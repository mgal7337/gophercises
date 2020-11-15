package main

import (
	"fmt"
	"strings"
)

func index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func include(vs []string, t string) bool {
	return index(vs, t) >= 0
}

func any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func all(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func mapp(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func collectionFunctions() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(index(strs, "pear"))
	fmt.Println(include(strs, "grape"))

	fmt.Println(any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(all(strs, func(v string) bool {
		return strings.Contains(v, "p")
	}))

	fmt.Println(filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(mapp(strs, strings.ToUpper))
}
