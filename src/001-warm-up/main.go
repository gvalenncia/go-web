package main

import "fmt"

//#golangsnippet
type human interface {
	speak()
}

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	licenceKill string
	person      person
}

func main() {

	german := person{
		fname: "German",
		lname: "Valencia",
	}

	james := secretAgent{
		licenceKill: "666",
		person: person{
			fname: "James",
			lname: "Bond",
		},
	}

	saySomething(german)
	saySomething(james)
}

func (p person) speak() {
	fmt.Println("Hello, I am", p.fname, p.lname)
}

func (s secretAgent) speak() {
	fmt.Println("Hello, I am the secret agent", s.person.fname, s.person.lname)
	fmt.Println("My licence to kill is", s.licenceKill)
}

func saySomething(h human) {
	h.speak()
}
