package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	return &p
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{10, 5}
	rp := &r

	fmt.Println("area: ", rp.area())
	// point to a struct
	fmt.Print(&person{"bob", 29})
	s := person{"Sean", 50}
	fmt.Println(s.name)

	sp := &s

	fmt.Println(sp.age)

	sp.age = 51

	fmt.Println(s.age)
}
