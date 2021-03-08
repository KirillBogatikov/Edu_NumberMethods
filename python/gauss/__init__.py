def gauss(n, matrix):
    x = []

    for i in range(0, n):
        x.append(0)

    k = 1
    while k < n:
        j = k
        while j < n:
            m = matrix[j][k - 1] / matrix[k - 1][k - 1]

            i = 0
            while i < n + 1:
                matrix[j][i] = matrix[j][i] - m * matrix[k - 1][i]
                i += 1

            j += 1

        i = n - 1
        while i >= 0:
            x[i] = matrix[i][n] / matrix[i][i]

            c = n - 1
            while c > i:
                x[i] = x[i] - matrix[i][c] * x[c] / matrix[i][i]
                c -= 1

            i -= 1

        k += 1

    return x

