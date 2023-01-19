import sys

sys.path.append('/jail/app/student')

from hello_python import say_hello_python

def test_say_hello_python():
    assert say_hello_python() == "Hello Python!"
