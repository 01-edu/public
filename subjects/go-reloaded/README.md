## go-reloaded

### Introduction

- Welcome back. Congratulations on your admission.

This subject is a selection of the best exercises of the piscine-go.
This selection was made to get back smoothly to programming in go after the pause period.
All the exercises have to be done and they must all pass each and every tests.

The goal of this project is to realize what you know and what you still need to practice on.
We strongly advise you to resist the temptation of looking at your old repository.

To really learn programming, you must practice programming, so copy and pasting some old code will not help the learning process.

One more detail. This time the project will be corrected by auditors. The auditors will be others students and you will be an auditor as well.

We advise you to create your own tests for yourself and for when you will correct your students.

---

## atoi

### Allowed functions

- `"--cast"`

### Instructions

- Write a [function](TODO-LINK) that simulates the behaviour of the `Atoi` function in Go. `Atoi` transforms a number represented as a `string` in a number represented as an `int`.

- `Atoi` returns `0` if the `string` is not considered as a valid number. For this exercise **non-valid `string` chains will be tested**. Some will contain non-digits characters.

- For this exercise the handling of the signs + or - **does have** to be taken into account.

- This function will **only** have to return the `int`. For this exercise the `error` result of atoi is not required.

### Expected function

```go
func Atoi(s string) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	student ".."
)

func main() {
	fmt.Println(student.Atoi("12345"))
	fmt.Println(student.Atoi("0000000012345"))
	fmt.Println(student.Atoi("012 345"))
	fmt.Println(student.Atoi("Hello World!"))
	fmt.Println(student.Atoi("+1234"))
	fmt.Println(student.Atoi("-1234"))
	fmt.Println(student.Atoi("++1234"))
	fmt.Println(student.Atoi("--1234"))
}
```

And its output :

```console
$ go run .
12345
12345
0
0
1234
-1234
0
0
$
```

---

## recursivepower

### Allowed functions

- `"--cast"`

### Instructions

Write an **recursive** function that returns the power of the `int` passed as parameter.

Negative powers will return `0`. Overflows do **not** have to be dealt with.

`for` is **forbidden** for this exercise.

### Expected function

```go
func RecursivePower(nb int, power int) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	student ".."
)

func main() {
	arg1 := 4
	arg2 := 3
	fmt.Println(student.RecursivePower(arg1, arg2))
}
```

And its output :

```console
$ go run .
64
$
```

---

## printcombn

### Allowed functions

- `"--cast"`
- `"github.com/01-edu/z01.PrintRune"`

### Instructions

- Write a function that prints all possible combinations of **n** different digits in ascending order.

- n will be defined as : 0 < n < 10

below are your references for the **printing format** expected.

- (for n = 1) '0, 1, 2, 3, ...8, 9'

- (for n = 3) '012, 013, 014, 015, 016, 017, 018, 019, 023,...689, 789'

### Expected function

```go
func PrintCombN(n int) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import student ".."

func main() {
	student.PrintCombN(1)
	student.PrintCombN(3)
	student.PrintCombN(9)
}
```

And its output :

```console
$ go run .
0, 1, 2, 3, 4, 5, 6, 7, 8, 9
012, 013, 014, 015, 016, 017, 018, ... 679, 689, 789
012345678, 012345679, ..., 123456789
$
```

---

## printnbrbase

### Allowed functions

- `"--cast"`
- `"github.com/01-edu/z01.PrintRune"`

### Instructions

Write a function that prints an `int` in a `string` base passed in parameters.

If the base is not valid, the function prints `NV` (Not Valid):

Validity rules for a base :

- A base must contain at least 2 characters.
- Each character of a base must be unique.
- A base should not contain `+` or `-` characters.

The function has to manage negative numbers. (as shown in the example)

### Expected function

```go
func PrintNbrBase(nbr int, base string) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"

	"github.com/01-edu/z01"
	student ".."
)

func main() {
	student.PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	student.PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	student.PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}
```

