package main

import (
	"fmt"
	"os"
)

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
	str2 := fmt.Sprintf("%s is %d years old.", name, age)
	fmt.Println(str2)

	// Printing #6. fmt.Sprintln()
	str3 := fmt.Sprintln("10 + 20", "= 30")
	fmt.Print(str3)

	// Printting #7. fmt.Fprint()
	file1, err1 := os.Create("output.txt")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer file1.Close()
	
	fmt.Fprint(file1, "Hello World!")

	// Printing #8. fmt.Fprintf()
	file2, err2 := os.Create("output2.txt")
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Fprintf(file2, "Hello World!!")

	// Printing #9. fmt.Append()
	buf1 := make([]byte, 0, 128)
	
	nametemp := fmt.Sprint("Name: ", "Kim", "\n")
	agetemp := fmt.Sprint("Age: ", 25, "\n")
	activetemp := fmt.Sprint("Active: ", true, "\n")
	
	buf1 = append(buf1, []byte(nametemp)...)
	buf1 = append(buf1, []byte(agetemp)...)
	buf1 = append(buf1, []byte(activetemp)...)

	os.WriteFile("user1.txt", buf1, 0644)

	buf2 := make([]byte, 0, 128)
	
	buf2 = fmt.Append(buf2, "Name: ", "Kim", "\n")
	buf2 = fmt.Append(buf2, "Age: ", 25, "\n")
	buf2 = fmt.Append(buf2, "Active: ", true, "\n")
	
	os.WriteFile("user2.txt", buf2, 0644)


	// Printing #10. fmt.Appendf()
	buf3 := make([]byte, 0, 128)

	buf3 = fmt.Appendf(buf3, "Name: %s\n", "Kim")
	buf3 = fmt.Appendf(buf3, "Age: %d\n", 25)
	buf3 = fmt.Appendf(buf3, "Active: %t\n", true)

	os.WriteFile("user3.txt", buf3, 0644)

	// Printing #11. fmt.Appendln()
	buf4 := make([]byte, 0, 128)

	buf4 = fmt.Appendln(buf4, "Name:", "Kim")
	buf4 = fmt.Appendln(buf4, "Age:", 25)
	buf4 = fmt.Appendln(buf4, "Active:", true)

	os.WriteFile("user4.txt", buf4, 0644)

	return 0
}

func main() {
	Printing()
}