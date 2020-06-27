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

	m := map[string]person{"Arion": p1, "Katy": p2} //mm := map[string]person{p1.last: p1, p2.last: p2}
	fmt.Println(m)
}

/*
package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}

*/
