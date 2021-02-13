package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	a.Print()
	// fmt.Println(a)
}

func (a A) Print() {
	fmt.Println("A")
}
