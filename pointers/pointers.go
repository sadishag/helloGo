package main

import "fmt"

func main() {
	var address *int  // declare an int pointer
	number := 42      // int
	address = &number // address stores the memory address of number
	value := *address // dereferencing the value

	fmt.Printf("address: %v\n", address)
	fmt.Printf("value: %v\n", value)
}
