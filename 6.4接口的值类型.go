package main

import (
	"demo-golang/cat"
	"demo-golang/dog"
	"fmt"
)

type Animal interface {
	Bark()
}

func bark(animal Animal) {
	animal.Bark()
}

func main() {
	var animal Animal

	animal = dog.Dog{Age: 1}
	bark(animal)
	fmt.Printf("%T %v\n", animal, animal)

	animal = cat.Cat{Age: 2}
	bark(animal)
	fmt.Printf("%T %v\n", animal, animal)

	switch v := animal.(type) {
	case cat.Cat:
		fmt.Println("这是一只猫, Age: ", v.Age)
	case dog.Dog:
		fmt.Println("这是一只狗. Age: ", v.Age)
	}

	if realAnimal, ok := animal.(*cat.Cat); ok {
		fmt.Println(realAnimal.Age)
	} else {
		fmt.Println("No match.")
	}
}
