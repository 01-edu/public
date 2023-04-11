## head-and-tail

### Instructions

Create the script `head-and-tail.sh` which will show the first and last lines of the file `HeadTail.txt`

- Where to look : https://assets.01-edu.org/devops-branch/HeadTail.txt

- What to use : `curl` and `head` and `tail`

- The output should be exactly like the example below but with the expected name

```console
$ ./head-and-tail.sh | cat -e
Hello My username is: <...>,$
so my passwd is: <...>$
$
```

### Hints

With `curl` you can get the content of the file from the url:

```console
$ curl https://assets.01-edu.org/devops-branch/HeadTail.txt
<...>
$
```

`head` print the top N number of data of the given input. By default, it prints the first 10 lines of the specified files:

```console
$head file
<first 10 lines>
$head -1 file
<first 1 line>
$
```

`tail` print the last N number of data of the given input. By default, it prints the last 10 lines of the specified files:

```console
$ tail file
<last 10 lines>
$ tail -1 file
<last 1 line>
$
```

You may need to use pipes `(|)` and the logical operator `&&`.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
