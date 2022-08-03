package main

import "fmt"

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
	array3[0] = "changed"

	fmt.Println(array)
	fmt.Println(array2)
	fmt.Println(array3)

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

}
