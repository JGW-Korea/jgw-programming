package main

import "fmt"

// 부모 클래스(Parent Class)
type Employee struct {
		name   string
		salary float64
		bonus  float64
}

func (e Employee) Calculate() float64 {
		return e.salary + e.bonus
}

// 자식 클래스(Child Class)
type Executives struct {
		Employee     // 상속(is-a)
		specialBonus float64
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
		
		// 단축 표현 방식
		// ex := Executives{
		//     Employee{"lee", 5_000_000, 1_000_000},
		// 	   200,
		// }

		fmt.Println("emp #1: ", int(ep1.Calculate()))
		fmt.Println("emp #2: ", int(ep2.Calculate()))
	
		fmt.Println("ex #3: ", int(ex.Employee.Calculate() + ex.specialBonus))
		fmt.Println("ex #4: ", int(ex.Calculate()))
}