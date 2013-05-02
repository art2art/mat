package mat

type Mat2 struct {
	r, c int
	v    []interface{}
}

func NewMat2(v []interface{}, r, c int) (*Mat2, error) {
	if len(v) != r*c {
		return nil, nil
	}
	return &Mat2{r: r, c: c, v: v}, nil
}

func (m *Mat2) Cell(idx int) (interface{}, error) {
	if idx >= m.Size() {
		return nil, nil
	}
	return m.v[idx], nil
}

func (m *Mat2) Size() int {
	return m.r * m.c
}

func (m *Mat2) Scale() (int, int) {
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
	//	c := make(vec, m.c)
	//	copy(c, m.v[(i-ic):(i-ic+m.c)])
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
