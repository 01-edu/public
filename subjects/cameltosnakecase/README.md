## Camel-to-snake-case 

### Instructions
Write a function that converts a string from `CamelCase` to `snake_case`.

Camel case is the practice of writing phrases without spaces or punctuation, it indicates the separation of words with a single capitalized letter. Snake case is a style of writing in which each space is replaced by an underscore (_) character.

Here are some rules for you to follow:

- If the string is empty, return an empty string.
- If the string is not `CamelCase`, return the string unchanged.
- If the string is `CamelCase`, return the `snake_case` version of the string.

Basic `CamelCase` Capitalization Rules:

- The first letter must be capitalized.
- The word must not have two capitalized letters together (CamelCAse) not end with a capitalized letter (CamelCasE).
- No numbers or punctuations are allowed in the word at any place (CamelCase1more).


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
    fmt.Println(CamelToSnakeCase("CamelToSnakeCase"))
    fmt.Println(CamelToSnakeCase("132322"))
}
```

and the output should be:

```console
$ go run .
Hello_World
helloWorld
Camel_To_Snake_Case
132322
```

