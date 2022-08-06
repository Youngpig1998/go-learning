# 基础数据类型
## 数字

* 整数：int, uint8, uint16, uint32, uint64, int8, int16, int32, int64

* 浮点数：float32, float64

* 复数：
  * complex64：实部和虚部都是float32类型的值
  
    ```go
    var v complex64 = 1 + 0.5i
    ```
  
  * complex128：实部和虚部都是float64类型的值
  
    ```go
    var v complex128 = 1 + 0.5i
    ```
  
  * **注意**：虚部为1的情况，1不能省略，否则编译报错
  
    ```go
    var v complex64 = 1 + i // compile error: undefined i
    var v complex64 = 1 + 1i // correct
    ```
  



## bool

值只能为`true`或`false`。



## 其它数字类型

* byte：等价于uint8，数据范围0-255，定义的时候超过这个范围会编译报错
* rune：等价于int32，数据范围-2147483648-2147483647
  
  * 字符串里的每一个字符的类型就是rune类型，或者说int32类型
* uint：在32位机器上等价于uint32，在64位机器上等价于uint64
* uintptr: 无符号整数，是内存地址的十进制整数表示形式，应用代码一般用不到（https://stackoverflow.com/questions/59042646/whats-the-difference-between-uint-and-uintptr-in-golang）

* reflect包的`TypeOf`函数或者`fmt.Printf`的`%T`可以用来获取变量的类型

    ```go
    var b byte = 10
    var c = 'a'
    fmt.Println(reflect.TypeOf(b)) // uint8
    fmt.Println(reflect.TypeOf(c)) // int32
    fmt.Printf("%T\n", c) // int32
    ```

## 类型转换

* 语法

  ```go
  type_name(expression)
  ```

* 示例

  ```go
  package main
  
  import "fmt"
  
  func main() {
      total_weight := 100
      num := 12
      // total_weight和num都是整数，相除结果还是整数
      fmt.Println("average=", total_weight/num) //  average= 8
      
      // 转成float32再相除，结果就是准确值了
      fmt.Println("average=", float32(total_weight)/float32(num)) // average= 8.333333
      
      /* 注意，float32只能和float32做运算，否则会报错，比如下例里float32和int相加，编译报错:
      invalid operation: float32(total_weight) + num (mismatched types float32 and int)
     
      res := float32(total_weight) + num
      fmt.Println(res)
      */
  }
  ```

  

* **注意**：Go不支持隐式类型转换，而且别名和原有类型之间也不能进⾏隐式类型转换。要做数据类型转换必须按照type_name(expression)方式做显示的类型转换

  ```go
  package main
  
  import "fmt"
  
  
  func main() {
      num := 10
      var f float32 = float32(num)
      fmt.Println(f) // 10
      
      /*
      不支持隐式类型转换，比如下例想隐式讲num这个int类型转换为float32就会编译报错:
       cannot use num (type int) as type float32 in assignment
       
      var f float32 = num
      */
  }
  ```


## References

* https://gfw.go101.org/article/basic-types-and-value-literals.html
* https://www.callicoder.com/golang-basic-types-operators-type-conversion/
