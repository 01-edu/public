# Job Control

### Objectives
In this project, you'll extend the `0-shell` project by adding `job control`. Job control refers to the ability to selectively stop (suspend) the execution of processes and continue (resume) their execution at a later point. With job control, your shell will let users run processes both in the foreground and background.

### Instructions
- The project has to be written in a Rust.
- The project must follow the same [principles](https://public.01-edu.org/subjects/0-shell/) as the first subject.
- The code must respect the [good practices](https://public.01-edu.org/subjects/good-practices/).
- You must implement the following commands:
	- The `&` operator to run processes in the background.
	- `jobs` (supporting `-r`, `-l`, `-p`, `-s`)
	- `bg`
	- `fg`
	- `kill` (including handling for job specifiers like `%1`)

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

### Learning objectives
This project will help you learn about:
- Job control
- Process creation and synchronization
- Commands syntax    
- Scripting language
