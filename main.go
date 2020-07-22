package main

import "fmt"

func main() {
	fmt.Println("This is 1st print statement")

	fmt.Println("Test CI")
	TestHello()
}

func TestHello() string {
	return "Hello World"
}
