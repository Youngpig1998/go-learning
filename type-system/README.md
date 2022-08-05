# Go类型系统

- Built-in Types (int/string/byte/float32)
- Reference Types (map/slice/channels/functions/method)
- Custom Types (struct)

------

# 值接收者与指针接收者

​	值接收者Value Receiver：调用时使用这个值的副本来执行

​	指针接收者Pointer Receiver：共享调用方法时接收者所指向的值

​	所有参数都是值传递：slice，map，channel 会有传引⽤用的错觉

------

​	在一个赋值语句句中可以对多个变量量进⾏行行同时赋值，如 a, b = b, a

​	Go语言不允许隐式类型转换

​	别名和原有类型也不能进行隐式类型转换

------

# 指针类型

1、不支持指针运算。**因此像Slice就不能作比较，如slicaA == sliceB会报错！**

2、string是值类型，其默认的初始化值为空字符串，而不是nil

------

# Is Go an object-oriented language?

​	Yes and no. Although Go has types and methods and allows an objectoriented style of programming, there is no type hierarchy. The concept of “interface” in Go provides a different approach that we believe is easy to use and in some ways more general. Also, the lack of a type hierarchy makes “objects” in Go feel much more lightweight than in languages such as C++ or Java.