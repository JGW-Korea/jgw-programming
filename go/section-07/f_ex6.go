package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func A(user User) {
	user.Name = user.Name + "_HOLLY_SHIT"
}

func B(user *User) {
	user.Age = user.Age * 10
}

func createDefaultUser() User {
	return User{
		Name: "A",
		Age: 20,
	}
}
func createPointerUser() *User {
	return &User{
		Name: "B",
		Age: 20,
	}
}

func main() {
	var a User = createDefaultUser()
	
	fmt.Println("#1.", a)

	A(a)
	fmt.Println("#2.", a)
	
	B(&a)
	fmt.Println("#3.", a)

	// ---

	var b *User = createPointerUser()
	
	fmt.Println("#1.", b)

	A(*b)
	fmt.Println("#2.", b)
	
	B(b)
	fmt.Println("#3.", b)

	var a3 int = 10
	
	{
		var a3 int = 20
		var b3 int = 30
		
		fmt.Println(a3 + b3) // 50
	}
	
	fmt.Println(a3 + b3)   // 
}