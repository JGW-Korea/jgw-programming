package main

import "fmt"

type User struct {
	Name     string
	Age      int
	Location string
}

func (user User) UpdateUserInfo() {
	fmt.Printf("memory address #2. %p\n", &user)      // user.Location: NewYork

	user.Name = "Micheal"
	user.Age = 29
	user.Location = "NewYork"
}

func main() {
	user := User{
		Name: "Vito",
		Age: 59,
		Location: "NewYork",
	}
	
	// 결과를 바꾸기 이전의 user 프로퍼티 출력
	fmt.Println("Before:")
	fmt.Println("user.Name:", user.Name)           // user.Name: Vito
	fmt.Println("user.Age:", user.Age)             // user.Age: 59
	fmt.Println("user.Location:", user.Location)   // user.Location: NewYork
	fmt.Printf("memory address #1. %p\n", &user)      // user.Location: NewYork
	
	user.UpdateUserInfo()
	fmt.Println()
	
	// 결과를 바꾼 이후 user 프로퍼티 출력 (값 변경)
	fmt.Println("After:")
	fmt.Println("user.Name:", user.Name)           // user.Name: Micheal
	fmt.Println("user.Age:", user.Age)             // user.Age: 29
	fmt.Println("user.Location:", user.Location)   // user.Location: NewYork
	fmt.Printf("memory address #3. %p\n", &user)      // user.Location: NewYork
}