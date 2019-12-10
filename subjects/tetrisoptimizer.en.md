## tetrisoptimizre

### Objectives

Develop a program that receives only one argument, a text file which will contain a list of [Tetrominoes](https://en.wikipedia.org/wiki/Tetromino) to assemble them in order to create the smallest square possible.

### Instructions

- Your program will be tested using our solution, you don't need to submit your text files

The program must :

- Expect at least 2 Tetrominoes in the text file
- Identify each Tetromino in the solution, by assigning different letters to different Tetrominoes
- Assemble all of the Tetrominoes in order to create the smallest square possible
- Print the solution where Tetrominoes are placed on their most upper left position, if there are more than one solution it
- In case of bad format on the Tetrominoes or bad file format it should print `ERROR`

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
student@ubuntu:~/tetrisoptimizre$ cat -e sample.txt
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
student@ubuntu:~/tetrisoptimizre$ ./tetrisoptimizre sample.txt | cat -e
ABBBB.$
ACCCEE$
AFFCEE$
A.FFGG$
HHHDDG$
.HDD.G$
student@ubuntu:~/tetrisoptimizre$
```
