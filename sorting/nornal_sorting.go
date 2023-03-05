package sorting

import (
	"fmt"
	"sort"
)

func Test_sort_nornal_sorting() {
	strs := []string{"b", "a", "c"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{4, 1, 2}
	sort.Ints(ints)
	fmt.Println(ints)
}
