## credentials_searches

### Instructions

Create a file `credentials_searches.py` that contains a function `credentials_search` which takes no parameters and searches for the keys `password` and `secret` in a file called `logs.json`.

- If both keys are found, the function should create a new json file named `credentials.json` and save the values of those keys in it.

- If only one of the keys is found, the function should create a new json file named `credentials.json` and save the value of that key in it.

- If neither key is found, the function should not create the `credentials.json` file.

- If the `logs.json` file is empty or is not a json file, the function should not create the `credentials.json` file.

- If the `logs.json` file does not exist, the function should not create the `credentials.json` file.

- If the keys are found in a nested object, the function should create the `credentials.json` file and save the values of the keys in it.

### Usage

Here is an example of how to use the credentials_search function:

With this file `logs.json`:

```json
{
  "password": "test_password",
  "other": "value",
  "secret": "test_secret"
}
```

If you run your function:

```python
credentials_search()
```

Your function should generate this file `credentials.json`:

```json
{
  "password": "test_password",
  "secret": "test_secret"
}
```

### Hints

- The `json` module can be used to read and write JSON files.

- The `os` module can be used to check if a file exists and remove a file.

- Remember that the function should be able to search for the keys `password` and `secret` in nested json objects.

- Using the `isinstance` function could be useful to check if a value is for example a dict.

### References

- [json module](https://www.w3schools.com/python/python_json.asp)

- [os module](https://docs.python.org/3/library/os.html)

- [isinstance function](https://www.w3schools.com/python/ref_func_isinstance.asp)
