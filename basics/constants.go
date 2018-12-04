package main

import "fmt"

const This = 4567

const (
	StatusOK = 200
	StatusCreated = 201
	StatusAccepted = 202
	StatusNonAuthoritativeInfo = 203
	StatusNoContent = 204
	StatusResetContent = 205
	StatusPartialContent = 206
)

const (
	Beeg = 1 << 62
	Smol = Beeg >> 61
)

func main() {
	fmt.Println(Beeg)
	fmt.Println(Smol)
}