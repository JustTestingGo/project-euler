package main

import (
	"fmt"
)

func sumSquares(squares ...int) int {
	var sum int
	for _, square := range squares {
		sum += square
	}
	return sum
}

func main() {
	var squares []int
	var sum int
	for i := 1; i < 101; i++ {
		squares = append(squares, i*i)
		sum += i
	}
	fmt.Println("The difference between the sum of the squares of the first one hundred natural numbers and the square of the sum is",
		sum*sum-sumSquares(squares...))
}

/*

The sum of the squares of the first ten natural numbers is,
12 + 22 + ... + 102 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)2 = 552 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

*/
