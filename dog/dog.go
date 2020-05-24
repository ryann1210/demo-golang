package dog

import "fmt"

type Dog struct {
	Age int
}

func (dog Dog) Bark() {
	fmt.Println("Dog: 汪汪汪")
}
