package main

import (
	"fmt"
	"sort"
	"strings"
)

func invert(slice []int) []int {
	newSlice := []int{}

	for i := len(slice) - 1; i >= 0; i-- {
		newSlice = append(newSlice, slice[i])
	}

	return newSlice
}

func isAnagram(a string, b string) bool {
	s1 := strings.Split(a, "")
	s2 := strings.Split(b, "")

	sort.Strings(s1)
	sort.Strings(s2)

	return strings.Join(s1, "") == strings.Join(s2, "")
}

func appendString(s []string, val string) []string {
	currLength := len(s)
	currCap := cap(s)

	if currCap >= currLength+1 {
		s = s[:currLength+1]

	} else {
		newCap := currCap * 2
		if currCap == 0 {
			newCap = 1
		}

		newS := make([]string, currLength+1, newCap)
		copy(newS, s)
		s = newS
	}

	s[currLength] = val

	return s
}

func main() {
	fmt.Println("Hello, 世界")

	numbers := []int{1, 2, 3, 4}
	fmt.Println(numbers)
	invNumbers := invert(numbers)
	fmt.Println(invNumbers)

	fmt.Println(isAnagram("abc", "bca"))
	fmt.Println(isAnagram("qaz", "zse"))

	names := []string{}
	fmt.Println(names, "cap", cap(names), "len", len(names))
	names = appendString(names, "Bob")
	fmt.Println(names, "cap", cap(names), "len", len(names))
	names = appendString(names, "Charlie")
	fmt.Println(names, "cap", cap(names), "len", len(names))
	names = appendString(names, "Jhon")
	fmt.Println(names, "cap", cap(names), "len", len(names))
	names = appendString(names, "Roger")
	fmt.Println(names, "cap", cap(names), "len", len(names))
	names = appendString(names, "Ric")
	fmt.Println(names, "cap", cap(names), "len", len(names))

}
