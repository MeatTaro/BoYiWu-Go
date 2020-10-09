package car

import (
	"errors"
)

type Car struct {
	Name string
	Price float32
}

func (e *Car) GetName(name string) string {
	if name != "" {
		e.Name = name
	}
	return e.Name
}

func New(name string, price float32) (*Car, error) {
	if name == "" {
		return nil, errors.New("missing name")
	}
	return &Car {
		Name: name,
		Price: price,
	}, nil
}