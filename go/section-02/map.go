package main

import "fmt"

func sectionMap() {
	vMap := map[string]string{
		"Name":    "Vito",
		"Address": "NY",
	}

	for i := range len(vMap) {
		var key string
		
		if i == 0 {
			key = "name"
		} else {
			key = "id"
		}
		
		_, ok := vMap[key]
		
		if ok {
			fmt.Println(key, "가 존재합니다.")
		} else {
			fmt.Println(key, "가 존재하지 않습니다.")
		}
	}

	vMap2 := map[string]int{
		"Alice": 25,
		"Bob": 30,
		"Charlie": 35,
	}

	fmt.Println("Befor deletion:")
	for key, value := range vMap2 {
		fmt.Printf("Key(%v): %v\n", key, value)
	}

	// 맵(Map) 선언 이후 데이터 삭제
	delete(vMap2, "Alice")

	fmt.Println("After deletion:")
	for key, value := range vMap2 {
		fmt.Printf("Key(%v): %v\n", key, value)
	}

	fmt.Println()

	// ------------
	// Set
	// -----------
	values := [5]int{1, 2, 2, 3, 3}

	set := make(map[int]bool)

	for _, value := range values {
		if set[value] {
			continue
		}
		
		set[value] = true
	}

	for key := range set {
		fmt.Println(key)
	}
}