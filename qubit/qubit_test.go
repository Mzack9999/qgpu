package qubit

import (
	"fmt"
	"math"
	"math/cmplx"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuperpositions(t *testing.T) {
	q0 := Zero()
	require.Equal(t, complex128(1), q0.Zero)
	q1 := One()
	require.Equal(t, complex128(1), q1.One)
	// diagonal superposition
	q2 := New(complex(1/math.Sqrt(2), 0), complex(1/math.Sqrt(2), 0))
	// in qubit |alpha^2| + |beta^2| = 1
	require.Equal(t, "(0.9999999999999998+0i)", fmt.Sprint(cmplx.Pow(q2.Zero, 2)+cmplx.Pow(q2.One, 2)))
}
