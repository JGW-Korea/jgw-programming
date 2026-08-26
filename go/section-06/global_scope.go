// 전역 스코프(Global Scope)
package main

import "fmt"

var GLOBAL_VARIABLE string = "This global variables";

func main() {
	for i := range len(GLOBAL_VARIABLE) {
		if (GLOBAL_VARIABLE[i] == ' ') {
			continue;
		}

		fmt.Printf("Current letter: %c\n", GLOBAL_VARIABLE[i]);
	}
}