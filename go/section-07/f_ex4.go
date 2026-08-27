// 순수 함수(Pure Function) #2-2. 잘못된 예시: 외부 상태 변경

struct User type {
	age int;
	name string;
}

func aYearLater(user *User) {
	// 매개변수로 전달된 데이터의 상태를 변경하게 된다.
	// 이로 인해 해당 함수는 부수 효과가 발생하기 때문에 순수 함수 조건에 만족하지 않는다.
	user.name = "asd"
}
