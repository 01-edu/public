## my-ls-1

### Objectives

- my-ls-1 consists on creating your own `ls` command.

- The `ls` command shows you the files and folders of the directory specified after the command. By exclusion of this directory, it shows the files and folders of the present directory.

- The behavior of your `ls` must be identical to the original `ls` command with the following variations :
  - You must incorporate in your project at least the following flags of the `ls` command :
    - `-l`
    - `-R`
    - `-a`
    - `-r`
    - `-t`
  - Note that you can use various flags, in various ways, just like in `ls`.
  - When it comes to the `ls -l` display, it must be identical to the system command.
  - Other flags displays are up to you.

This project will help you learn about :

- Unix system
- Ways to receive data.
- Ways to output data.
- Manipulation of strings.
- Manipulation of structures.

### Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices.en).
- It is recommended that the code should present a **test file**.

### Allowed packages

- fmt
- os
- os/user
- strconv
- strings
- syscall
- time
- math/rand
- errors

### Hint

- We strongly recommend that you account for the implications of the option `-R` from the very beginning of your code.

- The order that files and folders appear must be taken in consideration.

- We suggest that you consult the `ls` command manual.

### Usage

You can see how the `ls` command works, by using it on your terminal.
