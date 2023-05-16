package functions

import "fmt"

type Shape interface {
	draw()
}

type Circle struct{}
type Square struct{}

func (c *Circle) draw() {
	fmt.Println("Drawing a circle ...")
}

func (s *Square) draw() {
	fmt.Println("Drawing a square ...")
}

func testNonEmptyInterface() {
	c := Circle{}
	c.draw()

	s := Square{}
	s.draw()
}

func customPrintln(data interface{}) {
	fmt.Printf("value = %v, type = %T\n", data, data)
}

func testEmptyInterface() {
	customPrintln("a string")
	customPrintln([]int{1, 2, 3})
	customPrintln(true)
}

func TestInterfaces() {
	testNonEmptyInterface()
	testEmptyInterface()
}