And its output :

```console
$ go run .
125
-1111101
7D
-uoi
NV
$
```

---

## doop

### Allowed functions

- `"--cast"`
- `"github.com/01-edu/z01.PrintRune"`
- `"os.*"`
- `".."`

### Instructions

Write a program that is called `doop`.

The program has to be used with three arguments:

- A value
- An operator
- Another value

You should use `int64`.

The following operators are considered valid: `+`, `-`, `/`, `*`, `%`.

In case of an invalid operator or overflow the programs prints `0`.

In case of an invalid number of arguments the program prints nothing.

The program has to handle the modulo and division operations by 0 as shown on the output examples below.

### Usage

```console
$ go run .
$ go run . 1 + 1 | cat -e
2$
$ go run . hello + 1 | cat -e
0$
$ go run . 1 p 1 | cat -e
0$
$ go run . 1 / 0 | cat -e
No division by 0$
$ go run . 1 % 0 | cat -e
No modulo by 0$
$ go run . 9223372036854775807 + 1
0
$ go run . -9223372036854775809 - 3
0
$ go run . 9223372036854775807 "*" 3
0
$ go run . 1 "*" 1
1
$ go run . 1 "*" -1
-1
$
```

---

## atoibase

### Allowed functions

- `"--cast"`
- `"make"`

### Instructions

Write a function that takes a `string` number and its `string` base in parameters and returns its conversion as an `int`.

If the base or the `string` number is not valid it returns `0`:

Validity rules for a base :

- A base must contain at least 2 characters.
- Each character of a base must be unique.
- A base should not contain `+` or `-` characters.

Only valid `string` numbers will be tested.

The function **does not have** to manage negative numbers.

### Expected function

```go
func AtoiBase(s string, base string) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	student ".."
)

func main() {
	fmt.Println(student.AtoiBase("125", "0123456789"))
	fmt.Println(student.AtoiBase("1111101", "01"))
	fmt.Println(student.AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(student.AtoiBase("uoi", "choumi"))
	fmt.Println(student.AtoiBase("bbbbbab", "-ab"))
}
```

And its output :

```console
$ go run .
125
125
125
125
0
$
```

---

## splitwhitespaces

### Allowed functions

- `"--cast"`
- `"make"`

### Instructions

Write a function that separates the words of a `string` and puts them in a `string` slice.

The separators are spaces, tabs and newlines.

### Expected function

```go
func SplitWhiteSpaces(s string) []string {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	student ".."
)

func main() {
	fmt.Println(student.SplitWhiteSpaces("Hello how are you?"))
}
```

And its output :

```console
$ go run .
[Hello how are you?]
$
```

---

## split

### Allowed functions

- `"--cast"`
- `"make"`

### Instructions

Write a function that simulates the behaviour of the `strings.Split` function in Go. `strings.Split` separates the words of a `string` and puts them in a `string` slice.
The separators are the characters of the separator string given in parameter.

### Expected function

```go
func Split(s, sep string) []string {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	student ".."
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Println(student.Split(s, "HA"))
}
```

And its output :

```console
$ go run .
[Hello how are you?]
$
```

---

## convertbase

### Allowed functions

- `"--cast"`
- `".."`

### Instructions

Write a function that returns the conversion of a `string` number from one `string` baseFrom to one `string` baseTo.

Only valid bases will be tested.

Negative numbers will not be tested.

### Expected function

```go
func ConvertBase(nbr, baseFrom, baseTo string) string {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	student ".."
)

func main() {
	result := student.ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
}
```

And its output :

```console
$ go run .
43
$
```

---

## rotatevowels

### Allowed functions

- `"os.*"`
- `"--cast"`
- `"github.com/01-edu/z01.PrintRune"`
- `"make"`

### Instructions

Write a **program** that checks the arguments for vowels.

