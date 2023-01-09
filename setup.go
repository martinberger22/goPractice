package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("What's your name?")
	fmt.Scan(&name)
	fmt.Println("Hello", name)

	var age int
	fmt.Println("How old are you?")
	fmt.Scan(&age)
	fmt.Printf("User %v is %d years old \n", name, age)

	newMessage := fmt.Sprintf("This is %v, they are %d years old", name, age)
	fmt.Println(newMessage)
}
