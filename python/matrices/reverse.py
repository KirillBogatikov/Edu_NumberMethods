from matrices.algebraic import div
from matrices.basic import determinant_of, cofactor_matrix_of, transposed_of


def reversed_of(matrix):
    d = determinant_of(matrix)
    if d == 0:
        return None

    c = cofactor_matrix_of(matrix)
    t = transposed_of(c)

    return div(t, d)
