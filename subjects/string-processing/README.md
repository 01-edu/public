## String processing

### Instructions

Tokenization is the process of breaking down a string into smaller pieces, called tokens. In natural language processing, tokenization typically refers to the process of breaking down a sentence into words or breaking down a paragraph into sentences.

Create a file `string_processing.py` which will have a function `tokenize(sentence)` that given a sentence will do the following:

- remove all punctuation marks and special characters
- separate all words like so: `"it's not 3" => ['it', 's', 'not', '3']`
- put all the words in lowercase
- return a list of all the words.

### Usage

Here is a possible `test.py` to test your functions:

```python
import string_processing

if __name__ == '__main__':
    my_sentence = "It's not possible, you can't ask for a raise"
    print(string_processing.tokenize(my_sentence))
```

```bash
$ python test.py
['it', 's', 'not', 'possible', 'you', 'can', 't', 'ask', 'for', 'a', 'raise']
```

### Hints

The `re` library is a module for working with regular expressions. It provides a set of functions for working with regular expressions, including:

- `re.sub()` : Replaces all occurrences of a regular expression pattern in a string with a replacement string.

```python
text = "This is a test sentence. It has multiple punctuation marks!"

# Replace all exclamation marks with question marks
new_text = re.sub("!", "?", text)

print(new_text)
```

and the output:

```console
This is a test sentence. It has multiple punctuation marks?
```

The `.lower()` method is used to convert the sentence to lowercase before tokenizing it.

```python
text = "This Is A TeST Sentence."

lower_text = text.lower()

print(lower_text)
```

and the output:

```console
this is a test sentence.
```

The `.split()` method is used to split the sentence into a list of words.

```python
text = "This is a test sentence."

words = text.split()

print(words)
```

and the output:

```console
['This', 'is', 'a', 'test', 'sentence.']
```

### References

- [string methods](https://www.w3schools.com/python/python_ref_string.asp)
- [replace](https://www.w3schools.com/python/ref_string_replace.asp)
- [split](https://www.w3schools.com/python/ref_string_split.asp)
- [String punctuations](https://docs.python.org/3/library/string.html#string.punctuation)
- [Tokenization in text analysis](https://en.wikipedia.org/wiki/Lexical_analysis#Tokenization)
- [Word segmentation](https://en.wikipedia.org/wiki/Text_segmentation#Word_segmentation)
