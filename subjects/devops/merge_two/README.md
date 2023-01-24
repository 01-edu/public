## Merge two

One very useful data structure in Python are the dictionaries, in this exercise we will start to familiarize and use them.

### Instructions

Create a file `merge_two.py` which will have a function named `merge_two()`. This function will accept two dictionaries and return a third one which will be the merge of the two inputs.

Here is the prototype of the function:

```python
def merge_two(first_dict, second_dict):
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
second = {
    "Louise": 44,
    "Romolo": 30,
    "Lea": 22
}
print(merge_two(first, second))
```

Run your test file with the following command:

```console
$ python3 test.py
{'Bob': 36, 'Louise': 44, 'Lea': 22, 'Romolo': 30}
$
```

### Hints

- If a key is repeated in both dictionaries the value retained will be the one in `second_dict`.

- There is different ways to merge dictionaries, take the time to understand the differences in between those techniques and try more than one technique to better retain it.

### References

- [Merging dictionaries in Python](https://www.geeksforgeeks.org/python-merging-two-dictionaries/)
