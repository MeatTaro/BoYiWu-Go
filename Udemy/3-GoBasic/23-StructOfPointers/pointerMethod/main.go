package main

import (
	"time"
	"fmt"
)

type email struct {
	from	string
	to 			string
	send	string
}

func (e *email) From(s string) {
	e.from = s
}
func (e *email) To(s string) {
	e.to = s
}
func (e *email) Send() {
	fmt.Printf("From: %s , To: %s\n",e.from, e.to)
}
//此方法會在goroutine同步執行時造成pointer互相覆蓋,導致依序行喪失
// func main()  {
// 	e := &email{}
// 	for  i:=0; i < 10; i++ {
// 		e.From(fmt.Sprintf("example%02d@example.com", i))
// 		e.To(fmt.Sprintf("example%02d@example.com", i+1))
// 		e.Send()
// 	}
// }

func main()  {
//在每個迴圈開始執行method前都對struct進行初始化來確保i>>i+1
	for i:=0; i<10; i++ {
		go func(i int) {
			e := &email{}
			e.From(fmt.Sprintf("example%02d@example.com", i))
			e.To(fmt.Sprintf("example%02d@example.com", i+1))
			e.Send()
			}(i)	
	}
	time.Sleep(1*time.Second)
}

