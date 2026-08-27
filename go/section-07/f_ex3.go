package main

// 순수 함수(Pure Function) #2-1. 잘못된 예시: 외부 상태 의존
var tax float32 = 0.1

func calculate(price float32) float32 {
	// tax의 값이 변경되면 동일한 입력값에 대해 항상 동일한 결과를 반환하지 못하게 된다.
	// 이와 같이 순수 함수는 외부 상태에 의존하지 않아야 한다.
	return price + price * tax
}