- If the argument contains vowels (for this exercise `y` is not considered a vowel) the program has to **"mirror"** the position of the vowels in the argument (see the examples).
- If the number of arguments is less than 1, the program display a new line ("`\n`").
- If the arguments does not have any vowels, the program just prints the arguments.

Example of output :

```console
$ go run . "Hello World" | cat -e
Hollo Werld$
$ go run . "HEllO World" "problem solved"
Hello Werld problom sOlvEd
$ go run . "str" "shh" "psst"
str shh psst
$ go run . "happy thoughts" "good luck"
huppy thooghts guod lack
$ go run . "aEi" "Ou"
uOi Ea
$ go run .

$
```

---

## advancedsortwordarr

### Allowed functions

- `".."`
- `"len"`

### Instructions

Write a function `AdvancedSortWordArr` that sorts a `string` slice, based on the function `f` passed in parameter.

### Expected function

```go
func AdvancedSortWordArr(a []string, f func(a, b string) int) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	student ".."
)

func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	student.AdvancedSortWordArr(result, student.Compare)

	fmt.Println(result)
}
```

And its output :

```console
$ go run .
[1 2 3 A B C a b c]
$
```

---

## cat

### Allowed functions

- `"os.*"`
- `"--cast"`
- `"github.com/01-edu/z01.PrintRune"`
- `".."`
- `"ioutil.*"`
- `"io.*"`

### Instructions

Write a program that behaves like a simplified `cat` command.

- The `options` do not have to be handled.

- If the program is called without arguments it should take the standard input (stdin) and print it back on the standard output (stdout).

- In the program folder create two files named `quest8.txt` and `quest8T.txt`.

- Copy to the `quest8.txt` file the following sentence, with a new line ate the end of the file :

`"Programming is a skill best acquired by practice and example rather than from books" by Alan Turing`

- Copy to the `quest8T.txt` file the following sentence, with a new line ate the end of the file :

`"Alan Mathison Turing was an English mathematician, computer scientist, logician, cryptanalyst. Turing was highly influential in the development of theoretical computer science, providing a formalisation of the concepts of algorithm and computation with the Turing machine, which can be considered a model of a general-purpose computer. Turing is widely considered to be the father of theoretical computer science and artificial intelligence."`

- In case of error the program should print the error.

- The program must be submitted inside a folder named `cat`.

```console
$ go run . abc
ERROR: abc: No such file or directory
$ go run . quest8.txt
"Programming is a skill best acquired by pratice and example rather than from books" by Alan Turing
$ go run .
Hello
Hello
^C
$ go run . quest8.txt quest8T.txt
"Programming is a skill best acquired by pratice and example rather than from books" by Alan Turing
"Alan Mathison Turing was an English mathematician, computer scientist, logician, cryptanalyst. Turing was highly influential in the development of theoretical computer science, providing a formalisation of the concepts of algorithm and computation with the Turing machine, which can be considered a model of a general-purpose computer. Turing is widely considered to be the father of theoretical computer science and artificial intelligence."
$
```

---

## ztail

### Allowed functions

- `"os.*"`
- `"--cast"`
- `"github.com/01-edu/z01.PrintRune"`
- `"append"`
- `"fmt.Printf"`
- `".."`

### Instructions

Write a program called `ztail` that has the same behaviour as the system command `tail`, but which takes at least one file as argument.

- The only option to be handled is `-c`. This option will be used in all tests.

- For this program the "os" package can be used.

- For the program to pass the tests the convention for the return code of program in Unix systems should be followed.

- Handle the errors and output the same error messages as `tail`.

- For more information consult the man page for `tail`.

### Note:

This program is gonna be tested against `tail`, be sure to check all the different error messages with different permutations of the arguments.

---

## activebits

### Allowed functions

- `"--cast"`

### Instructions

Write a function, `ActiveBits`, that returns the number of active `bits` (bits with the value 1) in the binary representation of an integer number.

### Expected function

```go
func ActiveBits(n int) uint {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	student ".."
)

func main() {
	nbits := student.ActiveBits(7)
	fmt.Println(nbits)
}
```

