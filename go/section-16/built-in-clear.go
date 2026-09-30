package main

import "fmt"

// clear 사용 예시 #1. 맵 또는 맵 기반 타입 clear 사용 가능
type MapTypeAlias = map[string]string
type MapTypeDeclear map[string]string

func useClearInMap() {
	clearTargetMap1 := map[string]string{
		"name":     "vito",
		"location": "New York",
	}
	clearTargetMap2 := MapTypeAlias{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	clearTargetMap3 := MapTypeDeclear{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	fmt.Println("before clearTargetMap1", clearTargetMap1) // { name: "vito", location: "New York" }
	fmt.Println("before clearTargetMap2", clearTargetMap2) // { key1: "value1", key2: "value2", key3: "value3" }
	fmt.Println("before clearTargetMap3", clearTargetMap3) // { key1: "value1", key2: "value2", key3: "value3" }
	clear(clearTargetMap1)
	clear(clearTargetMap2)
	clear(clearTargetMap3)
	fmt.Println("after clearTargetMap1", clearTargetMap1) // { }
	fmt.Println("after clearTargetMap2", clearTargetMap2) // { }
	fmt.Println("after clearTargetMap3", clearTargetMap3) // { }
}

// clear 사용 예시 #2. 슬라이스 또는 슬라이스 기반 타입 clear 사용 가능
type SliceTypeAlias = []int
type SliceTypeDeclear []int

func useClearInSlice() {
	clearTargetSlice1 := []int{1, 2, 3, 4, 5}
	clearTargetSlice2 := make(SliceTypeAlias, 5, 10)
	for idx, _ := range clearTargetSlice2 {
		clearTargetSlice2[idx] = (idx + 1) * 10
	}
	clearTargetSlice3 := SliceTypeDeclear{100, 200, 300, 400, 500}

	fmt.Println("before clearTargetSlice1", clearTargetSlice1) // [1, 2, 3, 4, 5]
	fmt.Println("before clearTargetSlice2", clearTargetSlice2) // [10, 20, 30, 40, 50]
	fmt.Println("before clearTargetSlice3", clearTargetSlice3) // [100, 200, 300, 400, 500]
	clear(clearTargetSlice1)
	clear(clearTargetSlice2)
	clear(clearTargetSlice3)
	fmt.Println("after clearTargetSlice1", clearTargetSlice1) // [0, 0, 0, 0, 0]
	fmt.Println("after clearTargetSlice2", clearTargetSlice2) // [0, 0, 0, 0, 0]
	fmt.Println("after clearTargetSlice3", clearTargetSlice3) // [0, 0, 0, 0, 0]
}

func main() {
	useClearInMap()
	useClearInSlice()
}