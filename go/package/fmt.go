package main

import "fmt"

func Printing() (n int) {
	const (
		name string = "Kim"
		age  int    = 32
	)

	// Printing #1. fmt.Print()
	fmt.Print(name, " is ", age, " years old.\n")
	n1, err1 := fmt.Print(name, " is ", age, " years old.\n")
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(n1)

	// Printing #2. fmt.Printf()
	fmt.Printf("10 + 20 = %d\n", 30)
	n2, err2 := fmt.Printf("10 + 20 = %d\n", 10 + 20)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(n2)
	
	// Printing #3. fmt.Println()
	fmt.Println("10 + 20 = 30")
	n3, err3 := fmt.Println("10 + 20 = 30")
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(n3)

	// Printing #4. fmt.Sprint()
	str1 := fmt.Sprint(1, 2, 3, 4, 5)
	fmt.Println(str1)

	// Printing #5. fmt.Sprintf()
	str2 := fmt.Sprintf("%s is %d years old.\n", name, age)
	fmt.Println(str2)

	// Printing #6. fmt.Sprintln()
	str3 := fmt.Sprintln("10 + 20", "= 30")
	fmt.Print(str3)

	return 0
}

func main() {
	Printing()
}