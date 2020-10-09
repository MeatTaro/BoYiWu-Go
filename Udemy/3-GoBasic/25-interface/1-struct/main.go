package main

import (
	"fmt"
)
//建立struct
type Series struct {
	Series string
	Name string
}

func (m Series) ShowSeries()  {
	fmt.Println(m.Name, m.Series)
}

type Year struct {
	Series
	Year int
}

func (y Year) ShowYear() {
	fmt.Println(y.Name ,y.Year)
}

type MCU interface {
	ShowYear()
	ShowSeries()
}


func main()  {
	
	IronMan := Year {
		Series: Series{
			Series: "MCU",
			Name: "Iron Man",
		},
		Year: 2008,
	}
	TheIncredibleHulk := Year {
		Series: Series{
			Series: "MCU",
			Name: "The Incredible Hulk",
		},
		Year: 2008,
	}

	var film MCU
	film = IronMan
	film.ShowSeries()
	film.ShowYear()

	film = TheIncredibleHulk
	film.ShowSeries()
	film.ShowYear()
}

