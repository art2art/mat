package mat

type Mat interface {
	Transpose()
	Inverse() *Mat
	PseudoInverse() *Mat
	Mult(*Mat) *Mat
	Add(*Mat) *Mat
}
