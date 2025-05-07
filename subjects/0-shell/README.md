## 0-Shell

### Overview

In this project, you'll build **0-shell**, a minimalist Unix-like shell implemented in **Rust**, designed to run core Unix commands using system calls—without relying on external binaries or built-in shells like `bash` or `sh`.

Inspired by tools like [BusyBox](https://en.wikipedia.org/wiki/BusyBox), this project introduces you to key concepts in **Unix system programming**, including **process creation**, **command execution**, and **file system interaction**, all while leveraging Rust's safety and abstraction features to avoid manual memory management.

### Role Play

You are a **system-level developer** assigned to build a lightweight, standalone Unix shell for an embedded Linux environment. Your task is to create a shell that handles basic navigation, file manipulation, and process control—faithfully mimicking essential shell behaviors without relying on existing shell utilities.

### Learning Objectives

- Work with **file and directory operations**
- Manage **user input and output** within a shell loop
- Implement **robust error handling**
- Gain experience in **Unix process and system call APIs**

### Core Requirements

Your minimalist shell must:

- Display a prompt (`$ `) and wait for user input
- Parse and execute user commands
- Return to the prompt only after command execution completes
- Handle `Ctrl+D` (EOF) gracefully to exit the shell

You must implement the following commands **from scratch**, using system-level Rust abstractions:

- `echo`
- `cd`
- `ls` (supporting `-l`, `-a`, `-F`)
- `pwd`
- `cat`
- `cp`
- `rm` (supporting `-r`)
- `mv`
- `mkdir`
- `exit`

Additional constraints:

- Do **not** use any external binaries or system calls that spawn them
- If a command is unrecognized, print:  
  `Command '<name>' not found`

### Constraints

- Only basic command syntax is required  
  (No piping `|`, no redirection `>`, no globbing `*`, etc.)
- Shell behavior should align with Unix conventions
- Code must follow [good coding practices](https://public.01-edu.org/subjects/good-practices/)

### Bonus Features

Implementing any of the following will be considered bonus:

- Handle `Ctrl+C` (SIGINT) without crashing the shell
- Implement the commands exclusively using `low-level system calls` avoiding built-in functions or libraries that abstract file operations.

  - Instead of using functions like the Go `os.Open, os.Remove, and io.Copy`, you would use system calls directly through the `syscall` package using `syscall.Open, syscall.Close, syscall.Read, syscall.Write, syscall.Unlink`.
- Shell usability enhancements:
  - **Auto-completion**
  - **Command history**
  - **Prompt with current directory** (e.g., `~/projects/0-shell $`)
  - **Colorized output** for commands, directories, and errors
  - **Command chaining** with `;`
  - **Pipes** (`|`)
  - **I/O redirection** (`>`, `<`)
- Support for environment variables (e.g. `$HOME`, `$PATH`)
- A custom `help` command documenting built-in functionality

### Example Usage

```shell
student$ ./0-shell
$ cd dev
$ pwd
/dev
$ ls -l
total 0
crw-------  1 root   root     10,    58 Feb  5 09:21 acpi_thermal_rel
crw-r--r--  1 root   root     10,   235 Feb  5 09:21 autofs
drwxr-xr-x  2 root   root           540 Feb  5 09:21 block
...
$ something
Command 'something' not found
$ echo "Hello There"
Hello There
$ exit
student$
```

### Evaluation Criteria

This project will be reviewed and graded based on the following:

- **Functionality**: Commands perform correctly and emulate standard Unix behavior
- **Stability**: Shell handles user errors, invalid input, and edge cases without crashing

### Resources

- [man pages](https://man7.org/linux/man-pages/) (e.g., `man 2 open`, `man 2 execve`)
- [Unix Shell - Wikipedia](https://en.wikipedia.org/wiki/Unix_shell)
- [BusyBox](https://busybox.net/)
- [Rust std::fs and std::process Docs](https://doc.rust-lang.org/std/)
