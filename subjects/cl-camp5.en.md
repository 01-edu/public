## cl-camp5

### Instructions

"keep looking..."

Create a file `lookagain.sh`, which will look for, from the current directory and its sub-folders all the files:

-   all the files ending with `.sh`.

That command will only show the name of the files without the `.sh`.

### Usage

```console
student@ubuntu:~/piscine-go/test$ ./lookagain.sh | cat -e
file1$
file2$
file3$
student@ubuntu:~/piscine-go/test$
```

### Hint

A little `cut`ing might be useful...
