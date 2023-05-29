package mapdemo

import "fmt"

func MapDemo() {
	dict := map[string]int{"aa": 1, "bb": 2, "cc": 3}
	value, ok := dict["dd"]
	if ok {
		fmt.Println("value:", value)
	} else {
		fmt.Println("key not found")
	}
}
