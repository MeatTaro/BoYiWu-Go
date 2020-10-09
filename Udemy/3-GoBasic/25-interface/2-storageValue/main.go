package main

import (
	"fmt"
)

type Human struct {
	name string
	age int
}

func main()  {
	temp := []interface{}{}

	a := 1
	b := "foo"
	c := Human {
		name: "bar",
		age: 21,
	}

	temp = append(temp, a, b, c)

	for i, value := range temp {
		if v, ok := value.(int); ok {
			fmt.Printf("%d-value:%v\n", i,v)
		}
		if v, ok := value.(string); ok {
			fmt.Printf("%d-value:%v\n", i,v)
		}
		if v, ok := value.(Human); ok {
			fmt.Printf("%d-value:%v\n", i,v)
		}
	}

	for i, value := range temp {
		switch v := value.(type) {
		case int :
			fmt.Printf("%d-value:%v\n", i, v)
		case string:
			fmt.Printf("%d-value:%v\n", i, v)
		case Human:
			fmt.Printf("%d-value:%v\n", i, v)
		default:
			fmt.Println("Unknown")
		}
	}
}
