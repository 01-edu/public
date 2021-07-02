## cl-camp5

### Instructions

"keep looking..."

Create a file `lookagain.sh`, which will look, from the current directory and its sub-folders for:

- all the files ending with `.sh`.

That command will only show the name of the files without the `.sh`.

### Usage

```console
$ ./lookagain.sh | cat -e
file1$
file2$
file3$
$
```

### Hint

A little `cut`ing might be useful...