And its output :

```console
$ go run .
3
$
```

---

## sortlistinsert

### Allowed functions

- `"--no-array"`

### Instructions

Write a function `SortListInsert` that inserts `data_ref` in the linked list `l` while keeping the list sorted in ascending order.

- During the tests the list passed as an argument will be already sorted.

### Expected function and structure

```go
func SortListInsert(l *NodeI, data_ref int) *NodeI{

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"

	student ".."
)

func PrintList(l *student.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {

	var link *student.NodeI

	link = listPushBack(link, 1)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)

	PrintList(link)

	link = student.SortListInsert(link, -2)
	link = student.SortListInsert(link, 2)
	PrintList(link)
}
```

And its output :

```console
$ go run .
1 -> 4 -> 9 -> <nil>
-2 -> 1 -> 2 -> 4 -> 9 -> <nil>
$
```

---

## sortedlistmerge

### Allowed functions

- `"--no-array"`

### Instructions

Write a function `SortedListMerge` that merges two lists `n1` and `n2` in ascending order.

### Expected function and structure

```go
func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"

	student ".."
)

func PrintList(l *student.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *student.NodeI
	var link2 *student.NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)

	PrintList(student.SortedListMerge(link2, link))
}
```

And its output :

```console
$ go run .
-2 -> 3 -> 5 -> 7 -> 9 -> <nil>
$
```

---

## listremoveif

### Allowed functions

- `"--no-array"`

### Instructions

Write a function `ListRemoveIf` that removes all elements that are equal to the `data_ref` in the argument of the function.

### Expected function and structure

```go
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"

	student ".."
)

func PrintList(l *student.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func main() {
	link := &student.List{}
	link2 := &student.List{}

	fmt.Println("----normal state----")
	student.ListPushBack(link2, 1)
	PrintList(link2)
	student.ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	fmt.Println("----normal state----")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "Hello")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "There")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "How")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "are")
	student.ListPushBack(link, "you")
	student.ListPushBack(link, 1)
	PrintList(link)

	student.ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
}
```

And its output :

```console
$ go run .
----normal state----
1 -> <nil>
------answer-----
<nil>

----normal state----
1 -> Hello -> 1 -> There -> 1 -> 1 -> How -> 1 -> are -> you -> 1 -> <nil>
------answer-----
Hello -> There -> How -> are -> you -> <nil>
$
```

---

## btreetransplant

### Instructions

In order to move subtrees around within the binary search tree, write a function, `BTreeTransplant`, which replaces the subtree started by `node` with the node `rplc` in the tree given by `root`.

### Expected function

```go
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	student ".."
)

func main() {
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	node := student.BTreeSearchItem(root, "1")
	replacement := &student.TreeNode{Data: "3"}
	root = student.BTreeTransplant(root, node, replacement)
	student.BTreeApplyInorder(root, fmt.Println)
}
```

And its output :

```console
$ go run .
3
4
5
7
$
```

---

## btreeapplybylevel

### Instructions

Write a function, `BTreeApplyByLevel`, that applies the function given by `fn` to each node of the tree given by `root`.

### Expected function

```go
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error))  {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	student ".."
)

func main() {
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	student.BTreeApplyByLevel(root, fmt.Println)
}
```

And its output :

```console
$ go run .
4
1
7
5
$
```

---

## btreedeletenode

### Instructions

Write a function, `BTreeDeleteNode`, that deletes `node` from the tree given by `root`.

The resulting tree should still follow the binary search tree rules.

### Expected function

```go
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	student ".."
)

func main() {
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	node := student.BTreeSearchItem(root, "4")
	fmt.Println("Before delete:")
	student.BTreeApplyInorder(root, fmt.Println)
	root = student.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	student.BTreeApplyInorder(root, fmt.Println)
}
```

And its output :

```console
$ go run .
Before delete:
1
4
5
7
After delete:
1
5
7
$
```
