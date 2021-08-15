# short lowercase module and package names
# underscores if needed, but discouraged for packages
# package.my_module ie package\my_module.py

from typing import TypeVar

# CapWords for Type Variable names, and preferably short
VT_co = TypeVar("VT_co", covariant=True)
KT_contra = TypeVar("KT_contra", contravariant=True)

# constants: module level and all caps underscores
EULERS_NUMBER = 3
PI = 3


# classes: CapWords
class CustomClass:
    def __init__(self, a, b, c):
        # instance variable: same as function
        self.a = a
        # non public instance variable: same as function but suffix _
        self._b = b
        # name mangled instance variable: same as function but suffix __
        self.__c = c

    # instance method: same as function
    # instance method first arg: self
    def instance_method(self):
        pass

    # non public instance method: same as function but suffix _
    def _non_public_instance_method(self):
        pass

    # name mangled instance method: same as function but suffix __
    def __mangled_instance_method(self):
        pass

    # class method first arg: cls
    @classmethod
    def class_method(cls):
        pass


# classes used as callables: same as functions
class callable_class:
    def __call__(self):
        pass


# error exceptions: also CapWords, but Error suffix
class CustomError(Exception):
    pass


# functions: lowercase with underscores
def my_function():
    pass


def main():
    # variable: same as functions
    my_variable = None


if __name__ == "__main__":
    main()
