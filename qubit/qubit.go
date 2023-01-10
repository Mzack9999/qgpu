package qubit

import (
	"fmt"

	"github.com/Mzack9999/qgpu/vector"
)

type Qubit struct {
	vector vector.Vector
}

func New(z ...complex128) *Qubit {
	qubit := &Qubit{
		vector: vector.NewWithBits(z...),
	}
	return qubit
}

// prints qubit state vector
func (qubit *Qubit) String() string {
	return fmt.Sprint(qubit.vector)
}
