package main

import "mat"
import "fmt"

func main() {
	m := mat.Identity(3)
	m.Set(2, 0, 2)
	c := mat.NewMat2([]float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0}, 2, 3)
	fmt.Println(c)
	c.Transpose()
	fmt.Println(c)
	fmt.Println()
	d, _ := m.Mult(c)
	fmt.Printf("%v\n%v\n%v\n", m, c, d)
}

// @private
func max(l, r int) int {
	if l < r {
		return r
	}
	return l
}
