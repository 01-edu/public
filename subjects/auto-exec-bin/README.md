## auto-exec-bin

### Instructions

Create a file `auto-exec-bin.sh`, which will make a binary with the name `01exec` in `~/myBins` executable from any working directory.

You can use any binary from your choice to test your script.

> You can pick any binary for the tests!

Expected Output:

```console
$ ls -l ~/myBins # the binary
01exec
$ 01exec
error: command not found: 01exec
$ ./auto-exec-bin.sh
$ 01exec
Hello 01 Scripting Pool
$ cd /{random-path} && 01exec
Hello 01 Scripting Pool
$
```

### Hints

`PATH` environment variable is a variable where the shell search for the binaries for the execution.

When you put a command the shell will search for binary in `PATH` folders

![auto binary exec](resources/auto-exec-diagram.png)

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
