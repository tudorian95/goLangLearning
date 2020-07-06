package main

import (
	"fmt"
)

type person struct {
	name string
}

type human interface {
	speak() string
}

func saySomething(h human) string {
	return h.speak()
}

func (p *person) speak() string {
	return fmt.Sprint("you have called the one named\t", p.name)
}

func main() {
	p1 := person{"tudi"}
	s := saySomething(&p1)
	//s1 := saySomething(p1)//cannot use p1 (type person) as type human in argument to saySomething:
	//person does not implement human (speak method has pointer receiver)
	fmt.Println(s)
}
