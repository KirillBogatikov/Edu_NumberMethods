class Member(object):
    def __init__(self, base, exp):
        self.__base = base
        self.__exp = exp

    def __str__(self):
        return "%.3f" % pow(self.__base, self.__exp)


class VariableMember(Member):
    def __init__(self, base, exp, name = "x"):
        super().__init__(base, exp)
        self.__name = name

    def get(self, x):
        return pow(x, self.__exp) * self.__base

    def name(self):
        return self.__name

    def __str__(self):
        return "%.3f%s^%d" % (self.__base, self.__name, self.__exp)


class Equation(object):
    def __init__(self, members):
        self.__cached_variables_count = None
        self.__members = members

    def variables_count(self):
        if self.__cached_variables_count is None:
            c = 0
            names = []
            for m in self.__members:
                try:
                    names.index(m.name())
                except ValueError:
                    names.append(m.name)
                    c += 1

            self.__cached_variables_count = c

        return self.__cached_variables_count



