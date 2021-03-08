def max_row_of(matrix, p, use_abs=False):
    r = 0

    for i in range(1, len(matrix)):
        v = abs(matrix[i][p]) if use_abs else matrix[i][p]

        if v > abs(matrix[r][p]):
            r = i

    return matrix[r]


def calc_equation(equation, keys):
    size = len(equation) - 1
    v = 0.0
    for i in range(0, size):
        v += equation[i] * keys[i]

    return v - equation[size]


def check_equation_keys(equation, keys, accuracy):
    return abs(calc_equation(equation, keys) - equation[len(equation) - 1]) <= accuracy
