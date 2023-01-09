package qubit

import "fmt"

type Qubit struct {
	Zero complex128
	One  complex128
}

// prints qubit statevector containing zero/one superpositions with Dirac notation
func (qubit *Qubit) String() string {
	return fmt.Sprintf("%v |0⟩ + %v |1⟩", qubit.Zero, qubit.One)
}

// |0>
func Zero() *Qubit {
	return &Qubit{Zero: 1}
}

// |1>
func One() *Qubit {
	return &Qubit{One: 1}
}

// todo: support multi-bit qubits
func New(zero, one complex128) *Qubit {
	return &Qubit{Zero: zero, One: one}
}
