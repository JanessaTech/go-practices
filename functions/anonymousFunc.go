package functions

func case1() {
	f := func(input string) {
		println("input:", input)
	}
	f("hello world")
}

func case2() {
	f := func(int) { println("say case2") }
	f(3)
}

func case3() {
	f := func(int, int, bool) { println("say case3") }
	f(1, 2, true)

}

func TestAnonymousFunc() {
	//case1()
	//case2()
	case3()
}
