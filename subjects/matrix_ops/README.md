## matrix_ops

### Instructions

In this exercise you will define the basic operations with a matrix starting by implementing the `std::ops::Add` trait

Define the operation + (by defining the trait std::ops::Add) for two matrices. Remember that two matrices can only be added if they have the same size. Therefore the add method must handle the possibility of failure by returning an Option<T>.

You will be using your own `Matrix` and `Scalar` defined in the `matrix` and the `lalgebra_scalar` exercises.

### Expected Function

```rust
use crate::{Matrix, Scalar};
use std::ops::{ Add, Sub };

impl Add for Matrix {

}

impl Sub for Matrix {

}
```

### Usage

Here is a program to test your function

```rust
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
