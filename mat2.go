package mat

import "fmt"

type Mat2 struct {
	r, c int
	v    []interface{}
}

func NewMat2(r, c int, v []interface{}) (*Mat2, error) {
	if len(v) != r*c {
		return nil, fmt.Errorf("[NewMat2] wrong indecies (r:%d, c:%d)", r, c)
	}
	return &Mat2{r: r, c: c, v: v}, nil
}

func (m *Mat2) Set(idx int, v interface{}) error {
	if idx >= m.Length() {
		return fmt.Errorf("[Cell] wrong index (idx:%d)", idx)
	}
	m.v[idx] = v
	return nil
}

func (m *Mat2) Cell(idx int) (interface{}, error) {
	if idx >= m.Length() {
		return nil, fmt.Errorf("[Cell] wrong index (i:%d)", idx)
	}
	return m.v[idx], nil
}

func (m *Mat2) Length() int {
	return m.r * m.c
}

func (m *Mat2) Size() (int, int) {
	return m.r, m.c
}

func (m *Mat2) Ids(i int) (int, int) {
	return (i - i % m.c) / m.c, i % m.c
}

func (m *Mat2) Idx(r, c int) int {
	return m.c*r + c
}

func (m *Mat2) Row(i int) []interface{} {
	_, ic := m.Ids(i)
	return m.v[(i - ic):(i - ic + m.c)]
}

func (m *Mat2) Column(i int) []interface{} {
	column := make([]interface{}, m.r)
	ir, ic := m.Ids(i)
	for i, _ := range column {
		column[i] = m.v[m.Idx(ir, ic)]
		ir++
	}
	return column
}
