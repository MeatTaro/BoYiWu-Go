package main

import (
	"errors"
	"fmt"
)
//透過struct定義多個錯誤回報訊息
type errUserNameExist struct {
	UserName string
}
//Error()定義errUserNameExist 為error型態
func (e errUserNameExist) Error() string {
	return fmt.Sprintf("username %s exit", e.UserName)
}
//error 為 interface{} 所以可以轉為任何型態
func isUserNameExist(err error) bool {
	//若ok為true,表示err可轉換為自定義struck型態
	_, ok := err.(errUserNameExist)
	return ok
}

func checkUserNameExist(user_name string) (bool, error) {
	if user_name == "foo" {
		
		return true, errUserNameExist{UserName: user_name}
	}
	if user_name == "bar" {
		
		return true, errors.New("UserName is exist")
	}
	return false, nil
}

func main()  {
	if _, err := checkUserNameExist("foo"); err != nil {
		if isUserNameExist(err) {
			fmt.Println(err)
		}
	}

	if _, err := checkUserNameExist("bar"); err != nil {
		if isUserNameExist(err) {
			fmt.Println(err)
		}
			fmt.Println("isErrUserNameExist is false")
	}
}