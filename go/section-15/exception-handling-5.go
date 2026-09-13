package main

import (
	"errors"
	"fmt"
	"math"
)

func Power(f float64, i float64) (float64, error) {
	if (f == 0) {
		return 0, errors.New("0은 사용 불가능 합니다.")
	}

	return math.Pow(f, i), nil
}

func main() {
	if f, err := Power(7, 3); err != nil {
		// fmt.Printf("Error Message : %s\n", err)
		fmt.Printf("Error Message : %s\n", err.Error())
	} else {
		fmt.Println("ex 1 : ", f)
	}
	
	if f, err := Power(0, 3); err != nil {
		// fmt.Printf("Error Message : %s\n", err)
		fmt.Printf("Error Message : %s\n", err.Error())
	} else {
		fmt.Println("ex 2 : ", f)
	}
}