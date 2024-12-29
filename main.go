package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanf("My name is %s", &name)

	fmt.Printf("Hello, %s", name)
}
