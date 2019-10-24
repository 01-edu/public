## foreachprog

##**WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(```func main()```) even if empty.
- The function main declared needs to **also pass** the `Restrictions Checker`(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
- Every other rules are obviously the same than for a `program`.

### Instructions

Write a function `ForEach` that, for an `int` array, applies a function on each elements of that array.

### Expected function

```go
func ForEach(f func(int), arr []int) {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, arr)
}

func PrintNbr(int) {

}
```

And its output :

```console
student@ubuntu:~/test$ go build
student@ubuntu:~/test$ ./test
123456
student@ubuntu:~/test$
```
