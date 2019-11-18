## fiil it

### Objectives

Develop a program that receives only one parameter, a txt file which will contain a list of [**Tetrominos**](https://en.wikipedia.org/wiki/Tetromino) to assemble them in order to create the smallest square possible.

### Instructions

All files referring to the project, must be in the same folder.

The program must:

- Be written in **GO**
- Be compiled to be runned
- Assemble all of the **Tetrominos** in order to create the smallest square possible
- Show an error if the input doesn't correspond to a classic **Tetromino** piece
- Not allow the rotation of the **Tetrominos**
- Organize the **Tetrominos** by their apparition order in the file
- Identify each **Tetromino** in the solution, by assigning different charactes to different **Tetrominos**
- Have at least 2 **Tetrominos** in the text file

Some notes:

#### Example of a text File:

```console
#...
#...
#...
#...

....
....
..##
..##
```

- If it isn't possible to form a perfect square, the program should leave spaces between the **Tetrominos**. For example:

```console
ABB.
ABB.
A...
A...
```



- Among all the possible candidates for the smallest square, we only accpet the ones where **Tetrominos** is placed on their most upper-left position. In the example below, the program would choose option a).

```console
a) ABB.  b)A...  c)A...
   ABB.    ABB.    A...   ...
   A...    ABB.    ABB.
   A...    A...    ABB.
```

## Usage

```
student@ubuntu:~/fill-it$ cat -e sample.txt
...#$
...#$
...#$
...#$
$
....$
....$
....$
####$
$
.###$
...#$
....$
....$
$
....$
..##$
.##.$
....$
$
....$
.##.$
.##.$
....$
$
....$
....$
##..$
.##.$
$
##..$
.#..$
.#..$
....$
$
....$
###.$
.#..$
....$
student@ubuntu:~/fill-it$ ./fillit sample.txt | cat -e
ABBBB.$
ACCCEE$
AFFCEE$
A.FFGG$
HHHDDG$
.HDD.G$
student@ubuntu:~/fill-it$
```