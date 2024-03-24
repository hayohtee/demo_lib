// package demo_lib provides basic functions to perform
// addition and subtraction on two numbers.
package demo_lib

import "golang.org/x/exp/constraints"

// Add is a function that takes in two numbers values and returns their sum.
// It follows the rules specified by [Mathisfun]
//
// [Mathisfun]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}

// Number is a type element for all integers and floats
type Number interface {
	constraints.Integer | constraints.Float
}
