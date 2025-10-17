## custom-ls

### Instructions

Create the script `custom-ls.sh` which will create an alias `custom-ls`.

The alias `custom-ls`:

- shows the file details in long list format.
- does not list group information.
- does not ignore entries starting with `.`.
- prints the allocated size of each file, in blocks.
- sorts by file size, largest first.

Expected behavior:

```console
$ custom-ls
error: command not found: custom-ls
$ ./custom-ls.sh
$ custom-ls .
total ...
7784 -rw-r--r--  1 <user>  3983261 Dec 17 22:02 .file1 # just an example
3064 -rw-r--r--  1 <user>  1566444 Dec 17 22:12 file2 # just an example
$
```

### Hints

An alias is a shortcut that references a command. An alias replaces a string that invokes a command in the Linux shell with another user-defined string.

`alias` command instructs the shell to replace one string with another string while executing the commands.

```console
$ alias testcmd="echo 01school"
$ testcmd
01school
$ alias
testcmd='echo 01school'
<...>
$
```

> However, this update alias gets removed after closing the working environment.

To create and add aliases permanently to your bash shell on Linux and Unix-like systems:

1- Edit the `~/.bashrc`:

```console
vi ~/.bashrc
# or #
nano ~/.bashrc
```

2- Append your bash alias, For example append:

```console
alias testcmd="echo 01school"
```

3- Save and close the file.
4- Activate alias

```console
source ~/.bashrc
```

`unalias` unalias removes each alias name from the current shell execution environment.

```console
$ alias
testcmd='echo 01school'
<...>
$ unalias testcmd
$ alias
<...>
$
```

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!

### References

[alias command in linux with examples.](https://www.geeksforgeeks.org/alias-command-in-linux-with-examples/)
[man ls.](https://man7.org/linux/man-pages/man1/ls.1.html)
