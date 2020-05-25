package main

import (
	"fmt"
)

func main() {
	// a := []int
	// //zero value of slice will be nil
	// if a == nil {
	// 	fmt.Printf("Not yet Initialized ")
	// }
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("arr %v type of %T", arr, arr)
	fmt.Println(arr, len(arr), cap(arr))
	slice := arr[2:5]

	fmt.Println("Slice:", slice, len(slice), cap(slice))
	slice = slice[:cap(slice)]
	fmt.Println("After Reslicing", slice, len(slice), cap(slice))

	//creating slice using MAke Syntax - func make([]T, len, cap)
	i := make([]int, 5, 5)
	fmt.Println(i)

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Printf("Cars %v type of %T", cars, cars)
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Printf("Cars %v type of %T after append ", cars, cars)
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

}
