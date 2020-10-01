package main

import (
	"fmt"
	_ "2-GoEnvironmentBuild/22-initFunc/bar"
	_ "2-GoEnvironmentBuild/22-initFunc/foo"
)

var global = convert()

//init()實際執行時間為main之前

func init()  {
	global = 0
}

func convert() int {
	return 100
}

func main()  {
	fmt.Println("global is", global)
}
