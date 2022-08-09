package main

import "fmt"

type student struct {
	name  string
	age   int
	class string
}

type stuOption func(*student)


func withAge(a int) stuOption {
	// 设置年龄
	return func(s *student) {
		s.age = a
	}
}

func withClass(c string) stuOption {
	// 设置课程
	return func(s *student) {
		s.class = c
	}
}

func newStudent(name string, options ...stuOption) *student {
	stu := &student{name: name, age: 20, class: "math"} // 年龄默认为20,课程默认为math
	for _, o := range options {
		o(stu)
	}
	return stu
}

func main() {
	s := newStudent("Lilith", withAge(10))
	s2 := newStudent("Ulrica", withAge(100), withClass("chinese"))
	fmt.Println(*s)    // {Lilith 10 math}
	fmt.Println(*s2)  // {Ulrica 100 chinese}
}
