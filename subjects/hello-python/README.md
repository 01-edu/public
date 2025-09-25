## Hello Python

Here starts your journey in the marvelous world of Python, a very versatile programming language, popular for being accessible for beginner programmers and yet very powerful.

### Setup

First, you'll need to have Python installed on your machine. In those exercises we use Python 3.10, we then recommend you to install Python 3.10 or higher.

- [Download and install Python](https://www.python.org/downloads/)

To write your code, you may use your favorite IDE. If you don't have one, you can go for Visual Studio Code with Python extension pack installed:

- [Get VS Code](https://code.visualstudio.com/)
- (Optional) [The Python extension pack](https://marketplace.visualstudio.com/items?itemName=donjayamanne.python-extension-pack).

> Setup may vary in base of your Operating System or machine specificities, if you feel stuck check on internet for the specific steps that suits your configuration.

### Instructions

Create a file `hello_python.py` and write inside a function `say_hello_python` that returns the string `"Hello Python!"`.

```python
def say_hello_python():
    # this is a function,
    # write your code here
```

### Usage

Here is a possible code to test your function. Put it in another file (ex: `test.py`):

```python
from hello_python import say_hello_python

print(say_hello_python())
```

Run your test file with the following command:

```console
$ python3 test.py
Hello Python!
$
```

> `test.py` file should be in the same directory of `hello_python.py` in order to work.

### Hints

- Python uses indentation to indicate in which block your code will run (many other languages uses parenthesis instead). It is then very important to indent your code properly.

- `return` is a special world used to say which value a function should return (a function could also not return anything).
  For example `return 10` will return the number ten.

- A string is a set of characters wrapped by `"`.
  For example `"Ciao bella"` is a string containing the worlds `Ciao Bella`.

- In `test.py` we are calling another function named `print`. This function is already present in Python standard library so we don't have to create it. This function at its core is made to write strings in the standard output.

### References

- [Indentation in Python](https://www.w3schools.com/python/gloss_python_indentation.asp)
- [Python print function](https://www.w3schools.com/python/ref_func_print.asp)
