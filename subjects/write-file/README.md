## Write file

### Instructions

Create a file `write_file.py` which will have a function `to_do(input)` with one input as an argument `list[tuple[dt.date, str]]` which will be a to-do list like the following example:

```python
[
    (dt.date(2022, 6, 1), "fix the seat"),
    (dt.date(2022, 6, 2), "take the trash"),
]
```

Your function needs to format and write the input in a text file `output.txt` as follows:

```bash
$ cat output.txt
Wednesday 01 June 2022: fix the seat
Thursday 02 June 2022: take the trash
$
```

Make sure you follow the correct format for the output. The given lists will always have the correct format and will never be empty.

### Usage

Here is a possible `test.py` to test your function:

```python
import datetime as dt
from write_file import to_do

to_do_list = [
    (dt.date(2022, 6, 1), "fix the seat"),
    (dt.date(2022, 6, 2), "take the trash"),
]
to_do(to_do_list)
```

```bash
$ python test.py
$ cat output.txt
Wednesday 01 June 2022: fix the seat
Thursday 02 June 2022: take the trash
$
```

### Hints

- The `strftime() `method is used to format the date object as a string. You can use this method to format the date in the desired format (e.g. %A %d %B %Y).

```python
import datetime as dt

date = dt.date(2022, 6, 1)
formatted_date = date.strftime("%A %d %B %Y")
print(formatted_date)
```

The output:

```bash
Wednesday 01 June 2022
```

- The `%A` and `%B` format codes are used to represent the full weekday and month name respectively. The `%d` format code is used to represent the day of the month. The `%Y` format code is used to represent the year.

- The `write()` method is used to write a string to a file. You can use this method to write the formatted date and task to the output file.

```py
with open("output.txt", "w") as file:
    file.write("Hello World!")
```

In this example, the `open()` function is used to open the file named `output.txt` with the write mode `w`. The `with` statement is used to open the file, and automatically close it after the indented block of code is executed.

The `write()` method is used to write the string "Hello World!" to the file.

- Make sure to include a newline character `\n` at the end of each line of output, to separate the tasks in the output file.

- Test your function with different input formats, different date and task length, different date range and make sure that it works as expected.

### References

- [open a file](https://docs.python.org/3/tutorial/inputoutput.html#tut-files)
- [write in a file](https://www.pythontutorial.net/python-basics/python-write-text-file/)
