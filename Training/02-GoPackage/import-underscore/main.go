package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	//底線作用為導入package裡所有init func
)

func main() {
	project := os.Getenv("GOLANG_PROJECT")
	fmt.Println(project)
}


// package autoload

// /*
// 	You can just read the .env file on import just by doing
// 		import _ "github.com/joho/godotenv/autoload"
// 	And bob's your mother's brother
// */

// import "github.com/joho/godotenv"

// func init() {
// 	godotenv.Load()
// }