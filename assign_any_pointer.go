package main

import "fmt"

func main() {
	a := 5
	fmt.Printf("value: %v\n", a)
	var ptr any
	ptr = &a
	*(ptr.(*int)) = 2
	fmt.Printf("value: %v\n", a)
}
