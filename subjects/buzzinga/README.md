# buzzinga

### Instructions
- Write a function named `buzZinga()` that takes a number as an argument and prints the following:
    - If the number is divisible by 3, print `Buz` followed by ('`\n`')
    - If the number is divisible by 5, print `Zinga` followed by ('`\n`')
    - If the number is divisible by 3 and 5, print `BuzZinga` followed by ('`\n`')
    - If the number is not divisible by 3 or 5, print `*` followed by ('`\n`')
### Expected function
```go
func buzZinga(number int) {

}
```
### Usage
```go
package main

func main() {
	n := int32(15)
	buzZinga(n)
}
```
And its output :

```go
$ go run . | cat -e
*$
*$
Buz$
*$
Zinga$
Buz$
*$
*$
Buz$
Zinga$
*$
Buz$
*$
*$
BuzZinga$
```