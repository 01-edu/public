## Merge two

One very useful data structure in Python are the dictionaries, in this exercise we will start to familiarize and use them.

### Instructions

Create a file `merge_two.py` which will have a function named `merge_two()`. This function will accept one dictionary.
It will prompt the user to create a new dictionary asking for keys and values.
The function will always convert the `values` into integers.

As a return it will create a third one which will be the merge of the two dictionaries and return it as a serialized JSON string.

> If the key entered by the user is `exit` the function will stop asking for new key/values pairs and proceed to generate the desired output.

Here is the prototype of the function:

```python
def merge_two(first_dict):
    # this is a function,
    # write your code here
```

### Usage

Here is a possible code to test your function. Put it in another file (ex: `test.py`):

```python
from merge_two import merge_two

first = {
    "Bob": 36,
    "Louise": 23,
    "Lea": 34
}

print(merge_two(first))
```

Run your test file with the following command:

```console
$ python3 test.py
Add a new entry:
key: Louise
value: 44
Add a new entry:
key: Romolo
value: 30
Add a new entry:
key: Lea
value: 22
Add a new entry:
key: exit
{"Bob": 36, "Louise": 44, "Lea": 22, "Romolo": 30}
$
```

### Hints

- If a key is repeated in both dictionaries the value retained will be the last one.

- There are different ways to merge dictionaries, take the time to understand the differences in between those techniques and try more than one technique to better retain it.

- Add `import json` to use the standard functions for JSON manipulation.

- Use the function `input()` to read from `stdin` and `int()` to convert the value to a number.

> Your solution will be tested only for valid inputs (all the values will be convertible to `int`).

### References

- [Merging dictionaries in Python](https://www.geeksforgeeks.org/python-merging-two-dictionaries/)
- [JSON library in Python](https://docs.python.org/3/library/json.html)
- [Function input() in Python](https://www.w3schools.com/python/ref_func_input.asp)
