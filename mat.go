package mat

type Mat interface {
	Tran() *Mat
	Inv() *Mat
	PInv() *Mat
	Mult(*Mat) *Mat
	Add(*Mat) *Mat
}
