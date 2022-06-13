## leap-year

### Instructions

Write a function named `isLeap(int)` that takes a year as a parameter and returns true if the year is a leap year and false otherwise.
- A leap year is a year divisible by 4, but not by 100.
- A leap year is also divisible by 400.
- If the number is not positive, return false

### Expected function

```go
func isLeap(year int)bool{
  // your code here
}
```
### Usage

```go 
package main

import "fmt"

func main(){
  fmt.Println(isLeap(2020))
  fmt.Println(isLeap(2021))
  fmt.Println(isLeap(2022))
  fmt.Println(isLeap(-10))
}
```
and the output should be:

``` console
$ go run . | cat -e
true$
false$
false$
false$
```
