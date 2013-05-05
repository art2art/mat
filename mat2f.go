package mat

import "fmt"

type num int
type vec []float64
type Mat2f struct {
	r num
	c num
	v vec
}

func NewMat2f(v vec, r, c num) *Mat2f {
	if len(v) != int(r*c) {
		return nil
	}
	return &Mat2f{r: r, c: c, v: v}
}

func (m *Mat2f) Copy() *Mat2f {
	c := make(vec, m.r*m.c)
	copy(c, m.v)
	return NewMat2f(c, m.r, m.c)
}

func (m *Mat2f) Set(idx num, v float64) error {
	if idx >= m.Length() {
		return fmt.Errorf("[Set] wrong index (idx:%d)", idx)
	}
	m.v[idx] = v
	return nil
}

func (m *Mat2f) Size() (num, num) {
	return m.r, m.c
}

func (m *Mat2f) Length() num {
	return m.r * m.c
}

func (m *Mat2f) Tran() *Mat2f {
	out := &Mat2f{r: m.r, c: m.c}
	if m.r == 1 || m.c == 1 {
		out.v = make(vec, m.r*m.c)
		copy(out.v, m.v)
		return out
	}
	for i := num(0); i < m.c; i++ {
		c := m.column(i)
		out.v = append(out.v, c...)
	}
	return out
}

func (m *Mat2f) Inv() *Mat2f {
	if m.r != m.c {
		return nil
	}
	ext := m.connIdent()
	fcheck := func(i num) bool { return i < m.r }
	bcheck := func(i num) bool { return i >= 0 }
	gaussStep := func(start, step num, bound func(num) bool) bool {
		for i := start; bound(i); i += step {
			// #1
			if ext.v[ext.idx(i, i)] == 0 {
				mark := true
				for j := i + step; bound(j); j += step {
					if ida := ext.idx(j, i); ext.v[ida] != 0 {
						mark = false
						for jc, a := range ext.row(ida) {
							ext.v[ext.idx(i, num(jc))] += a
						}
					}
				}
				if mark {
					return false
				}
			}
			// #2
			for j := i + step; bound(j); j += step {
				k := ext.v[ext.idx(j, i)] / ext.v[ext.idx(i, i)]
				for jc, a := range ext.row(ext.idx(i, i)) {
					if num(jc) == i {
						ext.v[ext.idx(j, i)] = 0
					} else {
						ext.v[ext.idx(j, num(jc))] -= k * a
					}
				}
			}
		}
		return true
	}

	if !gaussStep(0, 1, fcheck) {
		return nil
	}
	if !gaussStep(m.r-1, -1, bcheck) {
		return nil
	}
	result := make(vec, m.r * m.r)
	for i := num(0); i < ext.r; i++ {
		ii := ext.idx(i, i)
		row := ext.row(ii)
		for j := ext.r; j < ext.c; j++ {
			result[m.idx(i, j - ext.r)] = row[j] / ext.v[ii]
		}
	}
	return &Mat2f{v: result, r: ext.r, c: ext.r}
}

// (X' * X)^-1 * X'
func (m *Mat2f) PInv() *Mat2f {
	xT := m.Tran()
	xM := xT.Mult(m)
	if xM == nil {
		return nil
	}
	xI := xM.Inv()
	if xI == nil {
		return nil
	}
	return xI.Mult(xT)
}

func (a *Mat2f) Add(b *Mat2f) *Mat2f {
	if a.r != b.r || a.c != b.c {
		return nil
	}
	result := NewMat2f(make(vec, a.r*a.c), a.r, a.c)
	for i, v := range a.v {
		result.v[i] = v + b.v[i]
	}
	return result
}

func (a *Mat2f) Mult(b *Mat2f) *Mat2f {
	if a.c != b.r {
		return nil
	}
	out := Mat2f{
		r: a.r,
		c: b.c,
		v: make(vec, a.r*b.c),
	}
	for r, i := num(0), num(0); r < a.c*a.r; r += a.c {
		row := a.row(r)
		for c := num(0); c < b.c; c++ {
			column := b.column(c)
			for k, v := range column {
				out.v[i] += row[k] * v
			}
			i++
		}
	}
	return &out
}

// Identity matrix
func I(sq num) *Mat2f {
	m := make(vec, sq*sq)
	for i := num(0); i < sq*sq; i += (sq + 1) {
		m[i] = 1.0
	}
	return &Mat2f{v: m, r: sq, c: sq}
}
