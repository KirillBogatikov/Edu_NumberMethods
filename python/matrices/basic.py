def minor_of(matrix, row, column):
    minor_matrix = []

    for y in range(0, len(matrix)):
        if y == row:
            continue

        row_array = []
        for x in range(0, len(matrix[y])):
            if x == column:
                continue

            row_array.append(matrix[y][x])

        minor_matrix.append(row_array)

    return minor_matrix


def multiply_coefficient_of(row, column):
    return 1 if (row + column) % 2 == 0 else -1


def cofactor_of(matrix, row, column):
    m = minor_of(matrix, row, column)
    md = determinant_of(m)
    k = multiply_coefficient_of(row, column)

    return md * k


def determinant_of(matrix):
    if len(matrix) == 0:
        return 0

    if len(matrix) != len(matrix[0]):
        return 0

    if len(matrix) == 2:
        return matrix[0][0] * matrix[1][1] - matrix[0][1] * matrix[1][0]

    if len(matrix) == 3:
        return matrix[0][0] * matrix[1][1] * matrix[2][2] + \
               matrix[2][0] * matrix[0][1] * matrix[1][2] + \
               matrix[1][0] * matrix[2][1] * matrix[0][2] - \
               matrix[0][2] * matrix[1][1] * matrix[2][0] - \
               matrix[0][0] * matrix[1][2] * matrix[2][1] - \
               matrix[0][1] * matrix[1][0] * matrix[2][2]

    d = 0.0
    h = len(matrix)
    for i in range(0, h):
        d += matrix[0][i] * cofactor_of(matrix, 0, i)

    return d


def cofactor_matrix_of(matrix):
    result_matrix = []

    for y in range(0, len(matrix)):
        row = []

        for x in range(0, len(matrix[y])):
            row.append(cofactor_of(matrix, y, x))

        result_matrix.append(row)

    return result_matrix


def transposed_of(matrix):
    width = len(matrix[0])
    height = len(matrix)

    result_matrix = []

    for y in range(0, width):
        row = []

        for x in range(0, height):
            row.append(matrix[x][y])

        result_matrix.append(row)

    return result_matrix
