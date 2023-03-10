## Numerical operations: the return!

### Instructions

Create a file `numerical_operations_the_return.py` containing the following functions:

- `modulo(a, b)`
- `divide(a, b)`
- `integer_division(a, b)`

We assume that `a` and `b` are numbers (`int` or `float`).

> In case of a division by zero or modulo zero your functions should return `0`.

### Usage

Here is a possible `test.py` to test your functions:

```python
import numerical_operations_the_return

print(numerical_operations_the_return.modulo(10, 3))
print(numerical_operations_the_return.divide(10, 3))
print(numerical_operations_the_return.divide(10, 0))
print(numerical_operations_the_return.integer_division(10, 3))
```

```bash
$ python3 test.py
1
3.3333333333333335
0
3
$
```

### Hints

- Some operations will panic in special cases (like division by zero), it is very important to always account for those cases and handle them properly in order to avoid bugs.
- In `Python 2` a division with two integers will return an integer, in `Python 3` it will return a float. We assume you are using `Python 3`, in case you want to force the `Python 3` behavior you can cast one of the operands to float like so: `float(a)`.

### References

- [conditions](https://www.w3schools.com/python/python_conditions.asp)
