package main

import "fmt"

type Car struct {
	name string
	price int
	color string
	tax int
}

func main() {
	bmw1 := Car{
		name: "520d",
		price: 500_000_000,
		color: "white",
		tax: 50_000_000,
	}
	
	bmw2 := Car{
		name: "220d",
		price: 600_000_000,
		color: "black",
		tax: 60_000_000,
	}

	fmt.Println("ex1 :", bmw1, &bmw1)
	fmt.Println("ex1 :", bmw2, &bmw2)
	fmt.Printf("ex1 : %p\n", &bmw1)
	fmt.Printf("ex1 : %p\n", &bmw2)
}