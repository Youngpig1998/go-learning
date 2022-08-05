package main

import (
	"fmt"
)

//Modify a slice
func modifySlice(slice []int) {
	slice[0] = 100
}

//Modify a map
func modifyMap(m map[string]int) {
	m["hehe"] = 100
}

func main() {

	//Array
	array := [3]string{"hello", "world", "whatever"}

	array2 := [...]string{"ni", "hao"}

	array3 := array
	sliceForArray1 := array[:]
	array3[0] = "changed"

	fmt.Println(array)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Print("sliceForArray1 is a slice. ")
	fmt.Println(sliceForArray1)

	//Slice(point,length,capacity)
	slice := []string{"hello", "world", "whatever", "sdwd", "Tseng"}
	fmt.Println(len(slice), cap(slice))

	newSlice := make([]int, 2, 2)
	fmt.Println(newSlice, len(newSlice), cap(newSlice))
	newSlice = append(newSlice, 1)
	fmt.Println(newSlice, len(newSlice), cap(newSlice))

	//The slice's value will be changed
	modifySlice(newSlice)
	fmt.Println(newSlice, len(newSlice), cap(newSlice))

	subSlice := slice[1:3]
	subSlice[0] = "Changed"
	fmt.Println(slice)

	//multiSlice
	multiSlice := make([][]int, 0)
	multiSlice = append(multiSlice, []int{1, 2, 4})
	multiSlice = append(multiSlice, []int{2, 4, 6})
	fmt.Println(multiSlice)

	//map
	scores := map[string]int{
		"xiaoming": 10,
		"zhangsan": 13,
		"deleted":  32,
	}

	delete(scores, "deleted")

	fmt.Println(scores["xiaoming"])
	fmt.Println(len(scores))

	score, exist := scores["zhang"]
	if exist {
		fmt.Println(score)
	}

	modifyMap(scores)
	fmt.Println(scores)

	for key, value := range scores {
		fmt.Printf("Key: %s \t Value: %d \n", key, value)
	}

	TestAccessNotExistingKey()

	TestMapWithFunValue()

	TestMapForSet()
}


func TestAccessNotExistingKey() {
	m1 := map[int]int{}
	fmt.Println(m1[1])
	m1[2] = 0
	fmt.Println(m1[2])
	m1[3] = 0
	if v, ok := m1[3]; ok {
		fmt.Printf("Key 3's value is %d", v)
	} else {
		fmt.Println("key 3 is not existing.")
	}
}

func TestMapWithFunValue() {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	fmt.Println(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet() {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		fmt.Printf("%d is existing\n", n)
	} else {
		fmt.Printf("%d is not existing\n", n)
	}
	mySet[3] = true
	fmt.Println(len(mySet))
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		fmt.Printf("%d is existing\n", n)
	} else {
		fmt.Printf("%d is not existing\n", n)
	}
}