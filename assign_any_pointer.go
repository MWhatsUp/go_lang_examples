package main

import "fmt"

func main() {
	a := 5
	fmt.Printf("value: %v\n", a)
	
	var ptr any
	ptr = &a
	*(ptr.(*int)) = 2
	fmt.Printf("value: %v\n", a)

	variable, ok := ptr.(*string)
	if !ok || variable == nil {
		_, file, line, _ := runtime.Caller(0)
		panic(fmt.Sprintf("Error: Incorrect pointer type was used in query parameter handler (%s:%v)", file, line))
	}
	*variable = "asInt"
	fmt.Printf("value: %v\n", a)
}
