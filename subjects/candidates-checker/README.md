## Candidates checker

### Instructions

Create a file `candidates_checker.py` which will receive the number of candidates as the only argument.
If a wrong number of arguments is passed the script will print `Error: wrong number of arguments` and exit with `1`.

When provided, the argument will always be convertible to `int`, the script will then ask for each candidate the name as string and the age as number.

Once the information for each candidate is retrieved for each of them the script will check the age and print one of the following results:

- `"[name] is not eligible (underaged)"` when the age is less than 18.
- `"[name] is not eligible (over the legal age)"` when the age is more than 60.
- `"[name] is eligible"` when the age is between 18 and 60 (included).

> You must use dictionaries to save the data about the candidates.

### Usage

Here is an example of your script running:

```console
$ python3 candidates_checker.py 3
Add a new candidate:
name: Edoardo
age: 17
Add a new candidate:
name: Michele
age: 60
Add a new candidate:
name: Lea
age: 61
Edoardo is not eligible (underaged)
Michele is eligible
Lea is not eligible (over the legal age)
$
```

Here is an example of bad usage:

```console
$ python3 candidates_checker.py
Error: wrong number of arguments
$
```

### Hints

- In order to succeed your script should print **exactly** the same output as the one in the usage section. So `Add a new candidate:`, `name: ` and `age: ` should be written in the exact same way and order.

- Tough it is not mandatory you could use `if __name__ == '__main__':` to specify the entrypoint of your script.

### References

- [Function strip()](https://docs.python.org/3.11/library/stdtypes.html?highlight=strip#str.strip)
- [String multiplication](https://www.geeksforgeeks.org/create-multiple-copies-of-a-string-in-python-by-using-multiplication-operator/)
- [Entrypoint for a Python script](https://realpython.com/if-name-main-python/)
