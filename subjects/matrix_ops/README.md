## matrix_ops

### Instructions

In this exercise, you will define some basic matrix operations, Implement traits for `Add` and `Sub`

Remember that two matrices can only be added or subtracted if they have the same dimensions. Therefore, you must handle the possibility of failure by returning an `Option<T>`.

You will be reusing your `Matrix` and `Scalar` structures defined in the [matrix](../matrix/README.md) and [lalgebra_scalar](../lalgebra_scalar/README.md) exercises.

### Expected Function

```rust
impl Add for Matrix {

}

impl Sub for Matrix {

}
```

### Usage

Here is a program to test your function

```rust
// Importing Matrix by defining it as a dependency in Cargo.toml
use matrix_ops::*;

fn main() {
	let matrix = Matrix(vec![vec![8, 1], vec![9, 1]]);
	let matrix_2 = Matrix(vec![vec![1, 1], vec![1, 1]]);
	println!("{:?}", matrix + matrix_2);

	let matrix = Matrix(vec![vec![1, 3], vec![2, 5]]);
	let matrix_2 = Matrix(vec![vec![3, 1], vec![1, 1]]);
	println!("{:?}", matrix - matrix_2);

	let matrix = Matrix(vec![vec![1, 1], vec![1, 1]]);
	let matrix_2 = Matrix(vec![vec![1, 1, 3], vec![1, 1]]);
	println!("{:?}", matrix - matrix_2);

	let matrix = Matrix(vec![vec![1, 3], vec![9, 1]]);
	let matrix_2 = Matrix(vec![vec![1, 1, 3], vec![1, 1]]);
	println!("{:?}", matrix + matrix_2);
}
```

And its output

```console
$ cargo run
Some(Matrix([[9, 2], [10, 2]]))
Some(Matrix([[-2, 2], [1, 4]]))
None
None
$
```
