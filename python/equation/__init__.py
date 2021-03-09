class Member(object):
    def __init__(self, base, exp=1):
        self.__base = base
        self.__exp = exp

    def base(self):
        return self.__base

    def exp(self):
        return self.__exp

    def get(self):
        return pow(self.__base, self.__exp)

    def neg(self):
        self.__base *= -1

    def __str__(self):
        return "%.3f" % pow(self.__base, self.__exp)


class VariableMember(Member):
    def __init__(self, base, exp = 1, name="x"):
        super().__init__(base, exp)
        self.__name = name

    def calc(self, x):
        return pow(x, self.exp()) * self.base()

    def name(self):
        return self.__name

    def __str__(self):
        return "%.3f%s^%d" % (self.base(), self.__name, self.exp())


class Equation(object):
    def __init__(self, members):
        self.__cached_variables = None
        self.__members = {}

        r = 0.0
        for m in members:
            if isinstance(m, VariableMember):
                try:
                    _ = self.__members[m.name()]
                except KeyError:
                    self.__members[m.name()] = m
            else:
                r += m.get()
            continue

        self.__result = Member(r * -1)

    def variables(self):
        if self.__cached_variables is None:
            names = {}
            for i in self.__members:
                member = self.__members[i]

                try:
                    names[member.name()] += 1
                except KeyError:
                    names[member.name()] = 0

            self.__cached_variables = names
        return self.__cached_variables

    def get(self, keys):
        v = self.variables()
        if len(keys) != len(v.keys()):
            return None

        m = self.__members.copy()

        result = 0.0
        for k in keys:
            print(k)
            result += m[k.name()].calc(k.calc(1))
            m.pop(k.name())

        for i in m:
            result += m[i].get()

        return result
