package main

import (
	"errors"
	"fmt"
)

func checkUserNameExist(user_name string) (bool, error) {
	if user_name == "foo" {
		//透過errors.New產生錯誤回報資訊
		return true, errors.New(fmt.Sprintf("username %s is exist", user_name))
	}
	return false, nil
}

func main()  {
	if _, err := checkUserNameExist("foo"); err != nil {
		fmt.Println(err)
	}

}