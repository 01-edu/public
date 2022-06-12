## matrix

### Instructions

Define a data structure to represent a matrix of any size and implement some basic operations.

We will consider a matrix as a rectangular arrangements of scalars. You can represent this as a 2 dimensional vector`. You will use the definition of scalars from the [lalgebra_scalar](../lalgebra_scalar/README.md) exercise.

Implement the following associated functions:
- `new`: which returns a matrix of size `1 x 1`.
- `identity`: which returns the identity matrix of size n.
- `zero`: which returns a matrix of size `row` x `col` with all the positions filled by zeros.
 
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
