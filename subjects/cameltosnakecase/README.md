## Camel-to-snake-case 

### Instructions

Write a function that converts a string from `camelCase` to `snake_case`.

For this exercise you need to know that camelCase has two different writing alternatives that will be accepted:

- lowerCamelCase
- UpperCamelCase

Some rules for writing in camelCase:

- The word does not end on a capitalized letter (CamelCasE).
- No two capitalized letters shall follow directly each other (CamelCAse).
- Numbers or punctuation are not allowed in the word anywhere (camelCase1).

Here are some rules for you to follow:

- If the string is empty, return an empty string.
- If the string is not `camelCase`, return the string unchanged.
- If the string is `camelCase`, return the `snake_case` version of the string.


### Expected function 

```go
func CamelToSnakeCase(s string) string{
    //Your code here
}
```

### Usage 

Here is a possible program to test your function:

```go
package main

import "fmt"

func main() {
    fmt.Println(CamelToSnakeCase("HelloWorld")) 
    fmt.Println(CamelToSnakeCase("helloWorld"))
    fmt.Println(CamelToSnakeCase("camelCase"))
    fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
    fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
    fmt.Println(CamelToSnakeCase("132322"))
}
```

and the output should be:

```console
$ go run .
Hello_World
hello_World
camel_Case
CAMELtoSnackCASE
camel_To_Snake_Case
132322
```

