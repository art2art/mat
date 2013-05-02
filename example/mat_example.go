package main

import "mat"
import "fmt"

func main() {
	m := mat.NewMat2f([]float64{4, 1, 2, 3, 0, 5, 0, 0, 0}, 3, 3)
	fmt.Println("transposed", m.Transpose())
	fmt.Println("mult", m.Mult(m))
	fmt.Println("inverted", m.Inverse())
	fmt.Println("pseudoinverse", m.PseudoInverse())
}