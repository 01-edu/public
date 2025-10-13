## in-back-ground

### Instructions

Create a script named `in-back-ground.sh` that executes a job in the background with the following behavior:

1. Run the `cat` command on the file [facts](./facts).
2. Pipe the output of `cat` into the `grep` command to search for the string `"moon"`.
3. If the `grep` command succeeds (i.e., it finds a line containing `"moon"`):
   - Append the message `"The moon fact was found!"` to the file `output.txt`.
   - Also, print the matching line to **stdout**.
4. If the `grep` command does not find the string, the file `output.txt` should not be created.

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
