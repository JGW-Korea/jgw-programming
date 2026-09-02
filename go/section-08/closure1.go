package main

import "fmt"

func increaseCnt() func()int {
	n := 0
	
	return func() int {
		if n <= 1 {
			n += 1
		} else {
			n *= n
		}
		return n
	}
}

func main() {
	cnt := increaseCnt()

	for i := range 5 {
		value := cnt()
		fmt.Printf("%d value: %d\n", i + 1, value)
	}
}