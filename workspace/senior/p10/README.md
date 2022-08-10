# 工厂模式

​	工厂模式（Factory Pattern）是面向对象编程中的常用模式。在 Go 项目开发中，你可以通过使用多种不同的工厂模式，来使代码更简洁明了。Go 中的结构体，可以理解为面向对象编程中的类，例如 Person 结构体（类）实现了 Greet 方法。

```go
type Person struct {
  Name string
  Age int
}

func (p Person) Greet() {
  fmt.Printf("Hi! My name is %s", p.Name)
}
```

​	有了 Person“类”，就可以创建 Person 实例。我们可以通过简单工厂模式、抽象工厂模式、工厂方法模式这三种方式，来创建一个 Person 实例。

​	这三种工厂模式中，简单工厂模式是最常用、最简单的。它就是一个接受一些参数，然后返回 Person 实例的函数：

```go
type Person struct {
  Name string
  Age int
}

func (p Person) Greet() {
  fmt.Printf("Hi! My name is %s", p.Name)
}

func NewPerson(name string, age int) *Person {
  return &Person{
    Name: name,
    Age: age,
  }
}
```

​	和`p：=＆Person {}`这种创建实例的方式相比，简单工厂模式可以确保我们创建的实例具有需要的参数，进而保证实例的方法可以按预期执行。例如，通过NewPerson创建 Person 实例时，可以确保实例的 name 和 age 属性被设置。

​	再来看抽象工厂模式，它和简单工厂模式的唯一区别，**就是它返回的是接口而不是结构体。**通过返回接口，可以在你不公开内部实现的情况下，让调用者使用你提供的各种功能，例如：

```go
type Person interface {
  Greet()
}

type person struct {
  name string
  age int
}

func (p person) Greet() {
  fmt.Printf("Hi! My name is %s", p.name)
}

// Here, NewPerson returns an interface, and not the person struct itself
func NewPerson(name string, age int) Person {
  return person{
    name: name,
    age: age,
  }
}
```

​	上面这个代码，定义了一个不可导出的结构体person，在通过 NewPerson 创建实例的时候返回的是接口，而不是结构体。通过返回接口，我们还可以实现多个工厂函数，来返回不同的接口实现，例如：

```go
// We define a Doer interface, that has the method signature
// of the `http.Client` structs `Do` method
type Doer interface {
  Do(req *http.Request) (*http.Response, error)
}

// This gives us a regular HTTP client from the `net/http` package
func NewHTTPClient() Doer {
  return &http.Client{}
}

type mockHTTPClient struct{}

func (*mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
  // The `NewRecorder` method of the httptest package gives us
  // a new mock request generator
  res := httptest.NewRecorder()

  // calling the `Result` method gives us
  // the default empty *http.Response object
  return res.Result(), nil
}

// This gives us a mock HTTP client, which returns
// an empty response for any request sent to it
func NewMockHTTPClient() Doer {
  return &mockHTTPClient{}
}
```

​	NewHTTPClient和NewMockHTTPClient都返回了同一个接口类型 Doer，这使得二者可以互换使用。当你想测试一段调用了 Doer 接口 Do 方法的代码时，这一点特别有用。因为你可以使用一个 Mock 的 HTTP 客户端，从而避免了调用真实外部接口可能带来的失败。来看个例子，假设我们想测试下面这段代码：

```go
func QueryUser(doer Doer) error {
  req, err := http.NewRequest("Get", "http://iam.api.marmotedu.com:8080/v1/secrets", nil)
  if err != nil {
    return err
  }

  _, err := doer.Do(req)
  if err != nil {
    return err
  }

  return nil
}
```

其测试用例为：

```go
func TestQueryUser(t *testing.T) {
  doer := NewMockHTTPClient()
  if err := QueryUser(doer); err != nil {
    t.Errorf("QueryUser failed, err: %v", err)
  }
}
```

​	另外，在使用简单工厂模式和抽象工厂模式返回实例对象时，都可以返回指针。例如，简单工厂模式可以这样返回实例对象：

```go
return &Person{
  Name: name,
  Age: age
}
```

抽象工厂模式可以这样返回实例对象：

```go
return &person{
  name: name,
  age: age
}
```

​	**在实际开发中，我建议返回非指针的实例，因为我们主要是想通过创建实例，调用其提供的方法，而不是对实例做更改。如果需要对实例做更改，可以实现SetXXX的方法。通过返回非指针的实例，可以确保实例的属性，避免属性被意外 / 任意修改。**

​	在**简单工厂模式**中，依赖于唯一的工厂对象，如果我们需要实例化一个产品，就要向工厂中传入一个参数，获取对应的对象；如果要增加一种产品，就要在工厂中修改创建产品的函数。这会导致耦合性过高，这时我们就可以使用**工厂方法模式**。

​	在工厂方法模式中，依赖工厂函数，我们可以通过实现工厂函数来创建多种工厂，将对象创建从由一个对象负责所有具体类的实例化，变成由一群子类来负责对具体类的实例化，从而将过程解耦。下面是**工厂方法模式**的一个代码实现：

```go
type Person struct {
  name string
  age int
}

func NewPersonFactory(age int) func(name string) Person {
  return func(name string) Person {
    return Person{
      name: name,
      age: age,
    }
  }
}
```

然后，我们可以使用此功能来创建具有默认年龄的工厂：

```go
newBaby := NewPersonFactory(1)
baby := newBaby("john")

newTeenager := NewPersonFactory(16)
teen := newTeenager("jill")
```

