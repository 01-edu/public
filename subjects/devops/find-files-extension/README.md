## find-files-extension

### Instructions

Create a file `find-files-extension.sh`, which will look, from the current directory and its sub-folders for:

Every file with a `.txt` extension and display only the file names.

- You can use this for testing: https://assets.01-edu.org/devops-branch/find-files-extension-example.zip

- What to use : `find`

- The output should be exactly like the example below but with the expected name

```console
$pwd
<..>/find-files-extension-example
$./find-files-extension.sh | cat -e
qwep$
pq1$
zzzz$
ziko$
wei$
ek$
$
```

### Hints

The `find` command is used to search and locate the list of files and directories based on conditions you specify for files that match the arguments:

You can use REGEX in the `find` command.

```console
$ find ~/ -iregex '<REGEX>'
<all files and folders in the user home that respect the specified regex>
$
```

REGEX for txt files: `.*\.\(txt\)`

The `cut` command is a command-line utility that allows you to cut out sections of a specified file or piped data and print the result:

```console
$ cat cut-example.txt
abc-lol
testtest-abcx
qwerty-azerty-x
$ cat cut-example.txt | cut -d "-" -f 1 # cut the lines with "-" as delimiter and print the first part.
abc
testtest
qwerty
$
```

You may need to use pipes `(|)` and the logical operator `&&`.

> You have to use Man or Google to know more about commands flags, in order to solve this exercice!
> Google and Man will be your friends!

### References

[Regex](https://www.computerhope.com/jargon/r/regex.htm)
