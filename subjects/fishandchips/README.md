## Fishandchips

### Instructions

Write a function called `Fishandchips()` Write a function that takes an integer and returns a string.

- If the number is divisible by 3, print `chips` followed by a newline `\n`.
- If the number is divisible by 2 and 3, print `fish and chips` followed by a newline `\n`.
- If is not divisible by any of 3 and 2 print newline `\n`.

### Expected function

```go
func Fishandchips(n int32) string {
}
```
### Usage

Here is a possible program to test your function :

```go
package main

func main() {
	args := []int32{0, 6, 33, -3, 5, 8}
	for i := 0; i < len(args); i++ {
		Fishandchips(args[i])
	}
}
```
And its output :

```go
fish and chips$
fish and chips$
chips$
chips$
$
$
```
