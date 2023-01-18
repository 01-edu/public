from numerical_operations_the_return import *
import sys

sys.path.append('/jail/app/student')


def test_modulo():
    assert modulo(100, 50) == 0
    assert modulo(11, 5) == 1
    assert modulo(27, 20) == 7
    assert modulo(3.5, 0.5) == 0
    assert modulo(3.5, 1.5) == 0.5
    assert modulo(0, 2) == 0
    assert modulo(20, 0) == 0


def test_divide():
    assert divide(100, 50) == 100 / 50
    assert divide(11, 5) == 11 / 5
    assert divide(27, 20) == 27 / 20
    assert divide(3.5, 0.5) == 7
    assert divide(3.5, 1.5) == 3.5 / 1.5
    assert divide(0, 2) == 0
    assert divide(20, 0) == 0


def test_integer_division():
    assert integer_division(100, 51) == 1
    assert integer_division(11.0, 5.0) == 2
    assert integer_division(27.0, 20.0) == 1
    assert integer_division(3.5, 0.5) == 7
    assert integer_division(3.5, 1.5) == 2
    assert integer_division(0, 2) == 0
    assert integer_division(20, 0) == 0
