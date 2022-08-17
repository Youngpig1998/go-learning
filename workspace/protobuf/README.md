# Protobuf

​	关于 protobuf 如何定义 message，及字段规则相关内容

## Protobuf介绍

​	Protobuf 是一种与平台无关、语言无关、可扩展且轻便高效的序列化数据结构的协议，可以用于网络通信和数据存储。（具体做什么的百度一下都有）

## Protobuf使用

​	使用 Protobuf 的流程基本就是：先创建 .proto 文件定义消息格式，然后用内嵌的 protoc 编译

## 创建 proto 文件

​	创建 .proto 文件，其实就相当于定义数据结构，规定一下我们发消息的格式和内容是什么。

### （1）例1：定义一个最基本的 message

​	假如我想叫小黑来我家吃饭，我决定给他发一个邀请，那么我的消息就可以这样定义：

```protobuf
message Invite {
	required string host = 1;
	required string name = 2;
	required string address = 3;
	optional string info = 4;
}
```

​	message 是用于定义**消息**的关键字，Invite 是消息的名字，消息名是根据需要自定义的。host，name，address，info 是我想告诉对方邀请人(host)是谁，想要邀请谁(name)，做客地址(address)是什么，以及附加的说明信息(info)，直观来看这些信息的类型应该都是 string 类型，另外根据常识前三项都是必须要写的，不然对方怎么知道是谁请客？所以把它们都设置为 required 类型（也就是必须写上的），至于最后的附加说明信息 info，也许需要说明一下为啥请人家吃饭吧，但是如果关系特铁还需要原因吗？不需要（赶紧来就完了），所以类型是 optional（也就是可选的）。
​	**最后的序号只是规定一下字段的顺序，只要一个消息里的字段序号不要重复就好。**



### （2）例2：定义含有枚举字段的message

​	假如我想让小黑本周一或者周五来我家吃饭，其他时间我都不在家，那么就可以添加一个枚举类型让小黑自己选一天：

```protobuf
enum Time {
	Monday = 0;
	Friday = 1;
}
message Invite {
	required string host = 1;
	required string name = 2;
	required string address = 3;
	required Time time = 4;
	optional string info = 5;
}
```


​	enum 是定义枚举类型的关键字，相当于 C++ 中的 enum，但不同点在于枚举值之间的分隔符是分号，不是逗号。上面的Time 为枚举变量的名字，Monday 和 Friday 都是枚举值，0 和 1 表示枚举值所对应的实际整型值，可以为枚举值指定任意的整型数值，不是必须从 0 开始定义。

### （3）例3：定义含有嵌套消息字段的message

enum 是定义枚举类型的关键字，相当于 C++ 中的 enum，但不同点在于枚举值之间的分隔符是分号，不是逗号
Time 为枚举变量的名字，Monday 和 Friday 都是枚举值，0 和 1 表示枚举值所对应的实际整型值，可以为枚举值指定任意的整型数值，不是必须从 0 开始定义
（3）例3：定义含有嵌套消息字段的message

​	假如我想把做客信息单独拿出来作为一个消息，可以进行如下定义：

```protobuf
enum Time {
	Monday = 0;
	Friday = 1;
}
message Info {
	required address = 1;
	required meal = 2;
	optional string des = 3;
}
message Invite {
	required string host = 1;
	required string name = 2;
	required Time time = 3;
	required Info info = 4;
}
```


​	Invite 消息的定义中包含另外一个消息类型 Info info 作为其字段。嵌套的消息是被定义在同一个 .proto 文件中的，如果想要将其他 .proto 文件中定义的消息嵌套进来，可以使用 import 关键字。

## 正儿八经的说明

Invite 消息的定义中包含另外一个消息类型 Info info 作为其字段
嵌套的消息是被定义在同一个 .proto 文件中的，如果想要将其他 .proto 文件中定义的消息嵌套进来，可以使用 import 关键字
正儿八经的说明

### 字段规则

​	例1里面对于各个说明信息的限制 required，optional 实际叫字段规则或者限定符，规定了一个 message 里字段是否必须，以及出现的次数。

| 字段规则 |                             说明                             |
| :------: | :----------------------------------------------------------: |
| required | 字段必须出现且仅能出现一次**（该字段在protoc 3中已被禁止）** |
| optional |                    字段可出现 0 次或 1 次                    |
| repeated |                字段可出现任意次（包括 0 次）                 |

1. 在每个 message 中至少要有一个 required 类型的字段。
2. 每个 message 中可以有任意个 optional 类型的字段。
3. 如果想要在原有的消息协议中添加新的字段，同时还要保证老版本的程序能够正常读取或写入，那么对于新添加的字段必须是optional或repeated。道理非常简单，老版本程序无法读取或写入新增的required限定符的字段。



### 序号

​	最后的序号 1 2 3 … 表示不同字段在序列化后的二进制数据中的布局位置，比如上面的例子中 name 字段编码后的数据一定位于 host 字段之后。

​	**需要注意的是该值在同一 message 中不能重复。**另外，对于 Protocol Buffer 而言，标签值为 1 到 15 的字段在编码时可以得到优化，既标签值和类型信息仅占有一个 byte，标签范围是 16 到 2047 的将占有两个 bytes，而 Protocol Buffer 可以支持的字段数量则为 2 的 29 次方减一。鉴于此，我们在设计消息结构时，可以尽可能考虑让 repeated 类型的字段标签位于 1 到 15 之间，这样便可以有效的节省编码后的字节数量。



### 字段定义格式

​	字段定义格式：[字段规则] [类型] [名称] = [字段编号]

### protoc编译

​	通过内置的 protoc 编译器对 protobuf 文件进行编译，通过如下命令生成接口代码：

```shell
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/xxx.proto
```

或

```shell
protoc --proto_path=$SRC_DIR --cpp_out=$DST_DIR $SRC_DIR/xxx.proto
```

-I 等同于 --proto_path，指示了待编译的 .proto 文件所在源目录，该选项可以同时指定多个

--go_out 表示生成go代码，同理 --cpp_out 表示生成 C++ 代码， --java_out 和 --python_out 分别表示生成 Java 和 Python 代码，其后的路径是生成的代码所存放的目录
最后的路径是待编译的 .protoc 文件



#### 示例：

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative route.proto
```


