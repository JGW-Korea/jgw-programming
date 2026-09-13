package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

type PowError struct {
	time    time.Time
	value   float64
	message string
}

func (e PowError) Error() string {
	return fmt.Sprintf("[%v] Error - Input Value(value: %g) - %s", e.time, e.value, e.message)
}

func Power(f, i float64) (float64, error) {
	if f == 0 {
		return 0, PowError{
			time: time.Now(),
			value: f,
			message: "0은 사용할 수 없습니다.",
		}
	}

	if math.IsNaN(f) {
		return 0, PowError{
			time: time.Now(),
			value: f,
			message: "숫자가 아닙니다.",
		}
	}
	
	if  math.IsNaN(i) {
		return 0, PowError{
			time: time.Now(),
			value: i,
			message: "숫자가 아닙니다.",
		}
	}

	return math.Pow(f, i), nil
}

func main() {
	if v, err := Power(10, 3); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("ex 1 : ", v)
	}
	
	if v, err := Power(0, 3); err != nil {
		fmt.Println(err.(PowError).value)
		fmt.Println(err.(PowError).message)
		fmt.Println(err.(PowError).time)
		log.Fatal(err)
	} else {
		fmt.Println("ex 1 : ", v)
	}
}