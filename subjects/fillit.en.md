## fill it

### Objectives

Develop a program that receives only one argument, a txt file which will contain a list of [Tetrominoes](https://en.wikipedia.org/wiki/Tetromino) to assemble them in order to create the smallest square possible.

### Instructions

All files referring to the project, must be in the same folder.

The program must :

- Be written in Go
- Compile successfully
- Assemble all of the Tetrominoes in order to create the smallest square possible
- Make the smallest square with the Tetrominoes
- Identify each Tetromino in the solution, by assigning different letters to different Tetrominoes
- Expect at least 2 tetrominoes in the text file

#### Example of a text File

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

- If it isn't possible to form a complete square, the program should leave spaces between the Tetrominoes. For example:

```console
ABB.
ABB.
A...
A...
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
