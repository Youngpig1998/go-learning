# Go命令

## go build命令

### 一个Go项目在GOPATH下，会有如下三个目录

![Go语言基础go build命令用法及示例详解](http://www.zzvips.com/uploads/allimg/211208/1-21120Q05004.png)

-  bin存放编译后的可执行文件
-  pkg存放编译后的包文件
-  src存放项目源文件

一般，bin和pkg目录可以不创建，go命令会自动创建（如 go install），只需要创建src目录即可。



### 使用：

add.go

```go
package cal
// 两个数相加 首字母大写，表示public 方法
func Add(a,b int) int {
	return a+b
}
```

subtraction.go

```go
package cal
// 两个数相减 首字母大写，表示public 方法
func Subtraction(a,b int) int {
	return a-b
}
```

main.go

```go
package main
import (
    "fmt"
    "cal"  //到入自定义的包，cal必须在src文件下，必须和main同一级
)
func main() {
    fmt.Println("hello word")
    fmt.Println("相加",cal.Add(10,20))
    fmt.Println("相减",cal.Subtraction(40,50))
}
```

1. 普通包 【非main包】

> `go build add.go` 【编译add.go,不生成exe执行文件】
> `go build -o add.exe add.go` 【指定生成exe执行文件，但是不能运行此文件，不是main包】

2. main包【package main】

> `go build main.go` 【生成exe执行文件】
> `go build -o main.exe main.go` 【指定生成main.exe执行文件】

3. 项目文件夹下有多个文件
   进入文件的目录

> `go build` 【默认编译当前目录下的所有go文件】
> `go build add.go subtraction.go` 【编译add.go 和 subtraction.go】



## 注意：

1. 如果是普通包，当你执行`go build`之后，它不会产生任何文件。【非main包】

2. 如果是main包，当你执行`go build`之后，它就会在当前目录下生成一个可执行文件exe。如果你需要在$GOPATH/bin下生成相应的文件，需要执行`go
   install`，或者使用`go build -o 路径/xxx.exe xxx.go`

3. 如果某个项目文件夹下有多个文件，而你只想编译某个文件，就可在go build之后加上文件名，例如`go build xxx.go`；`go build`命令默认会编译当前目录下的**所有go文件**。

4. 你也可以指定编译输出的文件名。我们可以指定`go build -o xxxx.exe`，默认情况是你的package名（main包），或者是第一个源文件的文件名（main包）。

5. `go build`会忽略目录下以“_”或“.”开头的go文件。