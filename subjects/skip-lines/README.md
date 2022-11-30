## Skip lines

### Instructions

Write a command line in a `skip-lines.sh` file that prints the result of a `ls -l` skipping 1 line out of 2, starting with the **first** one.

Example:

```console
User-> ls -l
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

**Tips:**

Here are some Commands that can help you:

- `sed`. Edit text in a scriptable manner. You can see also: awk. [For more information](https://www.gnu.org/software/sed/manual/sed.html).

  - Print just a first line to stdout:
    `{{command}} | sed -n '1p'`

```console
User-> cat file
AIX
Solaris
Unix
Linux
HPUX
User-> sed -n '1p' file
AIX
```

Use man to see how `awk` or `sed` works, they can do the job.
