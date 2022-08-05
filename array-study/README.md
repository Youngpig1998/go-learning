# 数组Array的复制和传参

​	讲数组Array进行复制和传参的时候，都是将其整个复制，而非只复制一个指针。

------

# 用==比较数组

1、相同维数且含有相同个数元素的数组才可以比较

2、每个元素都相同才相等

------

# 切片

​	切片是围绕动态数组的概念构建的一种数据结构，可以按需自动增长和缩小。它是由一个地址指针，和两个表示长的和容量的整数组成的。相对于数组而言，使用切片的一个好处是：可以按需增加切片的容量。Golang内置的append()函数会处理增加长度时的所有操作细节。

## make函数

Golang里的make函数仅能用来创建以下数据类型：

- Slice 
- Map 
- Channel

Ref:https://golang.org/pkg/builtin/#make

讲切片Slice进行复制和传参的时候，也是一样整个复制，只不过Slice只是一个指针加两个int，所以速度比传递数组要快得多。如果改变了被赋值的Slice的值，则原先的Slice的值也会随之改变。（相当于C++中的传左引用）

## 切片截取

截取一段切片

newSlice  := slice[lowerBound:upperBound]     遵循左闭右开原则

只是复制了切片指针，没有对底层数组进行复制操作。所以对上述的newSlice进行操作还是会影响到原先slice的值。

## 切片共享存储结构

<img src=".\slice.jpg" alt="image-20220805210306272" style="zoom: 50%;" />

------

# Map

1、Map的赋值传参与Slice相同，相当于传递一个指针而已。

2、在访问Map时，如果访问的 Key 不不存在，Go仍会返回零值，所以我们无法通过返回 nil 来判断元素是否存在

​	然而在返回的同时Go还会返回一个bool值来告诉我们Key是否存在

```go
if v, ok := m["four"]; ok {
	t.Log("four", v)
} else {
	t.Log("Not existing")
}
```

