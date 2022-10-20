## podiumposition

### Instructions

A F1 race just finished and the commentator is calling the finishing positions incorrectly.
Help to fix this before the contestants arrive at the podium by providing the commentator with the correct podium position.

Write a function `PodiumPosition` that takes a slice of slices of type `string` and returns the competitor positions correctly.

### Expected function

```go
func PodiumPosition(podium [][]string) [][]string {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
    "fmt"
    "piscine"
)

func main() {

    p4 := []string{"4th Place"}
    p3 := []string{"3rd Place"}
    p2 := []string{"2nd Place"}
    p1 := []string{"1st Place"}

    position := [][]string{p4, p3, p2, p1}
    fmt.Println(piscine.PodiumPosition(position))
}
```

And its output:

```console
$ go run . | cat -e
[[1st Place] [2nd Place] [3rd Place] [4th Place]]$
```
