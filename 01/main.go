package main

import (
	"fmt"
	"os"
)

const (
	Monday, Tuesday, Wendesday = 1, 2, 3
)

const (
	a = iota
	b
	c
	d, e, f = iota, iota, iota
	g       = iota
	h       = "h"
	i
	j = iota
)

var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

func main() {
	fmt.Printf("Good day\n%v %v %v\n", Monday, Tuesday, Wendesday)
	fmt.Println(a, b, c, d, e, f, g, h, i, j)
	fmt.Println(HOME, USER, GOROOT)
}
