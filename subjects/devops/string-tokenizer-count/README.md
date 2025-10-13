## string_tokenizer_count

### Instructions

Create a file `string_tokenizer_count.py` that contains a function `tokenizer_counter` which takes in a string as a parameter and returns a dictionary of words and their count in the string.

- The function should remove any punctuation from the string and convert it to lowercase before counting the words.

- The function should return a dictionary of words and their count, sorted alphabetically by word.

### Usage

Here is an example of how to use the function in a `test.py` script:

```python
string = "This is a test sentence, with various words and 123 numbers!"
result = tokenizer_counter(string)
print(result)
```

And its output:

```console
$ python3 test.py
{'123': 1, 'a': 1, 'and': 1, 'is': 1, 'numbers': 1, 'sentence': 1, 'test': 1, 'this': 1, 'various': 1, 'with': 1, 'words': 1}
$
```

### Hints

- The `re` module can be used to remove non-alphanumeric characters.

- The `Counter` class of the `collections` module can be used to count the words.

- The `operator` module can be used to sort the dictionary alphabetically by word.

Here is an example of how to sort a dictionary in python, using a `test.py` script:

```python
dictionary = {'a': 5, 'c': 1, 'b': 3}
sorted_dict = dict(sorted(d.items(), key=lambda item: item[1]))
print(sorted_dict)
```

And its output:

```console
$ python3 ./test.py
{'c': 1, 'b': 3, 'a': 5}
$
```

### References

- [`re` module](https://docs.python.org/3/library/re.html)

- [`collections` module](https://docs.python.org/3/library/collections.html)

- [`operator` module](https://docs.python.org/3/library/operator.html)
