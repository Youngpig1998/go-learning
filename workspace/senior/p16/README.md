# Go语言中命名函数参数和命名函数返回值的注意事项



## 命名返回值（named_results）

1. Go 的返回值可被命名，它们会被视作**定义在函数顶部的变量**。
2. 返回值的名称应当具有一定的意义，它可以作为文档使用。

3. 没有参数的 return 语句返回已命名的返回值。也就是直接返回。

4. 直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。


## 匿名返回值 

​	函数签名中命名返回值变量，只指定返回值类型。由return 指定返回值。

​	任何一个非命名返回值（使用非命名返回值是很糟的编程习惯）在 return 语句里面都要明确指出包含返回值的变量或是一个可计算的值（就像上面警告所指出的那样）

​	**建议：尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂。**

```go
package main

import "fmt"

//命名返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//匿名返回值
func split02(sum int) (int, int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func main() {
	fmt.Println(split(17))
	fmt.Println(split(18))
}
```



## References

* https://yourbasic.org/golang/named-return-values-parameters/