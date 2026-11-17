package main

import (
	"fmt"
	"strings"
)

func useStringsContains() {
	var str, substr string
	
	fmt.Printf("두 문자열을 공백으로 구분하여 작성해주세요: ")
	fmt.Scan(&str, &substr)
	
	if ok := strings.Contains(str, substr); ok {
		fmt.Printf("%s 안에 %s가 포함되어 있습니다.\n", str, substr)
	} else {
		fmt.Printf("%s 안에 %s가 포함되어 있습니다.\n", str, substr)
	}
}

func useStringsCut() {
	show := func (s, seq string) {
		before, after, found := strings.Cut(s, seq)
		fmt.Printf(
			"Cut(%q, %q) = %q, %q, %v\n",
			s,
			seq,
			before,
			after,
			found,
		)
	}
	
	show("Gopher", "Go")      // Cut(Gopher, Go) = "", pher, 
	show("Gopher", "ph")      // Cut(Gopher, Go) = Go, er, 
	show("Gopher", "er")      // Cut(Gopher, Go) = Goph, , 
	show("Gopher", "Badger")  // Cut(Gopher, Go) = "", pher, 

}

func main() {
	// useStringsContains()
	useStringsCut()
}