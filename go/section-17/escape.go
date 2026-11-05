package main

import "fmt"

func main() {
	// Alert escape sequence - \a
	fmt.Println("주의가 필요합니다! \a")
	
	// Backspace escape sequence - \b
	fmt.Println("Hello\b<--")

	// From Feed escape sequence - \f
	fmt.Println("Hello\fWorld") // 

	// Newline escape sequence - \n
	fmt.Println("Hello\nWorld")

	// Carriage Return escape sequence - \r
	fmt.Println("Hello\rWorld")

	// Horizontal Tab escape sequence
	fmt.Println("Hello\tWorld")

	// Vertical Tab escape sequence
	fmt.Println("Hello\vWorld")

	fmt.Println("\355\225\234")  // 한
}
