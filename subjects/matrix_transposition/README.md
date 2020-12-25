## matrix_transposition

### Instructions

- Define the structure matrix as a tuple of tuples of `i32`'s

- Define a function that calculate the transpose matrix of a 2x2 matrix.

- Note:

	- The transpose of a matrix `A` is the matrix `A'` where `A'`'s columns are `A`'s row and the rows are the columns:

Example:

```
( a b )     __ transposition __> ( a d )
( c d )                          ( b d )
```

- Matrix must implement Debug, PartialEq and Eq. You can use derive

- Remember that you're defining a library so you have to make public the elements that are going to be called from an external crate.

### Notions

[Chapter 7]( https://doc.rust-lang.org/stable/book/ch07-03-paths-for-referring-to-an-item-in-the-module-tree.html )

### Expected Function

```rust
pub fn transpose(m: Matrix) -> Matrix {
}
```

### Usage

Here is a posible program to test your function

```rust
fn main() {
    let matrix = Matrix((1, 3), (4, 5));
    println!("Original matrix {:?}", matrix);
    println!("Transpose matrix {:?}", transpose(matrix));
}
```

And it's output:

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
Original matrix Matrix((1, 3), (4, 5))
Transpose matrix Matrix((1, 4), (3, 5))
student@ubuntu:~/[[ROOT]]/test$
```
