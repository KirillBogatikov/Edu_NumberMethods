from util import max_row_of


def prepare(matrix):
    size = len(matrix)

    copy = []
    for y in range(0, size):
        row = []

        for x in range(0, size + 1):
            row.append(matrix[y][x])

        copy.append(row)

    prepared = []
    for i in range(0, size):
        max_row = max_row_of(copy, i, True)
        copy.remove(max_row)
        prepared.append(max_row)

    return prepared


def solve(matrix, accuracy):
    size = len(matrix)
    system = prepare(matrix)

    p: []
    x = []
    for i in range(0, size):
        # x.append(system[i][size])
        x.append(0)

    for i in range(0, 3):
        p = x.copy()
        x = iteration(system, x, size)
        print(x)

        for j in range(0, size):
            if abs(p[j] - x[j]) < accuracy:
                return x

    return None


def iteration(system, x, size):
    for i in range(0, size):
        r = 0.0
        for j in range(0, size):
            if j == i:
                continue

            r -= system[i][j] * x[j]

        r += system[i][size]
        r /= system[i][i]
        x[i] = r

    return x
