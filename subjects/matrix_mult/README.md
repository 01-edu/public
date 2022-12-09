## matrix_mult

### Instructions

Implement the methods:

- `number_of_cols`: which returns the number of columns in the matrix.
- `number_of_rows`: which returns the number of rows in the matrix.
- `row`: which returns the `n`th row in the matrix.
- `col`: which returns the `n`th column in the matrix.

Define the matrix multiplication by implementing the `std::ops::Mul` for the type matrix

### Expected Functions

```rust
impl Matrix<T> {
	pub fn number_of_cols(&self) -> usize {
	}

	pub fn number_of_rows(&self) -> usize {
	}

	pub fn row(&self, n: usize) -> Vec<T> {
	}

	pub fn col(&self, n: usize) -> Vec<T> {
	}
}

impl Mul for Matrix<T> {
}
```

### Usage

Here is a program to test your function.

```rust
use matrix_mult::*;

fn main() {
	let matrix: Matrix<u32> = Matrix(vec![vec![3, 6], vec![8, 0]]);
	println!("{:?}", matrix.col(0));
	println!("{:?}", matrix.row(1));

	let matrix_1: Matrix<u32> = Matrix(vec![vec![0, 1], vec![0, 0]]);
	let matrix_2: Matrix<u32> = Matrix(vec![vec![0, 0], vec![1, 0]]);
	let mult = matrix_1.clone() * matrix_2.clone();
	println!("{:?}", mult);
	println!("{:?}", matrix_1.number_of_cols());
	println!("{:?}", matrix_2.number_of_rows());
}
```

And its output

```console
$ cargo run
[3, 8]
[8, 0]
Some(Matrix([[1, 0], [0, 0]]))
2
2
$
```
