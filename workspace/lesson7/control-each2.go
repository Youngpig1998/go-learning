package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}

	var dst2 []*int
	for _, i := range src {
		dst2 = append(dst2, &i)
	}

	for _, p := range dst2 {
		fmt.Print(*p)
	}
	// 55555
}
