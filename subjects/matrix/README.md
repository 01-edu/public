## matrix

### Instructions

Define a data structure to represent a matrix of any size and implement the basic operations for this. The next steps need to be followed:

- You can use a 2 dimensional Vec<T>'s. We will consider a matrix as a rectangular arrangements of scalars.

- You have to use the definition of scalars done in the exercise: `lalgebra_scalar`

- Then define the associated function `identity` that returns the identity matrix of size n

- Finally, define the associated function `zero` that returns a matrix of size `row x col` with all the positions filled by zeroes

### Notions

[Traits](https://doc.rust-lang.org/book/ch19-03-advanced-traits.html)

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
Matrix([[0, 0, 0, 0], [0, 0, 0, 0], [0, 0, 0, 0]])
$
```
