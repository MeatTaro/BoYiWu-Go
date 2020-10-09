package main

import (
	"log"
	"strconv"
	"fmt"
)

func print01(num int) string {
	return fmt.Sprintf("%d", num)
}
func print02(num int64) string {
	return strconv.FormatInt(num, 10)
}
func print03(num int) string {
	return strconv.Itoa(num)
}

func main()  {
	log.Println(print01(10))
	log.Println(print02(10))
	log.Println(print03(10))

}