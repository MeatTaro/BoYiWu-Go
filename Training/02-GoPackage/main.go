package main

import (
	"fmt"
	"02-GoPackage/controller"
)



func main()  {
	hi := controller.HelloWorld("MeatTaro")

	fmt.Println(hi)
}