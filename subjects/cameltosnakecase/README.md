## Camel-to-snake-case 

### Instructions
Write a function that converts a string from `camelCase` to `snake_case`.

Lower camel case is the practice of writing phrases without spaces or punctuation, it indicates the separation of one word or more with a single capitalized letter between each word except the first one. Snake case is a style of writing in which each space is replaced by an underscore (_) character.

Here are some rules for you to follow:

- If the string is empty, return an empty string.
- If the string is not `camelCase`, return the string unchanged.
- If the string is `camelCase`, return the `snake_case` version of the string.

Basic `lowerCamelCase` Capitalization Rules:

- The first letter should be lower case.
- No numbers or punctuations are allowed in the word at any place (camelCase1).


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
    fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
    fmt.Println(CamelToSnakeCase("132322"))
}
```

and the output should be:

```console
$ go run .
HelloWorld
hello_World
camel_Case
camel_To_Snake_Case
132322
```

