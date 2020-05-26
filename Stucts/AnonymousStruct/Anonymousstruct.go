package main

import (
	"fmt"
)

//Anonymous struct means struct with no name .. just emp3 is created and no named struct is created.
func main() {
	emp3 := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)
}
