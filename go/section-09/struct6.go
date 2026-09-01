package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

type Executives struct {
	Employee // 상속(is-a)
	specialBonus float64
}

// 메서드 오버라이딩
func (e Executives) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

func main() {
	ep1 := Employee{"kim", 2_000_000, 150_000}
	ep2 := Employee{"park", 1_500_000, 200_000}

	ex := Executives{
		Employee: Employee{
			name:   "lee",
			salary: 5_000_000,
			bonus:  1_000_000,
		},
		specialBonus: 200,
	}

	// ex := Executives{
	// 	Employee{
	// 		name:   "lee",
	// 		salary: 5_000_000,
	// 		bonus:  1_000_000,
	// 	},
	// 	200,
	// }

	fmt.Println("emp #1: ", int(ep1.Calculate()))
	fmt.Println("emp #2: ", int(ep2.Calculate()))
	
	fmt.Println("ex #3: ", int(ex.Employee.Calculate() + ex.specialBonus))
	fmt.Println("ex #4: ", int(ex.Calculate()))
}

type Account struct {
		number  string
		balance  float64
		interest float64
}

func Constructor(number string, balance, interest float64) *Account {
		return &Account{number, balance, interest}
}