package car

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T)  {
	c, err := New("", 100)
	if err != nil {
		t.Fatal("got errors:", err)
	}
	if c == nil {
		t.Error("car should not be nil")
	}
}

func TestAssert(t *testing.T) {
	c, err := New("", 100)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Nil(t, c)
}