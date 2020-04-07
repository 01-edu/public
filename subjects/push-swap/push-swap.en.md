## push-swap

### Objectives

Push_swap is a very simple and highly effective algorithm.You have at your disposal a list of `int` values, two stacks (`a` and `b`) and a set of instructions.

You will have to write 2 programs:

- **push_swap**, which calculates and displays on the standard output the smallest program using push_swap instruction language that sorts integer arguments received.
- **checker**, which takes integer arguments and reads instructions on the standard output. Once read, checker executes them and displays `OK` if integers are sorted. Otherwise, it will display `KO`.

As said before, you will have two stacks at your disposal. Your goal is to sort stack `a`, that will contain the `int` values received, in ascending order, using both stacks and a set of instructions.

These are the instructions that you can use to sort the stack :

- `pa` push the top first element of stack `b` to stack `a`
- `pb` push the top first element of stack `a` to stack `b`
- `sa` swap first 2 elements of stack `a`
- `sb` swap first 2 elements of stack `b`
- `ss` execute `sa` and `sb`
- `ra` rotate stack `a` (shift up all elements of stack `a` by 1, the first element becomes the last one)
- `rb` rotate stack `b`
- `rr`execute `ra` and `rb`
- `rra` reverse rotate `a` (shift down all elements of stack `a` by 1, the last element becomes the first one)
- `rrb` reverse rotate `b`
- `rrr` execute `rra` and `rrb`

#### Example

```console
Init a and b :
2
1
3
6
8
5
= =
a b
---------------------------------------
Exec sa :
1
2
3
6
8
5
= =
a b
---------------------------------------
Exec pb pb pb :
6 3
8 2
5 1
= =
a b
---------------------------------------
Exec rb :
6 2
8 1
5 3
= =
a b
---------------------------------------
Exec rra and rrb (equivalent to rrr):
5 3
6 2
8 1
= =
a b
---------------------------------------
Exec pa pa pa:
1
2
3
5
6
8
= =
a b
```

This project will help you learn about :

- the use of basic algorithms
- the use of sorting algorithms
- the use of stacks

#### The push_swap program

- You have to write a program named push_swap, which will receive as an argument the stack a formatted as a list of integers. The first integer should be at the top of the stack.
- The program must display the smallest list of instructions possible to sort the stack `a`, with the smallest number being at the top.
- Instructions must be separated by a `\n` and nothing else.
- The goal is to sort the stack with the minimum possible number of operations.
- In case of error, you must display `Error` followed by a `\n` on the standard error. Errors are understood as: some arguments aren’t integers and/or there are duplicates.
- In case of there are no arguments the program displays nothing (0 instructions).

```console
student$ ./push_swap "2 1 3 6 5 8"
sa
pb
pb
pb
sa
pa
pa
pa
student$ ./push_swap "0 one 2 3"
Error
student$ ./push_swap
student$
```

#### The checker program

- You have to write a program named checker, which will get as an argument the stack `a` formatted as a list of integers. The first argument should be at the top of the stack (be careful about the order). If no argument is given, checker stops and displays nothing.
- Checker will then read instructions on the standard input, each instruction will be followed by `\n`. Once all the instructions have been read, checker will execute them on the stack received as an argument.
- If after executing those instructions, stack `a` is actually sorted and `b` is empty, then checker must display "OK" followed by a `\n` on the standard output. In every other case, checker must display "KO" followed by a `\n` on the standard output.
- In case of error, you must display Error followed by a `\n` on the standard error. Errors include for example: some arguments are not integers, there are duplicates, an instruction don’t exist and/or is incorrectly formatted.
- In case of there are no arguments the program displays nothing.

```console
student$ ./checker "3 2 1 0"
sa
rra
pb
KO
student$  echo -e "rra\npb\nsa\n" | ./checker "3 2 one 0"
Error
student$ echo -e "rra\npb\nsa\nrra\npa"
rra
pb
sa
rra
pa
student$ echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
OK
student$ ./checker
student$
```

### Allowed packages

- Only the [standard go](https://golang.org/pkg/) packages are allowed

### Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices.en).
- It is recommended that the code should present a **test file**.
- The first executable file must be named **checker** and the second **push_swap**.
- You have to be able to handle the errors.

### Usage

```console
student$ ARG="4 67 3 87 23"; ./push_swap "$ARG" | wc -l
6
student$ ARG="4 67 3 87 23"; ./push_swap "$ARG" | ./checker "$ARG"
OK
```

If the program checker displays KO, it means that your **push_swap** came up with a list of instructions that doesn't sort the list.
