package main

import (
	"fmt"
)

func main() {
	var nilSlice []string
	var emptySlice = []string{}

	//Nil:true Len:0 Capacity:0
	fmt.Printf("Nil:%v Len:%d Capacity:%d\n", nilSlice == nil, len(nilSlice), cap(nilSlice))

	//Empty:false Len:0 Capacity:0
	fmt.Printf("Empty:%v Len:%d Capacity:%d\n", emptySlice == nil, len(emptySlice), cap(emptySlice))
}
