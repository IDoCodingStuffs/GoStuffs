package main

import "fmt"

var (
	foo2, bar2 string
)

var foo3 int
var bar3 int

var (
	foo4 string = "This"
	bar4 int = 12345
)

func main(){
	foo5, bar5 := "This", "that"

	changeVal := func() {
		foo5 = "These"
		bar5 = "Those"
	}
	changeVal()

	fmt.Println(foo5)
	fmt.Printf("%s", bar5)
}