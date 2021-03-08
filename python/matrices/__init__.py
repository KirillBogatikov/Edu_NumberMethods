from matrices.basic import determinant_of, cofactor_matrix_of, transposed_of
from matrices.reverse import reversed_of


class Matrix(object):

    def __init__(self, matrix):
        self.__cached_determinant = None
        self.__cached_cofactor = None
        self.__cached_transpose = None
        self.__cached_reverse = None
        self.__rows = matrix

    def determinant(self):
        if self.__cached_determinant is None:
            self.__cached_determinant = determinant_of(self.__rows)

        return self.__cached_determinant

    def cofactor(self):
        if self.__cached_cofactor is None:
            self.__cached_cofactor = Matrix(cofactor_matrix_of(self.__rows))

        return self.__cached_cofactor

    def transpose(self):
        if self.__cached_transpose is None:
            self.__cached_transpose = Matrix(transposed_of(self.__rows))

        return self.__cached_transpose

    def reverse(self):
        if self.__cached_reverse is None:
            self.__cached_reverse = Matrix(reversed_of(self.__rows))

        return self.__cached_reverse

    def matrix(self):
        return self.matrix

    def __str__(self):
        rows = []

        for y in range(0, len(self.__rows)):
            row = []

            for x in range(0, len(self.__rows[y])):
                row.append(str(self.__rows[y][x]))

            rows.append(', '.join(row))

        return '\n'.join(rows)