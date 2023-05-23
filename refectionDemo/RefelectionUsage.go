package refectiondemo

import (
	"fmt"
	"reflect"
)

type Info struct {
	Person struct {
		Name   string `json:"name" binding:"required"`
		Age    int    `json:"age" binding:"required"`
		IsMale bool   `json:"isMale" binding:"required"`
	} `json:"person"`
}

type Person struct {
	Name   string `json:"name" binding:"required"`
	Age    int    `json:"age" binding:"required"`
	IsMale bool   `json:"isMale" binding:"required"`
}

type CustomStr string

func showInfoDetails(data interface{}) {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		e := t.Elem()
		f, _ := e.FieldByName("Age")
		ageTag := f.Tag.Get("json")
		fmt.Println("Age's tag = ", ageTag)
	}
}

func showPersonDetails(data interface{}) {
	t := reflect.TypeOf(data)
	k := t.Kind()
	v := reflect.ValueOf(data)
	fmt.Println("type=", t)  // refectiondemo.Person
	fmt.Println("kind=", k)  // struct
	fmt.Println("value=", v) // {Jane 10 false}
	if k == reflect.Struct {
		fields := reflect.ValueOf(data)
		fmt.Println("Number of fields for", t, " = ", fields.NumField())
		for i := 0; i < fields.NumField(); i++ {
			ft := reflect.TypeOf(fields.Field(i))
			fk := fields.Field(i).Kind()
			fv := fields.Field(i)
			fmt.Println(" field type=", ft)
			fmt.Println("field kind=", fk)
			fmt.Println("field value=", fv)

		}

	}
}

// show how to use relect.Type.Elem()
func readTag() {
	person := &Person{
		Name:   "Jane",
		Age:    10,
		IsMale: false,
	}

	tm := reflect.TypeOf(person).Elem()
	f, ok := tm.FieldByName("Name")
	if ok {
		tagName, _ := f.Tag.Lookup("json")
		fmt.Println("Name's tag=", tagName)
	}
}

func showStructFieldNames(data interface{}) {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		fName := t.Field(i).Name
		fmt.Println(fName)
	}
}

func showCustomStrDetails(data interface{}) {
	t := reflect.TypeOf(data)
	k := t.Kind()
	v := reflect.ValueOf(data)
	fmt.Println("type=", t)  // refectiondemo.CustomStr
	fmt.Println("kind=", k)  // string
	fmt.Println("value=", v) // hello
}

func Demo() {
	person := Person{
		Name:   "Jane",
		Age:    10,
		IsMale: false,
	}
	fmt.Println("Struct Person has the following field name:--------")
	showStructFieldNames(person)
	fmt.Println()
	fmt.Println("showPersonDetails")
	showPersonDetails(person)
	fmt.Println()
	fmt.Println("showCustomStrDetails")
	customStr := CustomStr("hello")
	showCustomStrDetails(customStr)
	fmt.Println()
	fmt.Println("readTag")
	readTag()
	info := Info{
		Person: Person{
			Name:   "Jane",
			Age:    10,
			IsMale: false,
		},
	}
	fmt.Println()
	showInfoDetails(&info.Person) // we see that Person is Ptr type, not Struct
}
