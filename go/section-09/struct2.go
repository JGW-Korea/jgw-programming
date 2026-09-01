package main

import "fmt"

type Account struct {
	policyNumber  string
	balance       float64
	interest      float64
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	kim := Account{policyNumber: "245-901", balance: 100_000_000, interest: 0.015}
	lee := Account{policyNumber: "245-901", balance: 120_000_000}
	park := Account{policyNumber: "245-901", interest: 0.015}
	cho := Account{"245-904", 150_000_000, 0.03}

	fmt.Println("ex 1:", kim)
	fmt.Println("ex 1:", lee)
	fmt.Println("ex 1:", park)
	fmt.Println("ex 1:", cho)
}