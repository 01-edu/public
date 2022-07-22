## StrisNegative

Write a function called `StrisNegative()` that defines whether a number is negative or positive.
- Your function prints `P` if the number is positive
- Your function prints `F` if the number is negative
- If the number is zero, print `0`.
- If it's not a number, print `('\n')`.
- Your program should always print `('\n')` at the end of the output.

### Expected function

```go
func StrisNegative(str string) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

func main() {
	StrisNegative("585")
	StrisNegative("-58")
	StrisNegative("55s44")
	StrisNegative("101-1331")
	StrisNegative("5544-")
}
```

And its output :

```console
$ go run .
P
N
!
!
!
```
