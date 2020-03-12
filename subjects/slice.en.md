## slice

## **WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(```func main()```) even if empty.
- The function main declared needs to **also pass** the `Restrictions Checker`(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
- Every other rules are obviously the same than for a `program`.

### Instructions

Write a **program** that replicates the javascript functions `slice`.

The program receives an array of strings and one or more integers, and returns an array of strings. The returned array is part of the received one but cut from the position indicated in the first int, until the position indicated by the second int. 

In case there only exists one int, the resulting array begins in the position indicated by the int and ends at the end of the received array.  

The ints can be lower than 0.

### Expected function

```go
func Slice(arr []string, nbrs... int) []string{

}
```


### Usage 

Here is a possible program to test your function :

```go
package main

import "fmt"

func main(){
    arr := []string{"coding", "algorithm", "ascii", "package", "golang"}
    fmt.Println(arr, 1)
    fmt.Println(arr, 2, 4)
    fmt.Println(arr, -3)
    fmt.Println(arr, -2, -1)
    fmt.Println(arr, 2, 0)
}
```

```console
student@ubuntu:~/student/test$ go build
student@ubuntu:~/student/test$ ./test
[algorithm ascii package golang]
[ascii package]
[algorithm ascii package golang]
[package]
[]
```