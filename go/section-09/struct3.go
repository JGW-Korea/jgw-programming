package main

import "fmt"

// Golang 구조체 #1. 초기화 방법
type User struct {
	name string
	age int
	location string
}

func main() {
	// Golang 구조체 #2. 익명 구조체
	person := struct{
		name string
		age int
	} {
		name: "Kim",
		age: 20,
	}
	
	// Golang 구조체 #3-1. 구조체 타입 할당
	var user1 User = User{
		name: "Vito",
		age: 49,
		location: "New York, NY",
	}
	
	// Golang 구조체 #3-2. 일부 필드 제외
	user2 := User{
		name: "Sonny",
		age: 27,
	}
	
	// Golang 구조체 #3-3. 단축 구조체 타입 할당
	user3 := User{"Michael", 21, "Chicago"}

	fmt.Printf("value: %v, address: %p\n", person, &person)
	fmt.Printf("value: %v, address: %p\n", user1, &user1)
	fmt.Printf("value: %v, address: %p\n", user2, &user2)
	fmt.Printf("value: %v, address: %p\n", user3, &user3)
}