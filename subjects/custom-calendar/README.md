## Custom calendar

### Instructions

Create a file `custom_calendar.py` which will have 2 functions:

- `day_from_number(day_number)`
- `day_to_number(day)`

Those functions perform conversion between day `index` and day `word`:

- 1 = Monday
- 2 = Tuesday
- 3 = Wednesday
- 4 = Thursday
- 5 = Friday
- 6 = Saturday
- 7 = Sunday

You should return `None` if the input is invalid (invalid number or day string).

### Usage

Here is a possible `test.py` to test your functions:

```python
import custom_calendar

print(custom_calendar.day_from_number(2))
print(custom_calendar.day_from_number(1))
print(custom_calendar.day_from_number(1000))
print(custom_calendar.day_to_number('Sunday'))
print(custom_calendar.day_to_number('invalid day'))
```

```console
$ python test.py
Tuesday
Monday
None
7
None
$
```

### Hints

Dictionaries:

A dictionary is a data type similar to arrays, but works with keys and values instead of indexes. Each value stored in a dictionary can be accessed using a key, which is any type of object (a string, a number, a list, etc.) instead of using its index to address it.

For example, a database of phone numbers could be stored using a dictionary like this:

```python
phonebook = {}
phonebook["John"] = 938477566
phonebook["Jack"] = 938377264
phonebook["Jill"] = 947662781
print(phonebook)
```

output:

```console
{'John': 938477566, 'Jack': 938377264, 'Jill': 947662781}
```

Alternatively, a dictionary can be initialized with the same values in the following notation:

```python
phonebook = {
    "John" : 938477566,
    "Jack" : 938377264,
    "Jill" : 947662781
}
print(phonebook)
```

Dictionaries can be iterated over, just like a list. To iterate over key value pairs, use the following syntax:

```python
phonebook = {"John": 938477566, "Jack": 938377264, "Jill": 947662781}
for name, number in phonebook.items():
    print("Phone number of %s is %d" % (name, number))
```

output:

```console
Phone number of John is 938477566
Phone number of Jack is 938377264
Phone number of Jill is 947662781
```

### References

- [None type](https://www.w3schools.com/python/ref_keyword_none.asp)
- [Dictionaries](https://docs.python.org/3/tutorial/datastructures.html#dictionaries)
- [Access an item in dictionary](https://www.w3schools.com/python/python_dictionaries_access.asp)
