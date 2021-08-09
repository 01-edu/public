## cl-camp5

### Instructions

"keep looking..."

Create a file `lookagain.sh`, which will look, from the current directory and its sub-folders for:

- all the files ending with `.sh`.

That command will only show the name of the files without the `.sh`.
The files will be in descending order (as shown in the below example).

### Usage

```console
$ ./lookagain.sh | cat -e
file3$
file2$
file1$
$
```

### Hint

A little `cut`ing might be useful...
