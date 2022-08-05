# 变量定义
## 全局变量

函数外定义的变量叫全局变量，以下是全局变量的定义方法。

* 方法1
```go 
var name type = value
```
* 方法2：注意，全局变量如果采用这个方式定义，那不能在全局范围内赋值，只能在函数体内给这个全局变量赋值
```go
var name type // value will be defaulted to 0, false, "" based on the type

/* 如果定义上面的全局变量，就不能紧接着在下一行通过name=value的方式对变量name做赋值，
比如name = 10，会编译报错：
 syntax error: non-declaration statement outside function body
*/
```
* 方法3
```go
var name = value 
```
* 方法4
```go
var (
	v1 int = 10
	v2 bool = true
)
var (
	v5 int   // the value will be defaulted to 0
	v6 bool  // the value will be defaulted to false
)
var (
	v3 = 20
	v4 = false
)
```

* **全局变量允许声明后不使用**，编译不会报错。

  

## 局部变量

函数内定义的变量叫局部变量。

* 和全局变量的定义相比，多了以下定义方法
  * 方法5
  ```go
  name := value
  ```
  * 方法6
	```go
	var name type
	name = value
	```
	
* **局部变量定义后必须要被使用，否则编译报错**，报错内容为`declared but not used`。

  

## 多变量定义：

一次声明和定义多个变量

* 全局变量

  * 方法1

    ```go
    var a, b, c int = 1, 2, 3
    ```

  * 方法2

    ```go
    var a, b, c bool
    ```

  * 方法3

    ```go
    var a, b, c = 1, 2, "str"
    ```

* 局部变量：和全局变量相比，多了以下定义方法

  * 方法4

    ```go
    var a, b int
    a, b = 1, 2
    
    var c, d int
    c = 10
    d = 20
    ```

  * 方法5

    ```go
    a, b := 1, 2
    a1, b1 := 1, "str"
    ```



## 变量类型及其零值

* 零值：英文叫[zero vaue](https://go.dev/ref/spec#The_zero_value)，没有显示初始化的变量，Go编译器会给一个默认值，也叫零值。

* 数值：所有数值类型的零值都是0

  * 整数，零值是0。byte, rune, uintptr也是整数类型，所以零值也是0。
  * 浮点数，零值是0
  * 复数，零值是0+0i

* bool，零值是false

* 字符串，零值是空串""

* 指针：var a *int，零值是nil

  ```go
  num := 100
  var a * int = &num
  ```

* 切片：var a []int，零值是nil

  ```go
  var a []int = []int{1,2}
  list := [6]int{1,2} //size为6的数组，前面2个元素是1和2，后面的是默认值0
  ```

* map：var a map[string] int，零值是nil

  ```go
  dict := map[string] int{"a":1, "b":2}
  ```

* 函数：var a func(string) int，零值是nil

  ```go
  function := func(str string) string {
    return str
  }
  result := function("hello fans")
  fmt.Println("result=", result)
  ```

* channel：var a chan int，通道channel，零值是nil

  ```go
  var a chan int = make(chan int)
  var b = make(chan string)
  c := make(chan bool)
  ```

* 接口：var a interface_type，接口interface，零值是nil

  ```go
  type Animal interface {
    speak()
  }
  
  type Cat struct {
    name string
    age int
  }
  
  func(cat Cat) speak() {
    fmt.Println("miao...")
  }
  
  // 定义一个接口变量a
  var a Animal = Cat{"gaffe", 1}
  a.speak() // miao...
  ```

* 结构体:  var instance StructName，结构体里每个field的零值是对应field的类型的零值

  ```go
  type Circle struct {
    radius float64
  }
  
  var c1 Circle
  c1.radius = 10.00
  ```

------

# 常量

* 常量定义的时候必须赋值，定义后值不能被修改

* 常量(包括全局常量和局部常量)可以定义后不使用，局部变量定义后必须使用，否则编译报错

* 常量可以用来定义枚举

* iota，特殊常量，可以理解为const语句块里的行索引，值从0开始

* 常量的定义方法

  * 方法1

    ```go
    const a int = 10
    const b bool = false
    ```

  * 方法2

    ```go
    const a = 10
    const b = false
    ```

  * 多个常量同时定义

    ```go
    const a, b int = 1, 2
    ```

  * iota，特殊常量，可以理解为每个独立的const语句块里的行索引

    ```go
    const a int = iota // the value of a is 0
    const b = iota // the value of b is still 0
    ```

  * 定义枚举方法1

    ```go
    const (
      unknown = 0
      male = 1
      female = 2
    )
    ```

  * 定义枚举方法2

    ```go
    const (
      unknown = iota // the value of unknown is 0
      male // the value of male is 1
      female // the value of female is 2
    )
    const (
      c1 = iota // the value of c1 is 0
      c2 = iota // the value of c2 is 1
      c3 = iota // the value of c3 is 2
    )
    ```

  * 注意事项

    * iota的值是const语句块里的行索引，行索引从0开始
    * const语句块里，如果常量没赋值，那它的值和上面的保持一样，比如下面的例子里class2=0, class6="abc"
    * 某个常量赋值为iota后，紧随其后的常量如果没赋值，那后面常量的值是自动+1，比如下面的例子里，class3的值是iota，该行的行索引是2，所以class3=2， class4常量紧随其后没有赋值，那class4=class3+1=3

    ```go
    const (
    	class1 = 0
    	class2 // class2 = 0
    	class3 = iota  //iota is 2, so class3 = 2
    	class4 // class4 = 3
    	class5 = "abc" 
    	class6 // class6 = "abc"
    	class7 = iota // class7 is 6
    )
    ```

    





## References

* https://go.dev/ref/spec#The_zero_value

