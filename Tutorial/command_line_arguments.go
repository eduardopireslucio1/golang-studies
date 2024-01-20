package main

import (
	"fmt"
	"os"
	"strings"
)

func methodOne() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func methodTwo() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func simpleSolution() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func dontFormat() {
	fmt.Println(os.Args[1:])
}

func printIndexAndValue() {
	for index, value := range os.Args[1:] {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
}

func main() {
	methodOne()
	methodTwo()
	simpleSolution()
	dontFormat()
	printIndexAndValue()
}