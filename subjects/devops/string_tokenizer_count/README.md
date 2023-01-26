## string_tokenizer_count

### Instructions

Create a file string_tokenizer_count.py that contains a function tokenizer_counter which takes in a string as a parameter and returns a dictionary of words and their count in the string.

- The function should remove any punctuation from the string and convert it to lowercase before counting the words.

- The function should return a dictionary of words and their count, sorted alphabetically by word.

### Usage

Here is an example of how to use the function:

```python
string = "This is a test sentence, with various words and 123 numbers!"
result = tokenizer_counter(string)
print(string)
```

And its output:

```console
string = "This is a test sentence, with various words and 123 numbers!"
result = tokenizer_counter(string)
```

### Hints

- The `re` module can be used to remove non-alphanumeric characters.

- The `collections` module can be used to count the words.

- The `operator` module can be used to sort the dictionary alphabetically by word.

### References

- [`re` module](https://docs.python.org/3/library/re.html)

- [`collections` module](https://docs.python.org/3/library/collections.html)

- [`operator` module](https://docs.python.org/3/library/operator.html)