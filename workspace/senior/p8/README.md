# Go语言设计模式之函数式选项模式

 本文主要介绍了Go语言中函数式选项模式及该设计模式在实际编程中的应用。



## 为什么需要函数式选项模式？

最近看[go-micro/options.go](https://github.com/micro/go-micro/blob/master/options.go)源码的时候，发现了一段关于服务注册的代码如下：

```go
type Options struct {
	Broker    broker.Broker
	Cmd       cmd.Cmd
	Client    client.Client
	Server    server.Server
	Registry  registry.Registry
	Transport transport.Transport

	// Before and After funcs
	BeforeStart []func() error
	BeforeStop  []func() error
	AfterStart  []func() error
	AfterStop   []func() error

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Broker:    broker.DefaultBroker,
		Cmd:       cmd.DefaultCmd,
		Client:    client.DefaultClient,
		Server:    server.DefaultServer,
		Registry:  registry.DefaultRegistry,
		Transport: transport.DefaultTransport,
		Context:   context.Background(),
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}
```

当时呢，也不是很明白`newOptions`这个构造函数为什么要这么写，但是后面在微信群里看到有人也再发类似的代码问为什么要这么写，后来在群里讨论的时候才知道了这是一种设计模式–**函数式选项模式**。

可能大家看到现在也不是很明白我说的问题到底是什么，我把它简单提炼一下。

我们现在有一个结构体，定义如下：

```go
type Option struct {
	A string
	B string
	C int
}
```

现在我们需要为其编写一个构造函数，我们可能会写成下面这种方式：

```go
func newOption(a, b string, c int) *Option {
	return &Option{
		A: a,
		B: b,
		C: c,
	}
}
```

上面的代码很好理解，也是我们一直在写的。有什么问题吗？

我们现在来思考以下两个问题：

1. 我们可能需要为Option的字段指定默认值
2. Option的字段成员可能会发生变更



## 选项模式

我们先定义一个`OptionFunc`的函数类型

```go
type OptionFunc func(*Option)
```

然后利用闭包为每个字段编写一个设置值的With函数：

```go
func WithA(a string) OptionFunc {
	return func(o *Option) {
		o.A = a
	}
}

func WithB(b string) OptionFunc {
	return func(o *Option) {
		o.B = b
	}
}

func WithC(c int) OptionFunc {
	return func(o *Option) {
		o.C = c
	}
}
```

然后，我们定义一个默认的`Option`如下：

```go
var (
	defaultOption = &Option{
		A: "A",
		B: "B",
		C: 100,
	}
)
```

最后编写我们新版的构造函数如下：

```go
func newOption2(opts ...OptionFunc) (opt *Option) {
	opt = defaultOption
	for _, o := range opts {
		o(opt)
	}
	return
}
```

测试一下：

```go
func main() {
	x := newOption("nazha", "小", 10)
	fmt.Println(x)
	x = newOption2()
	fmt.Println(x)
	x = newOption2(
		WithA("沙"),
		WithC(250),
	)
	fmt.Println(x)
}
GO 复制 全屏
```

输出：

```go
&{nazha 小 10}
&{A B 100}
&{沙 B 250}
```

这样一个使用函数式选项设计模式的构造函数就实现了。这样默认值也有了，以后再要为Option添加新的字段也不会影响之前的代码。



# Reference

https://www.cnblogs.com/xiao-xue-di/p/14452331.html#_label1

https://www.cnblogs.com/wangtaobiu/p/16159653.html