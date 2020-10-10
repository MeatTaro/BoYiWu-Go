package main

import (
	"testing"
)

func Test_print01(t *testing.T)  {
	if print01(100) != "100" {
		t.Fatal("error")
	}
}

func Test_print02(t *testing.T)  {
	if print02(int64(100)) != "100" {
		t.Fatal("error")
	}
}

func Test_print03(t *testing.T)  {
	if print03(100) != "100" {
		t.Fatal("error")
	}
}

//waring 在windows執行go test -v -bench=. 時 . 需要以"."表示
//若要只執行bechmark 則加入 -run=none .
//若要測試函式的記憶體使用狀況，則加入 -benchmem .
func BenchmarkPrint01(b *testing.B)  {
	for i:=0; i < b.N; i++ {
		print01(100)
	}
}

func BenchmarkPrint02(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		print02(int64(100))
	}
}

func BenchmarkPrint03(b *testing.B)  {
	for i:=0; i < b.N; i++ {
		print03(100)
	}
}

//由結果可看出print02效能最好，因為strconv.FormatInt為最底層的函式