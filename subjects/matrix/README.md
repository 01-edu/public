## matrix

### Instructions

Define a data structure to represent a matrix of any size and implement some basic operations.
In other words, create a simple way to represent a table with any number of rows and columns.

Think of this table as a grid of numbers. You can use a 2D vector to store these numbers.

Make sure to use the definition of numbers (scalars) from the [lalgebra_scalar](../lalgebra_scalar/README.md) exercise.

Implement the following associated functions:
- `new`: which returns a `1x1` matrix (just 1 row and 1 column).
- `identity`: Creates a special type of table called an "identity matrix" with n rows and n columns. It has 1s on the diagonal and 0s everywhere else.
- `zero`: which returns a matrix of specified number of `rows` and `columns`, and fills it with zeros.
 
### Expected Functions and Structure

```rust
pub struct Matrix<T>(pub Vec<Vec<T>>);

impl <T: Scalar<Item = T>> Matrix<T> {
	pub fn new() -> Matrix<T> {
	}

	pub fn zero(row: usize, col: usize) -> Matrix<T> {
	}

	pub fn identity(n: usize) -> Matrix<T> {
	}
}
```

### Usage

Here is a program to test your function.

```rust
use matrix::*;

fn main() {
	let m: Matrix<u32> = Matrix(vec![vec![0, 0, 0, 0], vec![0, 0, 0, 0], vec![0, 0, 0, 0]]);
	println!("{:?}", m);
	println!("{:?}", Matrix::<i32>::identity(4));
	println!("{:?}", Matrix::<f64>::zero(3, 4));
}
```

And its output:

```console
$ cargo run
Matrix([[0, 0, 0, 0], [0, 0, 0, 0], [0, 0, 0, 0]])
Matrix([[1, 0, 0, 0], [0, 1, 0, 0], [0, 0, 1, 0], [0, 0, 0, 1]])
Matrix([[0.0, 0.0, 0.0, 0.0], [0.0, 0.0, 0.0, 0.0], [0.0, 0.0, 0.0, 0.0]])
$
```

### Notions

- [Traits](https://doc.rust-lang.org/book/ch19-03-advanced-traits.html)
