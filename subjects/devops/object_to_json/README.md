## object_to_json

A lot of information on the web are shared in the JSON format. This exercise will be about transforming object to JSON and vice-versa.

### Instructions

You just landed a new job, congrats! Your new task is to build two functions to allow new users to register to your new shiny website.

The new registration information comes as string formatted as JSON. You need to create a file `registration.py` that will have the following functions inside:

1. `create_new_user` that will receive a `dict` and will return a new object of the class `User` provided below. To be valid, the input `dict` must have a `username` key and a `email` key to be valid. The new `User` will have the same `username` and `email` of the input `dict`. If the input `dict` is invalid, the default user will be returned.

```python
class User:
    username = 'user'
    email = 'something@email.com'
```

2. `user_to_json` that will receive a `User` and will return the object as a string in JSON format. Be aware of the Python types that can be converted to JSON!

### Usage

Here is a possible `test.py` to test your functions:

```python
import create_new_user
import user_to_json

registration_0 = '{"username": "mario", "email": "mario@me.it"}'
registration_1 = '{"city": "Rome", "country": "Italy"}'

user_0 = create_new_user(registration_0)
user_1 = create_new_user(registration_1)

print(user_to_json(user_0))
print(user_to_json(user_1))
```

```console
$ python3 test.py
{"username": "mario", "email": "mario@me.it"}
{}
$
```

### Hints

- It is possible to convert an object to `dict` using the `__dict__` casting.

```python
class C:
    a = 1
    b = 2

my_obj = C()
my_obj_dict = my_obj.__dict__ # my_obj_dict will be a dictionary with the object my_obj values
```

### Notions

- [Class and Object](https://docs.python.org/3/tutorial/classes.html#a-first-look-at-classes)
- [json.loads()](https://www.geeksforgeeks.org/json-loads-in-python/)
- [json.dumps()](https://www.geeksforgeeks.org/json-dumps-in-python/)
