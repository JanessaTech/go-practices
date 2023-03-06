package handlingerror

import "fmt"

func setPanic() {
	panic("a problem")
}
func Test_recover() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from :", r)
		}
	}()
	setPanic()
	fmt.Println("after panic")
}
