## anagram

##**WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(```func main()```) even if empty.
- The function main declared needs to **also pass** the `Restrictions Checker`(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
- Every other rules are obviously the same than for a `program`.

### Instructions

Write a function that returns `true` if two strings are anagrams, otherwise returns `false`.
Anagram is a string made by using the letters of another string in a different order.

Only lower case characters will be given.

### Expected function

```go
func IsAnagram(str1, str2 string) bool {

}
```

### Usage

Here is a possible [program] to test your function:

```go
package main

import (
    "fmt"
)

func main() {
    test1 := IsAnagram("listen", "silent")
    fmt.Println(test1)

    test2 := IsAnagram("alem", "school")
    fmt.Println(test2)

    test3 := IsAnagram("neat", "a net")
    fmt.Println(test3)

    test4 := IsAnagram("anna madrigal", "a man and a girl")
    fmt.Println(test4)
}
```

And its output:

```console
student@ubuntu:~/test$ go build
student@ubuntu:~/test$ ./test
true
false
true
true
student@ubuntu:~/test$
```
