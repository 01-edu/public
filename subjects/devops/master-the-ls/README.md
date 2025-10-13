## master-the-ls

### Instructions

Create a script `master-the-ls.sh`, that will do the following:

- list the files and directories of the current directory.
- Ignore the hidden files, the current directory `.` and the parent directory `..`.
- Separates the results with commas.
- Order them by ascending order of access time (the newest first).
- Have the directories ends with a `/`.

### Hints

Here are some Commands that can help you:

- `tr`. Translate characters: run replacements based on single characters and character sets.
  - Replace all occurrences of a character in a file, and print the result:
    `tr {{find_character}} {{replace_character}} < {{filename}}`

```console
$ cat uid.txt
My
UID
is
1000
$ tr "\n" " " < uid.txt
My UID is 1000
$
```

- `ls`. List directory contents.
- `sed`. Edit text in a scriptable manner.

```console
$ cat text.txt
unix is a great os. unix is opensource. unix is a free os.
$ sed 's/unix/linux/' text.txt
linux is a great os. unix is opensource. unix is a free os.
$
```

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!

### References

- [tr](https://www.gnu.org/software/coreutils/tr).
- [ls](https://www.gnu.org/software/coreutils/ls).
- [sed](https://www.gnu.org/software/sed/manual/sed.html).
