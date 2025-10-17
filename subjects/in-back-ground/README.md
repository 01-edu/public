## in-back-ground

### Instructions

Create a script `in-back-ground.sh` which will execute a job in the background that will do the following:

- Run the command `cat` on the file [facts](../../../sh/tests/left/facts) which will read the contents of the file and print it to `stdout`.
- The output of the cat command will be piped to the `grep` command, which will search for the string `"moon"` in the file.
- If the `grep` command succeeds (if it finds the string "moon"):
  - You will print `"The moon fact was found!"` to the `output.txt` file. If it fails the file `output.txt` is not created.
  - You should also print the line which contains the string to the `stdout`.

You must do all these steps running only one job and using the command `nohup`!

Expected output:

```console
$ ls
facts in-back-ground.sh
$ ./in-back-ground.sh
nohup: redirecting stderr to stdout
- Australia is wider than the moon. The moon sits at 3400km in diameter, while Australia's diameter from east to west is almost 4000km.
$ ls
facts in-back-ground.sh output.txt
$ cat output.txt
The moon fact was found!
$
```

```console
$ ./in-back-ground.sh   # If the string isn't found
nohup: redirecting stderr to stdout
$ ls
facts in-back-ground.sh
$ cat output.txt
cat: output.txt: No such file or directory
$
```

> In order to test your solution, you need to create your own `facts` file. This file must not be submitted!

### Hints

The `nohup` command is used to run a command in the background, even if you close the terminal or log out of the system. When a command is run with nohup, it ignores the "SIGHUP" signal, which is sent to processes when the terminal they are running in is closed.

For example, to run the `ls` command in the background with `nohup`, you can use the following command:

```console
$ nohup ls &
```

This will run the "ls" command in the background, and ignore the "SIGHUP" signal. The "&" symbol puts the "ls" command in the background, allowing the terminal to continue accepting input while the "ls" command is running.

In a Unix-like shell, stdin, stdout, and stderr are three standard streams that are used to communicate with a program or process.

- `stdin` (standard input) is a stream of data that a program or process reads from. By default, stdin is the keyboard, but it can be redirected to read from a file or the output of another command.

- `stdout` (standard output) is a stream of data that a program or process writes to. It can be redirected to write to a file or to the input of another command using the `>` operator.

- `stderr` (standard error) is a stream of data that a program or process writes to for error messages and other diagnostic output. It can also be redirected to write to a file or to the input of another command using the `2>` operator.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
