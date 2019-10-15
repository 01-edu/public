## raid2

### Instructions

-   Create a program that resolves a sudoku.

-   A valid sudoku has only one possible solution.

### Usage

#### Example 1:

Example of output for one valid sudoku :

```console
student@ubuntu:~/piscine-go/raid2$ go build
student@ubuntu:~/piscine-go/raid2$ ./raid2 ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
3 9 6 2 4 5 7 8 1$
1 7 8 3 6 9 5 2 4$
5 2 4 8 1 7 3 9 6$
2 8 7 9 5 1 6 4 3$
9 3 1 4 8 6 2 7 5$
4 6 5 7 2 3 9 1 8$
7 1 2 6 3 8 4 5 9$
6 5 9 1 7 4 8 3 2$
8 4 3 5 9 2 1 6 7$
student@ubuntu:~/piscine-go/raid2$
```

#### Example 2:

Examples of output for invalid input or sudokus :

```console
student@ubuntu:~/piscine-go/raid2$ ./raid2 1 2 3 4 +" | cat -e
Error$
student@ubuntu:~/piscine-go/raid2$ ./raid2 | cat -e
Error$
student@ubuntu:~/piscine-go/raid2$ ./raid2 ".96.4...1" "1...6.1.4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e"
Error$
student@ubuntu:~/piscine-go/raid2$
```
