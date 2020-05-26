package main

import (
	"fmt"
)

func change(val *int) {
	*val = 55
	fmt.Println("in func ", val)
}

//Local variables  can be accessed later
func hello() *int {
	i := 5
	return &i
}
func main() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
	*a++
	fmt.Printf("a incremented %d value b %d", *a, b)
	var c *int
	if c == nil {
		fmt.Println("C is not initialized ")

	}
	c = &b

	fmt.Println("Before func call", *c, *a, b)
	change(c)
	fmt.Println("After func call", *c, a, b, &b)

	//create pointer using new method
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)

	//access local vaiable of function which is not possible in C
	d := hello()
	fmt.Println("Value of d", *d)
}
