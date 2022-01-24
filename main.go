package main

import (
	"fmt"
	"go-multipackageproject/lib"
)

func main() {
	a := 9
	b := 17

	fmt.Printf("%v + %v = %v\n", a, b, lib.Add(a, b))
}
