package main

import "fmt"

func main() {
	x := [...]int{1, 2, 3, 4, 5}

	y := x[0:2] //Last index is not included
	z := x[1:4] //Capacity from index 1 to end

	fmt.Println(len(y), cap(y), len(z), cap(z)) //2 5 3 4
}
