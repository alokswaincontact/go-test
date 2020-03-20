package main

import "fmt"

func main() {
	x := [...]int{4, 8, 5}
	fmt.Println(x) // [4 8 5]

	//Last index is not included
	y := x[0:2]
	fmt.Println(y) // [4 8]

	z := x[1:3]
	fmt.Println(z) // [8 5]

	//The changes below change the underlying array
	y[0] = 1
	z[1] = 3
	fmt.Println(x) // [1 8 3]
}
