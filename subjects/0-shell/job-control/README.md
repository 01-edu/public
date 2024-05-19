## job control

### Objectives

You must follow the same [principles](https://public.01-edu.org/subjects/0-shell/) as the first subject.

Job control refers to the ability to selectively stop (suspend) the execution of processes and continue (resume) their execution at a later point.

In `job control`, you will have to implement the following [builtins](https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html#Job-Control-Builtins):

- jobs
- bg
- fg
- kill

You must also be able to stop jobs with the `Ctrl + Z`.

### Instructions

- The project has to be written in a compiled language like (C, Rust Go or other), **interpreted languages like (Perl and others) are not allowed**.
- The code must respect the [good practices](https://public.01-edu.org/subjects/good-practices/)

This project will help you learn about:

- Job control
- Process creation and synchronization
- Commands syntax
- Scripting language

### Usage

```
student$ ./0-shell
$ ls -lRr / 2>1 >/dev/null  &
[1] 8287
$ sleep 50 &
[2] 8870
$ jobs
[1]-  Running                 ls -lRr / 2>1 >/dev/null  &
[2]+  Running                 sleep 50 &
$ jobs -l
[1]- 8287 Running                 ls -lRr / 2>1 >/dev/null  &
[2]+ 8870 Running                 sleep 50 &
$ kill 8287
[1]+  Terminated              ls -lRr / 2>1 >/dev/null 
$ jobs
[2]+  Running                 sleep 50 &
$ exit
student$
```
