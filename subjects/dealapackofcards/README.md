## dealapackofcards

### Instructions

Deal a pack of `12` cards evenly between 4 players, `Player 1`, `Player 2`, `Player 3` and `Player 4`.

Write a function `DealAPackOfCards()` that takes the argument, `deck`, as a slice of `int` and prints the desired output.

- Each player must be printed in a different line.

### Expected function

```go
func DealAPackOfCards (deck []int) {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "piscine"

func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	piscine.DealAPackOfCards(deck)
}
```

And its output:

```console
$ go run . | cat -e
Player 1: 1, 2, 3$
Player 2: 4, 5, 6$
Player 3: 7, 8, 9$
Player 4: 10, 11, 12$
$
```
