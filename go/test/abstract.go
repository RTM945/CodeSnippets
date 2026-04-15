package main

func main() {
  m := make(map[int]IBase)
	m[1] = NewSub()
	b := m[1]
	b.Common()
  /*
  Output: 
  Sub.Require
  Sub.Test2
  Base.Test1
  */
}

// Require 必须要子类实现的方法
type Require interface {
	Require()
}

// Hook 可能会被子类重写的方法
type Hook interface {
	Require
	Test1()
	Test2()
}

// IBase 类似Java的 interface -> abstract class -> sub class 套娃
type IBase interface {
	Hook
	Common()
}

// Base abstract class
type Base[T Hook] struct {
	hook T
}

func NewBase[T Hook](hook T) *Base[T] {
	res := &Base[T]{
		hook: hook,
	}
	return res
}

func (b *Base[T]) Common() {
	b.hook.Require()
	b.hook.Test2()
	b.hook.Test1()
}

func (*Base[T]) Test1() {
	fmt.Println("Base.Test1")
}

func (*Base[T]) Test2() {
	fmt.Println("Base.Test2")
}

// Sub 抽象类的子类
type Sub struct {
	*Base[*Sub]
}

func NewSub() *Sub {
	res := &Sub{}
	// 注意 %v 输出可能会有套娃死循环
	res.Base = NewBase[*Sub](res)
	return res
}

// Test1 复用 Base 的 Test1
//func (*Sub) Test1() {
//	fmt.Println("Sub.Test1")
//}

// Test2 重写 Base 的 Test2
func (*Sub) Test2() {
	fmt.Println("Sub.Test2")
}

// Require 必须要子类实现的Require 不写会编译报错
func (*Sub) Require() {
	fmt.Println("Sub.Require")
}
