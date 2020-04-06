## anagram

## **WARNING! VERY IMPORTANT!**

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
package piscine

func IsAnagram(str1, str2 string) bool {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	piscine ".."
	"fmt"
)

func main() {
	test1 := piscine.IsAnagram("listen", "silent")
	fmt.Println(test1)

	test2 := piscine.IsAnagram("alem", "school")
	fmt.Println(test2)

	test3 := piscine.IsAnagram("neat", "a net")
	fmt.Println(test3)

	test4 := piscine.IsAnagram("anna madrigal", "a man and a girl")
	fmt.Println(test4)
}
```

Its output:

```bash
$> go build
$> ./main
true
false
true
true
```
