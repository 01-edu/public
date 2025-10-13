## skip_secrets

### Instructions

You need to write a script, `skip_secrets.py`, that will be able to decrypt the text coming from a specific file.

The script will receive a file name as the first argument, check if the file is readable, filter the content by skipping all the lines containing `pineapple` and save the result in a file `out.txt`

If the file passed as argument is not readable or the number of arguments is not the one expected, the script should exit with a status code `1`.

### Usage

Below an example of how the script is supposed to work:

```console
$ cat -e file.txt
A normal pizza $
Another normal pizza$
A pizza with pineapple$
Yet another very normal and delicious pizza$
$ python3 skip_secrets.py
$ echo $?
1
$ python3 skip_secrets.py file.txt
$ cat out.txt
A normal pizza $
Another normal pizza$
A pizza with pineapple$
Yet another very normal and delicious pizza$
$
```

### Hints

- It is possible to read arguments passed to a python script using the `sys` module. Here is an example of script (`argv.py`):

```python
import sys

for argv in sys.argv:
    print(argv)
```

And its output:

```console
$ python3 argv.py 1 2 3 something else
argv.py
1
2
3
something
else
$
```

- It is possible to interrupt the execution of a script and returning a status code different from `0` with the function `exit()`. For example, consider the following `example.py`:

```python
exit(1)
```

This would be the result:

```console
$ python3 example.py
$ echo $?  # check the status code
1
```

### Reference

[Python sys.argv](https://docs.python.org/3/library/sys.html#sys.argv)
