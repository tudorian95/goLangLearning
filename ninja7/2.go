package main

import "fmt"

type person struct {
	name string
}

func (p person) changeMe(q *person) {
	(*q).name = "pwned"
}

func main() {
	p := person{"tudi"}
	fmt.Println(p)
	p.changeMe(&p)
	fmt.Println(p)
}

/*
package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "James Bond",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.name = "Miss Moneypenny"
	// (*p).name = "Miss Moneyp"
}
*/
