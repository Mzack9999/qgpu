package vector

import "math/cmplx"

// A vector represents a n-qubit unit vector in hilbert space
type Vector []complex128

// New state vector initializes a unit vector with given complex bits
// (quantum state)
func New(zBits ...complex128) Vector {
	v := make(Vector, len(zBits))
	copy(v, zBits)
	return v
}

// NewWithSize initialize a new unit vector with zero-ed complex bits
func NewWithSize(n int) Vector {
	return make(Vector, n)
}

// Size returns the number of bits of the unit vector
func (v Vector) Size() int {
	return len(v)
}

// Tensor Product |vw>
// cartesian product of two unit vector
// (produces a column vector)
func (v Vector) TensorProduct(w Vector) Vector {
	tp := make(Vector, 0, v.Size()*w.Size())
	for i := 0; i < v.Size(); i++ {
		for j := 0; j < w.Size(); j++ {
			tp = append(tp, v[i]*w[j])
		}
	}
	return tp
}

// Conjugate/Hermitian Transpose of vector v
// for each item: a + ib => a - ib
func (v Vector) ConjugateTranspose() Vector {
	ct := make(Vector, 0, v.Size())
	for i := 0; i < v.Size(); i++ {
		ct = append(ct, cmplx.Conj(v[i]))
	}
	return ct
}

// Outer Product |vw|
// matrix of size len(n) x len(w)
// (projection)
func (v Vector) OuterProduct(w Vector) [][]complex128 {
	wConjugated := w.ConjugateTranspose()
	op := make([][]complex128, 0, v.Size())
	for i := 0; i < v.Size(); i++ {
		opRow := make([]complex128, 0, w.Size())
		for j := 0; j < w.Size(); j++ {
			opRow = append(opRow, v[i]*wConjugated[j])
		}
		op = append(op, opRow)
	}
	return op
}

// Inner Product <vw>
// single complex128
// (overlap)
func (v Vector) InnerProduct(w Vector) complex128 {
	wConjugated := w.ConjugateTranspose()
	var ip complex128
	for i := 0; i < len(v); i++ {
		ip += v[i] * wConjugated[i]
	}
	return ip
}
