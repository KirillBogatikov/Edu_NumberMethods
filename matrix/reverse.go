package matrix

import (
	"fmt"
)

func (m *Matrix) Reverse() *Matrix {
	determinant := m.Determinant()
	if verboseOut {
		fmt.Printf("Matrix determinant: %d\n", determinant.IntPart())
	}

	cofactor := m.Cofactor()
	if verboseOut {
		fmt.Printf("Matrix cofactor: %s\n", "\n"+cofactor.ToString())
	}

	transposedCofactor := cofactor.Transpose()
	if verboseOut {
		fmt.Printf("Transposed cofactor: %s\n", "\n"+transposedCofactor.ToString())
	}

	return transposedCofactor.Div(determinant)
}
