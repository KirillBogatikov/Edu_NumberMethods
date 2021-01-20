package matrix

import (
	"fmt"
)

func (m *Matrix) Reverse() *Matrix {
	determinant := m.Determinant()
	fmt.Printf("Matrix determinant: %d\n", determinant.IntPart())

	cofactor := m.Cofactor()
	fmt.Printf("Matrix cofactor: %s\n", "\n" + cofactor.ToString())

	transposedCofactor := cofactor.Transpose()
	fmt.Printf("Transposed cofactor: %s\n", "\n" + transposedCofactor.ToString())

	return transposedCofactor.Div(determinant)
}