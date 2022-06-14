## Camel-to-snake-case 

### Instructions

Write a function that converts a string from a camel case to a snake case.
- If the string is empty, return an empty string.
- If the string is not a camel case, return the string unchanged.
- If the string is a camel case, return the snake case version of the string.

### Expected function 
```go
func CamelToSnakeCase(s string) string{
    //Your code here
}
```

### Usage 

```go
package main

import "fmt"

func main() {
    fmt.Println(CamelToSnakeCase("HelloWorld")) 
    fmt.Println(CamelToSnakeCase("heloWorld"))
    fmt.Println(CamelToSnakeCase("CamelToSnakeCase"))
    fmt.Println(CamelToSnakeCase("132322"))
    fmt.Println(CamelToSnakeCase(""))
}
```

and the output should be:

```console
$ go run . | cat -e
Hello_World$
helo_World$
Camel_To_Snake_Case$
132322$
$
```

