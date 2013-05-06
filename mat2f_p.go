package mat

// @private
func (m *Mat2f) row(i num) vec {
	_, ic := m.ids(i)
	//	c := make(vec, m.c)
	//	copy(c, m.v[(i-ic):(i-ic+m.c)])
	return m.v[(i - ic):(i - ic + m.c)]
}

// @private
func (m *Mat2f) column(i num) vec {
	col := make(vec, m.r)
	ir, ic := m.ids(i)
	for i, _ := range col {
		col[i] = m.v[m.idx(ir, ic)]
		ir++
	}
	return col
}

// @private
func (m *Mat2f) ids(i num) (num, num) {
	return (i - i%m.c) / m.c, i % m.c
}

// @private
func (m *Mat2f) idx(r, c num) num {
	return m.c*r + c
}

// @private 
func (m *Mat2f) connIdent() *Mat2f {
	var merged vec
	for i := num(0); i < m.r; i++ {
		tail := make(vec, m.r)
		tail[i] = 1.0
		row := make(vec, m.r)
		copy(row, m.row(m.idx(i, 0)))
		merged = append(merged, append(row, tail...)...)
	}
	return &Mat2f{v: merged, r: m.r, c: 2 * m.c}
}

// @private
func (m *Mat2f) splitOffLeft() *Mat2f {
	var splitted vec
	for i := num(0); i < m.r; i++ {
		splitted = append(splitted, m.v[m.idx(i, m.r):m.idx(i, m.c)]...)
	}
	return &Mat2f{v: splitted, r: m.r, c: m.r}
}
