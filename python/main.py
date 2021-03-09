from equation import Member, VariableMember, Equation
from gauss import *
from matrices import Matrix
from seidel.basic import prepare, solve
from util import check_equation_keys, calc_equation


def gauss_main():
    matrix = [
        [0.10, 12, -0.13, 0.10],
        [0.12, 0.71, 0.15, 0.26],
        [-0.13, 0.15, 0.63, 0.38]
    ]

    x = gauss(3, matrix)
    print(x)


def matrix_main():
    matrix = Matrix([
        [1, 2, 3, 4],
        [2, 3, 4, 1],
        [3, 4, 1, 2],
        [4, 1, 2, 3]
    ])

    r = matrix.reverse()
    print(r)


def seidel_main():
    e = [
        [-0.13, 0.15, 0.63, 0.38],
        [0.10, 12.00, -0.13, 0.10],
        [0.12, 0.71, 0.15, 0.26]
    ]

    s = solve(e, 0.01)
    if s is None:
        print("No solution")
    else:
        print(s)
        print(calc_equation(e[0], s))


def equation_test():
    e = Equation([VariableMember(5, 2, "x1"), VariableMember(3, 3, "x2"), Member(5)])
    r = e.get([VariableMember(1, name="x1"), VariableMember(2, name="x2")])
    print(r)


if __name__ == '__main__':
    equation_test()
