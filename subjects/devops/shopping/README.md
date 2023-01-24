# shopping

## Instructions

Create a file `shopping.py` with a function `remember_the_apple(shopping_list)`.

This function input is a list of string, like this:

```python
['tomatoes', 'pastas', 'apple', 'salt']
```

If the string `'apple'` is not in the list, it should be added. The function return the updated list. If the apple is in the list, simply return it.

If the input list is empty, the returned list should be also empty (don't add the apple in that case).

## Usage

Here is a possible `test.py` to test your functions:

```python
import shopping

if __name__ == '__main__':
    list_with_apple = ['tomatoes', 'pastas', 'apple', 'salt']
    list_without_apple = ['tomatoes', 'pastas', 'salt']

    print(shopping.remember_the_apple(list_with_apple))
    print(shopping.remember_the_apple(list_without_apple))
```

```bash
$ python test.py
['tomatoes', 'pastas', 'apple', 'salt']
['tomatoes', 'pastas', 'salt', 'apple']
```

## Notions

- [lists](https://docs.python.org/3/tutorial/datastructures.html#)
- [Conditions](https://docs.python.org/3/tutorial/datastructures.html#more-on-conditions)
