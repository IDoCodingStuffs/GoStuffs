package main

import "fmt"

type Artist struct {
	Name, Genre string
	Songs int
}

func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	singer := &Artist{Name:"Bob", Genre:"Country", Songs: 41}
	fmt.Printf("%s has released their %dth song\n", singer.Name, newRelease(singer))
	fmt.Printf("%s has a total of %d songs", singer.Name, singer.Songs)
}