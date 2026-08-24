package main

import (
	"fmt"
)

func slice() {
	var slice []int 
	
	fmt.Println(len(slice), cap(slice))

	for i := 0; i < 5; i++ {
		slice = append(slice, len(slice) + 1);
	}
	
	fmt.Println(slice);

	var slice2 []int = slice

	fmt.Printf("%p\n", slice)
	fmt.Printf("%p\n", slice2)

	var makeSlice []int = make([]int, 5, 10)
	
	for i := range len(makeSlice) {
		makeSlice[i] = i + 1
	}
	
	// var arr [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("address: %p, values: %v\n", slice, slice)
	fmt.Printf("address: %p, values: %v\n", slice2, slice2)

	fmt.Println("before:", cap(makeSlice))
	for i := range 30 {
		makeSlice = append(makeSlice, i + 5)

		if cap(makeSlice) % 10 == 0 {
			fmt.Println("after:", cap(makeSlice))
		}
	}
}