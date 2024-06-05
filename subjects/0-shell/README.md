## 0-shell

### Objective

The objective of this project is for you to create a simple [shell](https://en.wikipedia.org/wiki/Unix_shell).

Through the `0-shell` you will get to the core of the `Unix` system and explore an important part of this system’s API which is the process creation and synchronization.
Executing a command inside a shell implies creating a new process, which execution and final state will be monitored by its parents processes. This set of functions will be the key to success for your project.

For this project you will only have to create a simple `Unix shell` where you can run some of the most known commands. For this part of the project, no advanced functions, pipes or redirection will be asked, but you can add them if you like.

### Instructions

- You must program a mini `Unix shell`, try to focus on something simple like [BusyBox](https://en.wikipedia.org/wiki/BusyBox).
- This interpreter must display at least a simple `$` and wait until you type a command line which will be validated by pressing enter.
- The `$` will be shown again only once the command has been completely executed.
- The command lines are simple, you will not have pipes, redirection or any other advanced functions.
- You must manage the errors, by displaying a message adapted to the error output.
- You must implement the following commands:
  - `echo`
  - `cd`
  - `ls`, including the flags `-l`, `-a` and `-F`
  - `pwd`
  - `cat`
  - `cp`
  - `rm`, including the flag `-r`
  - `mv`
  - `mkdir`
  - `exit`

> The commands need to be implemented from scratch, without calling any external binaries.

- You must manage the program interruption `Ctrl + D`.
- The project has to be written in a compiled language (like C, Rust, Go or other), **interpreted languages (like Perl and others) are not allowed**.
- The code must respect the [good practices](https://public.01-edu.org/subjects/good-practices/)

This project will help you learn about:

- Shell
- Operating systems services
- Command-line interfaces
- Unix system
- Process creation and synchronization
- Commands syntax
- Scripting language

### Bonus

You can also do more bonus features like:

- Implement the commands exclusively using `low-level system calls` avoiding built-in functions or libraries that abstract file operations.

  - Avoid High-Level Abstractions: Instead of using functions like the Go `os.Open, os.Remove, and io.Copy`, you would use system calls directly through the `syscall` package using `syscall.Open, syscall.Close, syscall.Read, syscall.Write, syscall.Unlink`.

- Manage the interruption `Ctrl + C`
- Auto complete when you are writing
- Add piping
- Add redirection
- Have your path behind the `$` like (~/Desktop/0-shell $)
- Add colors for the directories or errors
- Other advanced commands you may like.

### Usage

```
student$ ./0-shell
$ cd dev
$ pwd
dev
$ ls -l
total 0
crw-------  1 root   root     10,    58 fev  5 09:21 acpi_thermal_rel
crw-r--r--  1 root   root     10,   235 fev  5 09:21 autofs
drwxr-xr-x  2 root   root           540 fev  5 09:21 block
crw-------  1 root   root     10,   234 fev  5 09:21 btrfs-control
drwxr-xr-x  3 root   root            60 fev  5 09:20 bus
drwxr-xr-x  2 root   root          4400 fev  5 09:21 char
crw-------  1 root   root      5,     1 fev  5 09:21 console
lrwxrwxrwx  1 root   root            11 fev  5 09:20 core -> /proc/kcore
drwxr-xr-x  2 root   root            60 fev  5 09:20 cpu
crw-------  1 root   root     10,    59 fev  5 09:21 cpu_dma_latency
$ something
Command 'something' not found
$ echo "Hello There"
Hello There
$ exit
student$
```
