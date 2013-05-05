package main

import "mat"
//import "math"
import "fmt"

func main() {
	a := mat.NewMat2f([]float64{9, 2, 3, 4, 5, 6, 7, 8, 9}, 3, 3)
	fmt.Println()
	fmt.Println(a)
	fmt.Println("#0", a.Tran())
	fmt.Println("#1", a.Inv())
	fmt.Println("#2", a.PInv())
}