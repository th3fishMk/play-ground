package main

import "fmt"

func main() {
	fmt.Println("Hello world");

	var x int  = 42

	println(x)

	y := 42

	println(y)

	word := "variable"

	println(word)

	
	result := add(4,4)
	println(result)
	println(add(4,4))
}

func add(a int, b int) int  {
	return  a + b
}