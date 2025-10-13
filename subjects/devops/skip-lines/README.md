## Skip lines

### Instructions

Write a command line in a `skip-lines.sh` file that prints the result of a `ls -l` skipping every even line, starting with the **first** one.

Example:

```console
$ ls -l
drwxr-xr-x 16 User User      4096 nov 11 10:55 Desktop
drwxr-xr-x 22 User User      4096 nov  4 10:02 Documents
drwxr-xr-x  6 User User      4096 nov 11 10:40 Downloads
drwxr-xr-x  2 User User      4096 mar 31  2022 Music
drwxr-xr-x  2 User User      4096 set 29 10:34 Pictures
drwxr-xr-x  2 User User      4096 nov 23  2021 Public
```

What we want your script to do is:

```console
drwxr-xr-x 16 User User      4096 nov 11 10:55 Desktop
drwxr-xr-x  6 User User      4096 nov 11 10:40 Downloads
drwxr-xr-x  2 User User      4096 set 29 10:34 Pictures
```

### Hints

Here are some Commands that can help you:

- `sed`. Edit text in a scriptable manner.
  - Print just a first line to stdout:
    `{{command}} | sed -n '1p'`

```console
$ cat file
AIX
Solaris
Unix
Linux
HPUX
$ sed -n '1p' file
AIX
```

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!

### References

- [Sed](https://www.gnu.org/software/sed/manual/sed.html).
