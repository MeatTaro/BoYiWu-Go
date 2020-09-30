package controller

import (
	"testing"
)

func TestHelloWorld(t *testing.T)  {
	hello := HelloWorld("MeatTaro")
	if hello != "Hi MeatTaro" {
		t.Errorf("Testing Fail")
	}
}