## Punishment

### Instructions

Often in schools we are asked to copy hundreds of sentences in order to better remember not to do something, this punishment is very old and quite boring!

Thankfully we now have Python that can do the job for us.
In order to do so create a file `punishment.py` which will contain a function `do_punishment` having 3 arguments:

- `first_part`: which will be a string.
- `second_part`: which will be also a string.
- `nb_lines`: which will be a number.

Here is the prototype of the function:

```python
def do_punishment(first_part, second_part, nb_lines):
```

The function will concatenate `first_part` and `second_part`, adding a space in between them and a `.` at the end of `second_part`. It will repeat this process for `nb_lines` times.

The function will return a single string containing all the repeated sentences.

> In case `first_part` and `second_part` have empty spaces at the start or at the end those spaces should be trimmed (removed from the strings).

### Usage

Here is a possible `test.py` to test your functions:

```python
import punishment

print(punishment.do_punishment('   The first half   ', '   and the second  ', 4), end='')
print(punishment.do_punishment('Will not', 'show', 0), end='')
print(punishment.do_punishment('', '', 3), end='')
```

```bash
$ python test.py
The first half and the second.
The first half and the second.
The first half and the second.
The first half and the second.
 .
 .
 .
```

### Hints

- Removing spaces at the start and end of a string is so common that almost all languages implement this feature. Here you can check for the `strip()` method.
- Instead of using loops you can try the `string multiplication` operator, which is a very nice feature of Python and will make your code more readable.

### References

- [Function strip()](https://docs.python.org/3.11/library/stdtypes.html?highlight=strip#str.strip)
- [String multiplication](https://www.geeksforgeeks.org/create-multiple-copies-of-a-string-in-python-by-using-multiplication-operator/)
