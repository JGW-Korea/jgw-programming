package main

import "fmt"

type Animal struct {
	name    string
	weight  int
	breed   string
}

func (a Animal) bite() {
	bite := ""

	if a.breed == "dog" {
		bite = ": Dog bites!"
	} else {
		bite = ": Cat 할퀴다!"
	}

	fmt.Printf("%s : %s\n", a.name, bite)
}

func (a Animal) sounds() {
	bite := ""

	if a.breed == "dog" {
		bite = ": Dog barks!"
	} else {
		bite = ": Cat cries!"
	}

	fmt.Printf("%s : %s\n", a.name, bite)
}

// func (a Animal) run() {
// 	bite := ""
	
// 	if a.breed == "dog" {
// 		bite = ": Dog is running!"
// 	} else {
// 		bite = ": Cat is running!"
// 	}

// 	fmt.Printf("%s : %s\n", a.name, bite)
// }

type Dog = Animal
type Cat = Animal

type Behavivor interface {
	bite()
	sounds()
	run()
}

func action(animal Behavivor) {
	animal.bite()
	animal.sounds()
	animal.run()
}

func main() {
	// const a: interface = {...}

	infers := []Behavivor{
		Dog{"poll", 10, "dog"},
		Cat{"marry", 5, "cat"},
	}

	for _, infer := range infers {
		action(infer)
	}
}