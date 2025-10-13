## Hello Python

### Welcome to Python (Finally!)

After a week and a half of shell scripting, file permissions, and wrestling with Unix commands, you've earned yourself a change of scenery. Welcome to Python - where syntax errors are friendlier and you don't have to remember whether it's `chmod +x` or `chmod 755` anymore.

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

### A Moment of Relief

Take a breath. No file permissions to debug. No mysterious shell variables to track down. No grep patterns to perfect. Just a simple function that returns a string. This should feel refreshingly straightforward after your recent adventures in the terminal.

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
