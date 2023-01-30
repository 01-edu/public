## Clean the list

### Instructions

Create a file `clean_the_list.py` that contains a function `clean_list` which takes and returns a list of strings and performs the following operations on each list item:

- Removes all spaces before and after (but not between words).

- Capitalizes the first letter (first letter only, other ones should be to lowercase).

- Adds its index number before a separator `x/ `.

- An empty list as input should return an empty list as output.

- And don't forget the milk! (add it at the end of the list if missing).

### Usage

Here is an example of how to use the `clean_list` function:

```python
import clean_the_list
shopping_list = ['tomatoes', 'pastas', 'milk', 'salt']
print(clean_the_list.clean_list(shopping_list))
```

This will output:

```console
$ python test.py
['1/ Tomatoes', '2/ Pastas', '3/ Milk', '4/ Salt']
$
```

### References

[String strip method](https://www.w3schools.com/python/ref_string_strip.asp)
[String capitalize method](https://www.w3schools.com/python/ref_string_capitalize.asp)
