package main

import "fmt"

// copy 사용 방법 #1. 기본적인 방법
func basicCopyUsage() {
	dst := make([]int, 5, 10)
	src := make([]int, 10, 20)
	
	for idx, _ := range src {
		src[idx] = idx + 1
	}
	
	count := copy(dst, src)
	fmt.Println("Copy count:", count);
	fmt.Println("Copy value:", dst);
}

// copy 사용 방법 #2. []byte
func byteCopyUsage() {
	dst := make([]byte, 5, 10)
	src := "Hello, Golang"
	
	count := copy(dst, src)
	fmt.Println("Copy count:", count)
	fmt.Printf("Copy value: %s\n", dst)
}

func main() {
	basicCopyUsage()
	byteCopyUsage()
}