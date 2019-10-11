## printcomb

### Instructions

Write a [function](TODO-LINK) that prints in ascending order on a single line all unique combinations of three different digits so that the first digit is lower than the second and the second is lower than the third.

These combinations are separated by a comma and a space.

### Expected function

```go
func PrintComb() {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import piscine ".."

func main() {
	piscine.PrintComb()
}
```

This is the incomplete output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
012, 013, 014, 015, 016, 017, 018, 019, 023, ..., 689, 789
student@ubuntu:~/piscine-go/test$
```

`000` or `999` are not valid combinations because the digits are not different.

`987` should not be shown because the first digit is not less than the second.
