package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		"Amnanda",
		"Smith",
		34,
	}
	s, i := p1.speak()
	fmt.Println(s, i)
}

func (p person) speak() (string, string) {
	return fmt.Sprintf("I am %v", p.last), fmt.Sprintf("and I am %v\n", p.first)
}

/*
package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		"Amnanda",
		"Smith",
		34,
	}
	p1.speak()
}

func (p person) speak() {
	fmt.Printf("I am %v and I am %v\n", p.first, p.age)
}
*/
