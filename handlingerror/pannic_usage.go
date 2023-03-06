package handlingerror

import "os"

func Test_panic() {
	panic("first error")

	_, err := os.Create("./tmp/file")
	if err != nil {
		panic(err)
	}
}
