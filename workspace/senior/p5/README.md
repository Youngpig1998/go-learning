# 什么是空结构体

我们说不包含任何字段的结构体叫做空结构体，可以通过如下的方式定义空结构体：

原生定义

```go
var a struct{}
```

类型别名

```go
type empty struct{}
var e empty
```



# 特点



## 地址相同

我们分别定义两个非空结构体和空结构体变量，然后取地址打印，发现空结构体变量的地址是相同的：

```go
// 定义一个非空结构体
type User struct {
    name string
}

func main() {
 
  // 两个非空结构体的变量地址不同
  var user1 User
  var user2 User
  fmt.Printf("%p \n", &user1) // 0xc000318670
  fmt.Printf("%p \n", &user2) // 0xc000318680
  
  // 定义两个空结构体，地址相同
  var first struct{}
  var second struct{}
  fmt.Printf("%p \n", &first)    // 0x1ca15f0 
  fmt.Printf("%p \n", &second)   // 0x1ca15f0 
}

```

我们知道 Go 语言中的变量传递都是值传递，对于传参前后的变量地址应该不同，我们通过传参的方式再来试一下：

```go
// 非空结构体
type NonEmptyUser struct {
    name string
}

// 空结构体
type EmptyUser struct{}

// 打印非空结构体参数地址
func testNonEmptyUser(user NonEmptyUser) {
    fmt.Printf("%p \n", &user)
}

// 打印空结构体参数地址
func testEmptyUser(user EmptyUser) {
    fmt.Printf("%p \n", &user)
}


func main() {
  
    // 两个非空结构体的变量地址不同
    var user1 NonEmptyUser
    fmt.Printf("%p \n", &user1) // 0xc0001986c0
    testNonEmptyUser(user1)     // 0xc0001986d0

  
    // 两个空结构体变量的地址相同
    var user2 EmptyUser
    fmt.Printf("%p \n", &user2) // 0x1ca25f0
    testEmptyUser(user2)        // 0x1ca25f0
  
}
```

发现对于非空结构体，传参前后的地址是不同的，但是对于空结构体变量，前后地址是一致的。



## 大小为0

在Go中，我们可以使用 unsafe.Sizeof 来计算一个变量占用的字节数，那么就举几个例子来看下：

```go
type EmptyUser struct{}

func main() {
    var i int
    var s string
    var m []string
    var u EmptyUser
  
    fmt.Println(unsafe.Sizeof(i)) // 8
    fmt.Println(unsafe.Sizeof(s)) // 16
    fmt.Println(unsafe.Sizeof(m)) // 24
    fmt.Println(unsafe.Sizeof(u)) // 0
}
```

可以看到空结构体占用的内存空间大小为0，同时对于空结构体的组合，占用空间大小也为0：

```go
// 空结构体的组合
type EmptyUser struct {
    name struct{}
    age  struct{}
}

func main() {
    var u EmptyUser
    fmt.Println(unsafe.Sizeof(u)) // 0
}
```



# 原理探究

为什么空结构体的地址都相同，而且大小都为0呢，我们一起来看下源码（go/src/runtime/malloc.go）：

```go
// base address for all 0-byte allocations
var zerobase uintptr

// 创建新的对象时，调用 mallocgc 分配内存
func newobject(typ *_type) unsafe.Pointer {
    return mallocgc(typ.size, typ, true)
}

func mallocgc(size uintptr, typ *_type, needzero bool) unsafe.Pointer {
    if gcphase == _GCmarktermination {
        throw("mallocgc called with gcphase == _GCmarktermination")
    }

    if size == 0 {
        return unsafe.Pointer(&zerobase)
    }
    ......
}
```

通过源码可以看出，创建新的对象时，需要调用 malloc.newobject() 进行内存分配，进一步调用 mallocgc 方法，在该方法中，如果判断类型的 size==0 ，固定返回 zerobase 的地址。 zerobase 是一个 uintptr 全局变量，占用 8 个字节。

因此我们可以确定的是，在Go语言中，所有针对 size==0 的内存分配，用的都是同一个地址 &zerobase ，所以我们在一开始看到的所有空结构体地址都相同。



# 使用场景

空结构体不包含任何数据，那么其应用场景也应该不在乎值内容，只当做一个占位符。在这种场景下，由于其不占用内存空间，使用空结构体既可以做到节省空间，又可以提供语义支持。



## 集合(Set)

使用过 Java 的同学应该都用过 Set 类型，Set 是保存不重复元素的集合，但是 Go 语言没有提供原生的 Set 类型。但是我们知道 Map 结构存储的是 key-value 类型，key 不允许重复，因此可以利用 Map 来实现 Set，key存储需要的数据，value 给个固定值就可以了。那么 value 给什么值好呢？这时候我们的 空结构体 就可以出场了，不占用空间，还可以完成占位操作，堪称完美，下面我们看怎么实现吧。

```go
// 定义了一个保存 string 类型的 Set集合
type Set map[string]struct{}

// 添加一个元素
func (s Set) Add(key string) {
    s[key] = struct{}{}
}

// 移除一个元素
func (s Set) Remove(key string) {
    delete(s, key)
}

// 是否包含一个元素
func (s Set) Contains(key string) bool {
    _, ok := s[key]
    return ok
}

// 初始化
func NewSet() Set {
    s := make(Set)
    return s
}
```



## channel中信号传输

空结构体 与 channel 可谓是一个经典组合，有时候我们只是需要一个信号来控制程序的运行逻辑，并不在意其内容如何。

在下面的例子中，我们定义了两个 channel 用于接收两个任务完成的信号，当接收到任务完成的信号时，就会触发相应的动作。

```go
func doTask1(ch chan struct{}) {
    time.Sleep(time.Second)
    fmt.Println("do task1")
    ch <- struct{}{}
}

func doTask2(ch chan struct{}) {
    time.Sleep(time.Second * 2)
    fmt.Println("do task2")
    ch <- struct{}{}
}

func main() {

    ch1 := make(chan struct{})
    ch2 := make(chan struct{})
    go doTask1(ch1)
    go doTask2(ch2)

    for {
        select {
        case <-ch1:
            fmt.Println("task1 done")
        case <-ch2:
            fmt.Println("task2 done")
        case <-time.After(time.Second * 5):
            fmt.Println("after 5 seconds")
            return
        }
    }
}
```



## [stop channel](../p6)

# 总结

本篇文章，我们学习了如下内容：

- 空结构体是一种特殊的结构体，不包含任何元素
- 空结构体的大小都为0
- 空结构体的地址都相同
- 由于空结构体不占用空间，从节省内存的角度出发，适用于实现Set结构、在 channel 中传输信号等