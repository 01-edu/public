# podiumposition

### Instructions

A F1 race just finished and the commentator is calling the finishing positions incorrectly.
Help to fix this before the contestants arrive at the podium by providing the commentator with the correct podium position.

Write a function `PodiumPosition` that takes an array of arrays of type `string` and returns the competitor positions correctly.

### Expected function

```go
func PodiumPosition(podium[][] string) [][]string {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"

func main() {
    position := [][]string{{"4th Place"}, {"3rd Place"}, {"2nd Place"}, {"1st Place"}}
    	fmt.Println(PodiumPosition(position))
}
```

And its output:

```go
$ go run . | cat -e
[[1st Place] [2nd Place] [3rd Place] [4th Place]]$
```
