## leap-year

### Instructions

Write a function named `LeapYear(int)` that takes a year as a parameter and returns true if the year is a leap year and false otherwise.
- A leap year is a year divisible by 4, but not by 100.
- A leap year is also divisible by 400.
- If the number is not positive, return false

### Expected function

```go
func LeapYear(year int)bool{
  // your code here
}
```
### Usage

Here is a possible program to test your function:

```go 
package main

import "fmt"

func main(){
  fmt.Println(LeapYear(2020))
  fmt.Println(LeapYear(2021))
  fmt.Println(LeapYear(2022))
  fmt.Println(LeapYear(-10))
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
