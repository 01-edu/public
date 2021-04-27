## matrix_transposition

### Instructions

- Define the structure matrix as a tuple of 2 tuples of 2 `i32`'s.

- Define a **function** that calculates the transpose matrix of a 2x2 matrix.

- Note:

  - The transpose of a matrix `A` is the matrix `A'` where `A'`'s columns are `A`'s row and the rows are the columns:

Example:

```
( a b )   __ transposition __>   ( a c )
( c d )                          ( b d )
```

- Matrix must implement Debug, PartialEq and Eq. You can use derive.

- Remember that you are defining a library so you have to make public the elements that are going to be called from an external crate.

### Notions

- [Defining a struct](https://doc.rust-lang.org/stable/book/ch05-01-defining-structs.html)

- [The Tuple Type](https://doc.rust-lang.org/stable/book/ch03-02-data-types.html?highlight=accessing%20a%20tuple#compound-types)

- [Tuples](https://doc.rust-lang.org/rust-by-example/primitives/tuples.html)

- [Tuple Structs without Named Fields](https://doc.rust-lang.org/stable/book/ch05-01-defining-structs.html?highlight=tuple#using-tuple-structs-without-named-fields-to-create-different-types)

- [Adding Useful Functionality with Derived Traits](https://doc.rust-lang.org/stable/book/ch05-02-example-structs.html?highlight=debug%20deriv#adding-useful-functionality-with-derived-traits)

- [Chapter 7](https://doc.rust-lang.org/stable/book/ch07-03-paths-for-referring-to-an-item-in-the-module-tree.html)

### Expected Function

```rust
pub fn transpose(m: Matrix) -> Matrix {
}
```

### Usage

Here is a posible program to test your function

```rust
use matrix_transposition::transpose;
use matrix_transposition::Matrix;

fn main() {
    let matrix = Matrix((1, 3), (4, 5));
    println!("Original matrix {:?}", matrix);
    println!("Transpose matrix {:?}", transpose(matrix));
}
```

And its output:

```console
student@ubuntu:~/matrix_transposition/test$ cargo run
Original matrix Matrix((1, 3), (4, 5))
Transpose matrix Matrix((1, 4), (3, 5))
student@ubuntu:~/matrix_transposition/test$
```
