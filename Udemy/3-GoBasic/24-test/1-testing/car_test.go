package car

import (
	"testing"
)

func TestNew(t *testing.T)  {
	c, err := New("", 100)
	if err != nil {
		t.Error("got errors:", err)
	}
	if c == nil {
        t.Error("car should be nil")
    }
}