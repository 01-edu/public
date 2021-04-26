## volumechanger

### Instructions

It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).
Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum" (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, "Lorem ipsum dolor sit amet..", comes from a line in section 1.10.32.

Dagar loves watching TV. He always makes the volume to be **b**. Today he is angry, because he found out that someone changed the volume, and now, the volume is **a**.
His remote controller can change the volume by: -5, -2, -1, +1, +2, +5.
Please help Dagar, to minimize number of button presses, and calculate how many times he should press buttons of remote controller so that it changes from **a** to **b**.
There are **t** requests in total.
Input and output should be displayed in standard input and output respectively.

### Expected function

```go
func Volumechanger(a, b int) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
	fmt.Println(Volumechanger(4, 0))
	fmt.Println(Volumechanger(5, 14))
	fmt.Println(Volumechanger(3, 9))
}
```

And its output :

```console
student@ubuntu:~/volumechanger/test$ go build
student@ubuntu:~/volumechanger/test$ ./test
2
3
2
student@ubuntu:~/volumechanger/test$
```
