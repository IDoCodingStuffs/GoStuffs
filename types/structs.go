package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	Lat float64
	Lon float64

	Date time.Time
}

type Point struct {
	X, Y int
}

var (
	p = Point{1, 2}
	q = &Point{1, 2}
	r = Point{X: 1}
	s = Point{}
	)


func main(){
	event := Bootcamp{
		Lat: 34.012836,
		Lon: -118.495338,
		Date: time.Now(),
	}

	fmt.Printf("Event on %s, location (%f, %f)\n",
		event.Date, event.Lat, event.Lon)

	fmt.Println(p,q,r,s)

	x := new(Bootcamp)
	y := &Bootcamp{}

	fmt.Println(*x == *y)
}