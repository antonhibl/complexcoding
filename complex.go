package main

import (
	"fmt"
	"os"
	"strconv"
)

func complex_add(x complex128, y complex128) (result complex128) {
	result = x + y
	return result
}

func complex_multiply(x complex128, y complex128) (result complex128) {
	result = x * y
	return result
}

func main() {
	// inputs for complex X
	// Real
	x1, _ := strconv.Atoi(os.Args[1])
	// Imaginary
	x2, _ := strconv.Atoi(os.Args[2])
	// complex = real + imaginary
	var x complex128
	x = complex(float64(x1), float64(x2))

	// inputs for complex Y
	// Real
	y1, _ := strconv.Atoi(os.Args[3])
	// Imaginary
	y2, _ := strconv.Atoi(os.Args[4])
	// complex = real + imaginary
	var y complex128
	y = complex(float64(y1), float64(y2))

	// Add the complex numbers
	var complex_sum complex128
	complex_sum = complex_add(x, y)

	// Calculate the product
	var complex_product complex128
	complex_product = complex_multiply(x, y)

	// print X
	fmt.Print("\nX = ", x)
	// print Y
	fmt.Print("\nY = ", y)
	// Print the sum
	fmt.Print("\nX + Y = ", complex_sum)
	// print the product
	fmt.Print("\nX * Y = ", complex_product)
	// print divisor operations
	// x/y
	fmt.Print("\nX / Y = ", x/y)
	// y/x
	fmt.Print("\nY / X = ", y/x)
}
