## flex_function

### Instructions

Create a file `flex_function.py` that contains a function `create_person` which takes in first name and last name as required positional parameters, age and gender as positional or keyword parameters, size and job as keyword only parameters, with size and job default values: size=1.83, job="taxidermist", and returns a dictionary object like this:

```python
{
    'first_name': 'Kevin',
    'last_name': 'Boulin',
    'age': 34,
    'gender': 'male',
    'size': 1.83,
    'job': 'taxidermist',
}
```

### Usage

Here is an example of how to use the create_person function:

```python
person = create_person("Kevin", "Boulin", 34, "male")
print(person)
```

And its output:

```console
$ python3 ./test.py
{'first_name': 'Kevin', 'last_name': 'Boulin', 'age': 34, 'gender': 'male', 'size': 1.83, 'job': 'taxidermist'}
$
```

### Hints

**Positional parameters** are the parameters that are passed to a function in the same order as they are defined in the function signature. In the create_person function, the first name and last name are considered as positional parameters.

```python
print("Hello World")
```

In this example, "Hello World" is passed as the first and only positional parameter of the function.

**Keyword parameters** are the parameters that are passed to a function by explicitly specifying the parameter name and its value. In the print() function, the end keyword parameter is used to specify the string to be printed at the end of the line.

```python
print("Hello", end=" World")
```

In this example, the value " World" is passed as the value of the end keyword parameter.
