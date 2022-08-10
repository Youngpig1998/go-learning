# 策略模式

​	策略模式（Strategy Pattern）定义一组算法，将每个算法都封装起来，并且使它们之间可以互换。

​	在什么时候，我们需要用到策略模式呢？

​	在项目开发中，我们经常要根据不同的场景，采取不同的措施，也就是不同的**策略**。比如，假设我们需要对 a、b 这两个整数进行计算，根据条件的不同，需要执行不同的计算方式。我们可以把所有的操作都封装在同一个函数中，然后通过 if ... else ... 的形式来调用不同的计算方式，这种方式称之为**硬编码**。

​	在实际应用中，随着功能和体验的不断增长，我们需要经常添加 / 修改策略，这样就需要不断修改已有代码，不仅会让这个函数越来越难维护，还可能因为修改带来一些 bug。所以为了解耦，需要使用策略模式，定义一些独立的类来封装不同的算法，每一个类封装一个具体的算法（即策略）。下面是一个实现策略模式的代码：

```go
package strategy

// 策略模式

// 定义一个策略类
type IStrategy interface {
  do(int, int) int
}

// 策略实现：加
type add struct{}

func (*add) do(a, b int) int {
  return a + b
}

// 策略实现：减
type reduce struct{}

func (*reduce) do(a, b int) int {
  return a - b
}

// 具体策略的执行者
type Operator struct {
  strategy IStrategy
}

// 设置策略
func (operator *Operator) setStrategy(strategy IStrategy) {
  operator.strategy = strategy
}

// 调用策略中的方法
func (operator *Operator) calculate(a, b int) int {
  return operator.strategy.do(a, b)
}
```

在上述代码中，我们定义了策略接口 IStrategy，还定义了 add 和 reduce 两种策略。最后定义了一个策略执行者，可以设置不同的策略，并执行，例如：

```go

func TestStrategy(t *testing.T) {
  operator := Operator{}

  operator.setStrategy(&add{})
  result := operator.calculate(1, 2)
  fmt.Println("add:", result)

  operator.setStrategy(&reduce{})
  result = operator.calculate(2, 1)
  fmt.Println("reduce:", result)
}
```





------



## Implementation

Implementation of an interchangeable operator object that operates on integers.

```go
type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}
```

## Usage

### Addition Operator

```go
type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}
```

```go
add := Operation{Addition{}}
add.Operate(3, 5) // 8
```

### Multiplication Operator

```go
type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}
```

```go
mult := Operation{Multiplication{}}

mult.Operate(3, 5) // 15
```



