## invert tree

## **WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(`func main()`) even if empty.
- The function main declared needs to **also pass** the `Restrictions Checker`(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
- Every other rules are obviously the same than for a `program`.

### Instructions

Write a function that takes tree and inverts(flips) and returns it.

### Expected function and structure

```go
type TNode struct {
    Val     int
    Left    *TNode
    Right   *TNode
}

func InvertTree(root *TNode) *TNode {

}
```

Example:

```shell
Input:
      7
    /   \
   5     10
  / \    / \
 3   6  9  13

Output:
        7
     /    \
   10      5
  / \     / \
13   9   6   3
```
