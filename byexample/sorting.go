package main

import (
	"fmt"
	"sort"
)

func sorting() {
	strs := []string{"c", "a", "b", "e", "z", "w"}
	sort.Strings(strs)
	fmt.Println("strings:", strs)

	ints := []int{7, 1, 65, 1324, 2, 4, 0, -123, -345, -3}
	sort.Ints(ints)
	fmt.Println("ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted:", s)
}
