import sys

sys.path.append('student')

from concat_string import concat

def test_concat():
    assert concat("Hello", "World") == "Hello, World"
    assert concat("Hi", "there") == "Hi, there"
    assert concat("", "") == ", "
    assert concat("Hello", "") == "Hello, "
    assert concat("", "World") == ", World"
test_concat()
