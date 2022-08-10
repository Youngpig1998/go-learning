# 模版模式

​	模版模式 (Template Pattern) 定义一个操作中算法的骨架，而将一些步骤延迟到子类中。这种方法让子类在不改变一个算法结构的情况下，就能重新定义该算法的某些特定步骤。简单来说，模板模式就是将一个类中能够公共使用的方法放置在抽象类中实现，将不能公共使用的方法作为抽象方法，强制子类去实现，这样就做到了将一个类作为一个模板，让开发者去填充需要填充的地方。以下是模板模式的一个实现：

```go
package template

import "fmt"

type Cooker interface {
  fire()
  cooke()
  outfire()
}

// 类似于一个抽象类
type CookMenu struct {
}

func (CookMenu) fire() {
  fmt.Println("开火")
}

// 做菜，交给具体的子类实现
func (CookMenu) cooke() {
}

func (CookMenu) outfire() {
  fmt.Println("关火")
}

// 封装具体步骤
func doCook(cook Cooker) {
  cook.fire()
  cook.cooke()
  cook.outfire()
}

type XiHongShi struct {
  CookMenu
}

func (*XiHongShi) cooke() {
  fmt.Println("做西红柿")
}

type ChaoJiDan struct {
  CookMenu
}

func (ChaoJiDan) cooke() {
  fmt.Println("做炒鸡蛋")
}
```

这里来看下测试用例：

```go
func TestTemplate(t *testing.T) {
  // 做西红柿
  xihongshi := &XiHongShi{}
  doCook(xihongshi)

  fmt.Println("\n=====> 做另外一道菜")
  // 做炒鸡蛋
  chaojidan := &ChaoJiDan{}
  doCook(chaojidan)

}
```

