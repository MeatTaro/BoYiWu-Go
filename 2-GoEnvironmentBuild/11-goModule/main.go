package main

import (
	"fmt"
	"github.com/appleboy/com/random"
	"11-goModule/Hello"
)

func main()  {
	fmt.Println(hello.Hello(random.String(10)))
}

//go module
// main func 需要引入2種package 來執行 （random & Helllo）
//Hello : 本地package 需要import 本地位置""11-goModule/Hello""
//random : github 第三方package 需要import網址"github.com/appleboy/com/random"

//建立go module : go mod init 專案目錄
//go build 檔案名稱.go : 建立執行檔並更新go.mod載入package