// First exercise

// # Instructions
// Define a data structure to represent a matrix of any size and
// implement the basic operations for this you will need to follow the
// next steps:

// You can use a 2 dimensional Vec<T>'s
// We will consider a matrix as a rectangular arrangements of scalars
// You can use the definition of scalars done in the last exercise:
// `lalgebra_scalar`

// Then define the associated function `identity` that returns the identity matrix
// of size n
// Ex:
// Matrix::identity(3) == [[1,0,0], [0,1,0], [0,0,1]]

// And the associated function `zero` that returns a matrix of size
// `row x col` with all the positions filled by zeroes
// Ex:
// Matrix::zero(3, 3) == [[0,0,0],[0,0,0],[0,0,0]]

// Resources: https://doc.rust-lang.org/book/ch19-03-advanced-traits.html

use matrix::Matrix;

#[allow(dead_code)]
fn main() {
	let m: Matrix<u32> = Matrix(vec![vec![0, 0, 0, 0], vec![0, 0, 0, 0], vec![0, 0, 0, 0]]);
	println!("{:?}", m);
	println!("{:?}", Matrix::<i32>::identity(4));
	println!("{:?}", Matrix::<i32>::zero(3, 4));
}

#[test]
fn zero_property() {
	let matrix: Matrix<u32> = Matrix::zero(3, 4);
	let expected: Matrix<u32> = Matrix(vec![vec![0, 0, 0, 0], vec![0, 0, 0, 0], vec![0, 0, 0, 0]]);
	assert_eq!(matrix, expected);

	let matrix: Matrix<u32> = Matrix::zero(2, 2);
	let expected: Matrix<u32> = Matrix(vec![vec![0, 0], vec![0, 0]]);
	assert_eq!(matrix, expected);
}

#[test]
fn identy_matrix() {
	let matrix: Matrix<u32> = Matrix::identity(2);
	let expected: Matrix<u32> = Matrix(vec![vec![1, 0], vec![0, 1]]);
	assert_eq!(matrix, expected);

	let matrix: Matrix<u32> = Matrix::identity(3);
	let expected: Matrix<u32> = Matrix(vec![vec![1, 0, 0], vec![0, 1, 0], vec![0, 0, 1]]);
	assert_eq!(matrix, expected);
}
