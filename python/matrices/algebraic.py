def process(matrix, p):
    result_matrix = []

    for y in range(0, len(matrix)):
        row = []

        for x in range(0, len(matrix[y])):
            row.append(p(y, x, matrix[y][x]))

        result_matrix.append(row)

    return result_matrix


def div(matrix, v):
    return process(matrix, lambda y, x, c: c / v)


def mul(matrix, v):
    return process(matrix, lambda y, x, c: c * v)