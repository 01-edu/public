## cameltosnakecase

### Instructions

Write a function that converts a `string` from `camelCase` to `snake_case`.

- If the `string` is empty, return an empty `string`.
- If the `string` is not `camelCase`, return the `string` unchanged.
- If the `string` is `camelCase`, return the `snake_case` version of the `string`.

For this exercise you need to know that `camelCase` has two different writing alternatives that will be accepted:

- lowerCamelCase
- UpperCamelCase

Rules for writing in `camelCase`:

- The word does not end on a capitalized letter (CamelCasE).
- No two capitalized letters shall follow directly each other (CamelCAse).
- Numbers or punctuation are not allowed in the word anywhere (camelCase1).

### Expected function

```go
func CamelToSnakeCase(s string) string{

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
```

And its output:

```console
$ go run .
Hello_World
hello_World
camel_Case
CAMELtoSnackCASE
camel_To_Snake_Case
hey2
```
