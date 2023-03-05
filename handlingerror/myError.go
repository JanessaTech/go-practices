package handlingerror

import (
	"errors"
	"fmt"
)

type customError struct {
	code int
	msg  string
}

func (e *customError) Error() string { // notice here we use *obj
	return fmt.Sprintf("error[%d] : %s", e.code, e.msg)
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cannot work with 42")
	}
	return 1, nil
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &customError{code: -1, msg: "cannot work with 42"} // notice that here we use &customError
	}
	return 1, nil
}

func Test() {
	for _, i := range []int{7, 42} { // use _ to hold the place where we don't want to know the exact value
		if r, e := f1(i); e != nil { // we could use this way  instead of condition1 && condition2 in if condition
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed ", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*customError); ok { // pay attention on how to check if a variale(e) belongs to a type(customError)   using the sytax: v.(type)
		fmt.Println(ae.code)
		fmt.Println(ae.msg)
	}

}
