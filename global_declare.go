package main

import "log"

/*
Reference:
Non-declaration statement outside function body
https://yourbasic.org/golang/short-variable-declaration-outside-function/
*/

var number int

//Uncommenting this will lead to Compiler Error
//syntax error: non-declaration statement outside function body
//number = 1

func init() {
	number = 1
}

func main() {
	log.Println(number)
}
