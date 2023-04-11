## Numerical operation

### Instructions

Create a file `numerical_operations.py` containing the following functions:

- `add(a, b)`
- `subtract(a, b)`
- `multiply(a, b)`
- `power(a, b)`
- `square_root(a)`

We assume that `a` and `b` are numbers (`int` or `float`).

### Usage

Here is a possible `test.py` to test your functions:

```python
import numerical_operations

print(numerical_operations.add(2, 2))
print(numerical_operations.subtract(10, 5))
print(numerical_operations.multiply(3, 4))
print(numerical_operations.power(3, 3))
print(numerical_operations.square_root(3))
```

```bash
$ python test.py
4
5
12
27
1.73205080757
$
```

### [Optional] Use a virtual environnement to run python code locally

Virtual environments can help you to run your code locally.

[Learn all you need about virtual environments](https://openclassrooms.com/fr/courses/6951236-mettez-en-place-votre-environnement-python/7013854-decouvrez-les-environnements-virtuels).

Here, we setup a virtual environment with Miniconda.

First, [download and install Miniconda](https://docs.conda.io/en/latest/miniconda.html).

Then, use those commands to create a new environment:

```bash
# create a new virtual environment for python 3.10
$ conda create --name my_env python=3.10

# activate your new environment
$ conda activate my_env

# From now, all python command use this python version in this terminal
$ python --version
Python 3.10.4
```

> We advise you to create one virtual environment per python project. Later, we could also install external packages on our environment.

### Hints

- You could `import math` at the start of your file to use the functions defined in that library (for example `math.sqrt()`).

### References

- [Square root in Python](https://www.geeksforgeeks.org/python-math-function-sqrt/)
- [Operations in Python](https://www.geeksforgeeks.org/python-arithmetic-operators/)
