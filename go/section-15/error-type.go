package main

import (
	"errors"
	"fmt"
)

func main() {
	argError1 := errors.New("첫 번째 오류");
	argError2 := errors.New("두 번째 오류");
	argError3 := errors.New("세 번째 오류");
	
	// %w 서식 지정자를 통해 argError1를 감싼 error 값 생성
	error1 := fmt.Errorf("%w!!", argError1)
	fmt.Println(error1)                  // -> argError1을 감싼 error를 출력
	fmt.Println(error1.Error())          // -> argError1을 감싼 error를 출력 (위와 동일)
	fmt.Println(errors.Unwrap(error1))   // -> error가 감싼 argError1를 출력

	// %w 서식 지정자를 통해 여러 개의 error를 감싼 error 값 생성
	error2 := fmt.Errorf("%w %w %w", argError1, argError2, argError3)
	fmt.Println(error2)                  // -> argError1, argError2, argError3을 감싼 error를 출력
	fmt.Println(errors.Unwrap(error2))   // -> nil 출력
}