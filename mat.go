package mat

type num int
type vec []float64
type Mat2 struct {
	r num 
	c num
	v vec
}

type Mat interface {
	Transpose()
	Inverse()   error
	Mult(* Mat) (* Mat, error)
	Add(* Mat)  (* Mat, error)
}

func NewMat2(v vec, r, c num) *Mat2 {
	return &Mat2{
		r: r,
		c: c,
		v: v,
	}
}

func (m *Mat2) Copy() *Mat2 {
	newv := make(vec, m.r * m.c)
	copy(newv, m.v)
	return NewMat2(newv, m.r, m.c)
}

func Identity(sq num) *Mat2 {
	m := make(vec, sq * sq)
	for i := num(0); i < sq * sq; i += (sq + 1)  {
		m[i] = 1.0
	}
	return &Mat2{v: m, r: sq, c: sq}
}

func (m *Mat2) Set(v float64, r, c num) {
	m.v[m.idx(r, c)] = v
}

func (m *Mat2) Size() (num, num) {
	return m.r, m.c
}

func (m *Mat2) Transpose() {
	out := make(vec, m.r * m.c)
	var cr, cc num = 0, 0	
	for oi, _ := range m.v {
		if cc > (m.c * m.r - m.c) {
			cc = 0
			cr++
		}
		ni := cc + cr
		out[ni] = m.v[oi]
		cc += m.r
	}
	m.v, m.r, m.c = out, m.c, m.r
}

func (a *Mat2) Inverse() error {
	return nil
}

func (a *Mat2) Mult(b *Mat2) (*Mat2, error) {
	if a.c != b.r {
		return nil, nil
	}
	out := Mat2{
		r: a.r,
		c: b.c,
		v: make(vec, a.r * b.c),
	}
	for r, i := num(0), num(0); r < a.c * a.r; r += a.c {
		row := a.row(r)
		for j := num(0); j < b.c; j, i = j+1, i+1 {
			column := b.column(j)
			for k, v := range column {
				out.v[i] += row[k] * v
			}
		}
	}
	return &out, nil
}

// @private
func (m *Mat2) row(i num) vec {
	_, ic := m.ids(i)
	return m.v[(i - ic):(i - ic + m.c)]
}

// @private
func (m *Mat2) column(i num) vec {
	col    := make(vec, m.r)
	ir, ic := m.ids(i)
	for i, _ := range col {
		col[i] = m.v[m.idx(ir, ic)]
		ir++
	}
	return col
}

// @private
func (m *Mat2) ids(i num) (num, num) {
	return (i - i % m.c) / m.c, i % m.c
}

// @private
func (m *Mat2) idx(r, c num) num {
	return m.c * max(r - 1, 0) + c
}

// @private
func max(l, r num) num {
	if l < r {
		return r
	}
	return l
}
