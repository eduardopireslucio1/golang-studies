package main

import (
	"fmt"
)

/* In type declarations: A type declaration makes it possible to give a name to an
existing type. Since structure types are often long, they are nearly always named. 
A familiar example is the definition of a Point type for a 2-D graphics system:
*/
type Point struct {
	X, Y int
}

var p Point

func main() {
	heads := 0
	tails := 0
	/* Control flow: We covered the two fundamental control-flow statements, if and for, but not the switch statement, which is a multi-way branch. Here's a small example: */
	switch coinflip() {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
		fmt.Println("landed on edge!")
	}
	fmt.Println(heads, tails)
}

/* This for m is cal le d a tagless switch; it’s equivalent to switch true. */
func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}

func coinflip() string {
	return "heads"
}
