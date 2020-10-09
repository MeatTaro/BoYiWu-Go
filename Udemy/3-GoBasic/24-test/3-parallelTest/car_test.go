package car

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c, err := New("", 100)
	if err != nil {
		t.Fatal("get errors :", err)
	}
	if c == nil {
		t.Error("car should not be nil")
	}
}

func TestNewAssert(t *testing.T) {
	c, err := New("", 100)
	assert.NotNil(t, err)
	assert.Nil(t, c)
}

func TestCar_GetName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		c    *Car
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "not input name",
			c: &Car{
				Name: "foo",
				Price: 100,
			},
			args: args{
				name: "",
			},
			want: "foo",
		},
		{
			name: "input name",
			c: &Car{
				Name: "foo",
				Price: 100,
			},
			args: args{
				name: "bar",
			},
			want: "bar",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()//平行測試
			if got := tt.c.GetName(tt.args.name); got != tt.want {
				t.Errorf("Car.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
