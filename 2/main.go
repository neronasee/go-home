package main

import (
	"fmt"
	"homework/2/morse"
	"strings"
)

func min(values ...int) int {
	min := values[0]

	for _, val := range values {
		if min > val {
			min = val
		}
	}
	return min
}
func max(values ...int) int {
	max := values[0]

	for _, val := range values {
		if max < val {
			max = val
		}
	}
	return max
}

type JustPanic struct{}

// type IFeelSomethingWrong struct{}
// type IFeelSomethingWrongTheLast struct{}

func checkPanic(word string) {
	defer func() {
		switch p := recover(); p {
		case nil:
			fmt.Println(word)
		case JustPanic{}:
			fmt.Println("JustPanic")
		default:
			panic(p)
		}
	}()

	if word == "panic" {
		panic(JustPanic{})
	}
}

func panicTest() {
	text := `I feel something's wrong
But the panic is on
I feel something's wrong
But the panic is on
But the panic is on
But the panic is on
I feel something's wrong
But the panic is on
I feel something's wrong
But the panic is on
`

	for _, word := range strings.Split(text, " ") {
		checkPanic(word)
	}
}

func main() {
	// 1
	fmt.Println(min(1, 5, 3))
	fmt.Println(max(1, 5, 3))

	// 2
	rawCode := "--. --- / .-- .- ... / -.. . ... .. --. -. . -.. / .- - / --. --- --- --. .-.. . / - --- / .. -- .--. .-. --- ...- . / .--. .-. --- --. .-. .- -- -- .. -. --. / .--. .-. --- -.. ..- -.-. - .. ...- .. - -.--"
	result, err := morse.Translate(morse.Code(rawCode))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	// 3
	panicTest()
}
