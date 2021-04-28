## tetrisoptimizer

### Objectives

Develop a program that receives only one argument, a path to a text file which will contain a list of [tetrominoes](https://en.wikipedia.org/wiki/Tetromino) to assemble them in order to create the smallest square possible.

### Allowed packages

- Only the [standard go](https://golang.org/pkg/) packages are allowed

### Instructions

The program must :

- Compile successfully
- Assemble all of the tetrominoes in order to create the smallest square possible
- Identify each tetromino in the solution by printing them with uppercase latin letters (`A` for the first one, `B` for the second, etc)
- Expect at least one tetromino in the text file
- In case of bad format on the tetrominoes or bad file format it should print `ERROR`
- The project must be written in **Go**.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices/).
- It is recommended that the code should present a **test file**.

This project will help you learn about:

- The use of algorithms
- Reading from files

#### Example of a text File

```
#...
#...
#...
#...

....
....
..##
..##
```

- If it isn't possible to form a complete square, the program should leave spaces between the tetrominoes. For example:

```
ABB.
ABB.
A...
A...
```

## Usage

```console
$ cat -e sample.txt
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
$ go run . sample.txt | cat -e
ABBBB.$
ACCCEE$
AFFCEE$
A.FFGG$
HHHDDG$
.HDD.G$
$
```
