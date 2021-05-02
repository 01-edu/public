## halfcontest

### Instructions

It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).
Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum" (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, "Lorem ipsum dolor sit amet..", comes from a line in section 1.10.32.

The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from "de Finibus Bonorum et Malorum" by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation by H. Rackham.

You are given 4 non-negative integers 'h1', 'm1', 'h2', and 'm2'. The contest starts at h1:m1 minutes and finishes at h2:m2 minutes. Your task is to find out when half of the contest will be over and return it in the format decribed in the example.
Contest cannot finish before it was started.

### Expected function

```go
func HalfContest(h1, m1, h2, m2 int) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"

	"piscine"
)

func main() {
	fmt.Println(piscine.HalfContest(1, 15, 3, 33))
	fmt.Println(piscine.HalfContest(10, 3, 11, 55))
	fmt.Println(piscine.HalfContest(9, 2, 11, 3))
}
```

And its output :

```console
$ go run .
224
1059
1002
$
```
