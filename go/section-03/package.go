package main

import (
	fs "fmt"
	"section-03/lib"
	"section-03/math"
	"section-03/user"
	"section-03/utils"
)

func Calculator(c math.Calculator) {
	fs.Println(c.Add(10, 30))
	fs.Println(c.subtract(50, 30))
}

func main() {
	lib.A()
	utils.B()

	

	var user user.User = user.User{
		Name: "as",
		Age: 1,
	}

	user.GetUser()

	math := math.Math{}
	
	Calculator(math)
}