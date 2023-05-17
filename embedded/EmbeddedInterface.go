package embedded

import "fmt"

type animal interface {
	breathe()
	walk()
}

type human interface {
	animal
	speak()
}

type employee struct {
	name string
}

func (e employee) breathe() {
	fmt.Println("Employee breathes")
}

func (e employee) walk() {
	fmt.Println("Employee walk")
}

func (e employee) speak() {
	fmt.Println("Employee speaks")
}

func printDetails(h human) {
	h.breathe()
	h.walk()
	h.speak()
}

func EmbedInterface2Interface() {
	e := employee{name: "jane"}
	printDetails(e)
}

func EmbedInterface2Struct() {

}
