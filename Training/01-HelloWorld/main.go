package main

import (
	"fmt"
)

func Hello(name string) string {
	return fmt.Sprintf("Hi %s ", name)
}

func main()  {
	a := 1
	fmt.Printf(Hello("MeatTaro"))
	fmt.Println("learn golang")
	if a >= 1 {
		fmt.Println("a >= 1")
	}
}