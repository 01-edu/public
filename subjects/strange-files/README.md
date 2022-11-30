## Strange Files

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
User-> code foo\!\\.txt
```

### Instructions

Create a file `"\?$*'First_file'*$?\"` that will contain `Random text inside!` and **nothing else**.

### Usage

```console
User-> ls | cat -e
"\?$*'First_file'*$?\"$
User->
```
