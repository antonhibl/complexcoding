package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type quantum_gate struct {
	r1c1 complex128
	r1c2 complex128
	r2c1 complex128
	r2c2 complex128
}

func complex_add(x complex128, y complex128) (result complex128) {
	result = x + y
	return result
}

func complex_subtract(x complex128, y complex128) (result complex128) {
	result = x - y
	return result
}

func complex_multiply(x complex128, y complex128) (result complex128) {
	result = x * y
	return result
}

func complex_divide(x complex128, y complex128) (result complex128) {
	result = x / y
	return result
}

func matrix_maker(r1c1 complex128, r1c2 complex128, r2c1 complex128, r2c2 complex128) (result quantum_gate) {
	result = quantum_gate{
		r1c1: r1c1,
		r1c2: r1c2,
		r2c1: r2c1,
		r2c2: r2c2,
	}
	return result
}

func matrix_display(matrix quantum_gate) {
	fmt.Print("[ ", matrix.r1c1)
	fmt.Print(" ", matrix.r1c2)
	fmt.Print(" ]\n")
	fmt.Print("[ ", matrix.r2c1)
	fmt.Print(" ", matrix.r2c2)
	fmt.Print(" ]\n\n")
}

func matrix_add(matrix1 quantum_gate, matrix2 quantum_gate) (result quantum_gate) {
	r1c1 := complex_add(matrix1.r1c1, matrix2.r1c1)
	r1c2 := complex_add(matrix1.r1c2, matrix2.r1c2)
	r2c1 := complex_add(matrix1.r2c1, matrix2.r2c1)
	r2c2 := complex_add(matrix1.r2c2, matrix2.r2c2)

	result = quantum_gate{
		r1c1: r1c1,
		r1c2: r1c2,
		r2c1: r2c1,
		r2c2: r2c2,
	}

	return result
}

func matrix_subtract(matrix1 quantum_gate, matrix2 quantum_gate) (result quantum_gate) {
	r1c1 := complex_subtract(matrix1.r1c1, matrix2.r1c1)
	r1c2 := complex_subtract(matrix1.r1c2, matrix2.r1c2)
	r2c1 := complex_subtract(matrix1.r2c1, matrix2.r2c1)
	r2c2 := complex_subtract(matrix1.r2c2, matrix2.r2c2)

	result = quantum_gate{
		r1c1: r1c1,
		r1c2: r1c2,
		r2c1: r2c1,
		r2c2: r2c2,
	}

	return result
}

func matrix_multiply(matrix1 quantum_gate, matrix2 quantum_gate) (result quantum_gate) {
	r1c1 := complex_multiply(matrix1.r1c1, matrix2.r1c1)
	r1c2 := complex_multiply(matrix1.r1c2, matrix2.r1c2)
	r2c1 := complex_multiply(matrix1.r2c1, matrix2.r2c1)
	r2c2 := complex_multiply(matrix1.r2c2, matrix2.r2c2)

	result = quantum_gate{
		r1c1: r1c1,
		r1c2: r1c2,
		r2c1: r2c1,
		r2c2: r2c2,
	}

	return result
}

func matrix_divide(matrix1 quantum_gate, matrix2 quantum_gate) (result quantum_gate) {
	r1c1 := complex_divide(matrix1.r1c1, matrix2.r1c1)
	r1c2 := complex_divide(matrix1.r1c2, matrix2.r1c2)
	r2c1 := complex_divide(matrix1.r2c1, matrix2.r2c1)
	r2c2 := complex_divide(matrix1.r2c2, matrix2.r2c2)

	result = quantum_gate{
		r1c1: r1c1,
		r1c2: r1c2,
		r2c1: r2c1,
		r2c2: r2c2,
	}

	return result
}

func main() {
	// inputs for complex X
	// Real
	x1, _ := strconv.ParseFloat(os.Args[1], 128)
	// Imaginary
	x2, _ := strconv.ParseFloat(os.Args[2], 128)
	// complex = real + imaginary
	var x complex128
	x = complex(float64(x1), float64(x2))

	// inputs for complex Y
	// Real
	y1, _ := strconv.ParseFloat(os.Args[3], 128)
	// Imaginary
	y2, _ := strconv.ParseFloat(os.Args[4], 128)
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
	fmt.Print("\n\n")

	// construct matrices
	q := matrix_maker(x, -x, y, -y)
	v := matrix_maker(-y, x, -x, x)

	// display both matrices
	fmt.Print("Q =\n")
	matrix_display(q)
	fmt.Print("V =\n")
	matrix_display(v)

	fmt.Print("Q + V = \n")
	added_matrixes := matrix_add(q, v)
	matrix_display(added_matrixes)

	fmt.Print("Q - V = \n")
	subtracted_matrixes := matrix_subtract(q, v)
	matrix_display(subtracted_matrixes)

	fmt.Print("Q * V = \n")
	multiplied_matrixes := matrix_multiply(q, v)
	matrix_display(multiplied_matrixes)

	fmt.Print("Q / V = \n")
	divided_matrixes := matrix_divide(q, v)
	matrix_display(divided_matrixes)

	quant := complex(1/math.Sqrt(2), 1/math.Sqrt(2))
	h := matrix_maker(quant, quant, quant, -quant)

	fmt.Print("The Hadamard Gate: \n")
	matrix_display(h)

	SH := matrix_multiply(h, h)
	fmt.Print("H * H =\n")
	matrix_display(SH)
}
