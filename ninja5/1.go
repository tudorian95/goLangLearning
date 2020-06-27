package main

import "fmt"

type person struct {
	first    string
	last     string
	flavours []string
}

func main() {
	p1 := person{
		first:    "Tudi",
		last:     "Arion",
		flavours: []string{"stawberry", "choco"},
	}

	p2 := person{
		first:    "Katy",
		last:     "Marin",
		flavours: []string{"cioco", "vanilla"},
	}
	for i := range p1.flavours {
		fmt.Println(p1.flavours[i])
	}
	for i := range p2.flavours {
		fmt.Println(p2.flavours[i])
	}

	for i, v := range p2.flavours { //same as above but in another way
		fmt.Println(i, v)
	}
}
