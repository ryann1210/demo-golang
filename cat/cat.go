package cat

import "fmt"

type Cat struct {
	Age int
}

func (cat Cat) Bark() {
	fmt.Println("Cat: 汪汪汪")
}
