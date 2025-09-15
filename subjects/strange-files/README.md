## Strange Files

### Instructions

Create the following files:

- `firstFile` which contains `Random text inside!` and **nothing else**
- `"medium_File!"` which contains `Random text inside!` and **nothing else**
- `"\?$*'Hard_file'*$?\"` which contains `Random text inside!` and **nothing else**.

> Watch the videos:
>
> - [Escape Characters](https://www.youtube.com/watch?v=ZenUrQeHl7c)

### Usage

```console
$ ls | cat -e
"\?$*'Hard_file'*$?\"$
firstFile$
"medium_File!"$
$
```

### Hints

We come across files and folders name very regularly. In most of the cases file/folder names are related to the content of the file/folder and starts with numbers or letters. Alpha-Numeric file name are pretty common and very widely used, but this is not the case when we have to deal with file/folder name that has special characters in them.

Example of most common file names are:

```console
something.txt
alphanum21.txt
674659.txt
```

Examples of file names that have special characters:

```console
#121day.txt
some/file#.txt
*File$*.txt
```

When You create a file with some special character, you will need to `escape` those special characters in order to create that file.

A non-quoted backslash `\` is the escape character. The backslash `\` preceding a character tells the shell to interpret that special character literally.

Example:

If you want to create a file named `foo!\.txt` You have to escape the characters like so:

```console
$ touch foo\!\\.txt
```

- `ls`. List directory contents.
- `touch` used to create, change and modify timestamps of a file.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!

### References

- [ls](https://www.gnu.org/software/coreutils/ls).
- [touch](https://www.gnu.org/software/coreutils/touch).
