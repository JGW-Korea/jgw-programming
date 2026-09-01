package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type Car struct {
	name    string  "차량명"
	color   string  "색상"
	company string  "제조사"
}

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Gender string  `json:"gender"`
}

type User struct {
	Name  string  `validate:"required"`
	Email string  `validate:"required,email"`
	Age   int     `validate:"gte=0,lte=130"`
}

func main() {
	// 구조체 #1. 기본 사용법
	tag := reflect.TypeOf(Car{})

	for i := 0; i < tag.NumField(); i++ {
		fmt.Println(
			"ex1 :",
			tag.Field(i).Tag,
			tag.Field(i).Name,
			tag.Field(i).Type,
		)
		// i = 0 -> ex1 : 차량명 name string
		// i = 1 -> ex1 : 색상 name string
		// i = 2 -> ex1 : 제조가 name string
	}

	// 구조체 #2. JSON 직렬화
	person := Person{"John", 30, "Male"}
		
	jsonData, err := json.Marshal(person)
	if err != nil {
			fmt.Println(err)
	}
	
	fmt.Println(string(jsonData))

	// 구조체 #3. 검증
	validate := validator.New()
		
	user := &User{
		Name:  "",
		Email: "invalid-email",
		Age:   150,
	}
	
	if err := validate.Struct(user); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf(
				"필드 '%s'가 '%s' 규칙을 위반했습니다.\n",
				err.Field(),
				err.Tag(),
			)
		}
	}
}