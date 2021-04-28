## sametree

### Instructions

Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

Write a function, `IsSameTree`, that returns `bool`.

### Expected function

```go
type TreeNodeL struct {
    Left    *TreeNodeL
    Val     int
    Right   *TreeNodeL
}


func IsSameTree(p *TreeNodeL, q *TreeNodeL) bool {

}
```

Example 1:

Input:

          1
         / \
        2   3

       [1,2,3]

          1
         / \
        2   3

       [1,2,3]

Output: true

Input:

           1
          /
         2

      [1,2]

            1
             \
              2

      [1,null,2]

Output: false

Input:

```

           1
          / \
         2   1

        [1,2,1]

            1
           / \
          1   2

         [1,1,2]
```

Output: false

### Usage

Here is a possible program to test your function :

```go
package main

func main() {
	t1 := NewRandTree()
	t2 := NewRandTree()

	fmt.Println(IsSameTree(t1, t2))
}
```

### Output

```console
$ go run .
true
$
```
