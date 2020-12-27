// # Instructions
// A scalar type must implement the operations
// Addition, Subtraction, Multiplication and Division (you might
// also have to use more restrictions). For this use a trait
// inheritance (supertraits)

// Another condition for a number to be a scalar is to have a zero
// (neutral element in the addition) and a one (neutral element in the
// multiplication). Therefore the Scalar trait will require 2
// functions zero() and one()

// After finishing implement the Scalar trait for u32, u64, i32, i64,
// f32, f64

use lalgebra_scalar::Scalar;

#[allow(dead_code)]
fn main() {
	println!("{:?}", f64::zero());
	println!("{:?}", i32::zero());
}

#[test]
fn scalar_u32() {
	let a: u32 = u32::zero();
	assert_eq!(a, 0 as u32);

	let b = u32::one();
	assert_eq!(b, 1 as u32);
}

#[test]
fn scalar_u64() {
	let a = u64::zero();
	assert_eq!(a, 0 as u64);

	let b = u64::one();
	assert_eq!(b, 1 as u64);
}

#[test]
fn scalar_i32() {
	let a: i32 = i32::zero();
	assert_eq!(a, 0 as i32);

	let b = i32::one();
	assert_eq!(b, 1 as i32);
}

#[test]
fn scalar_i64() {
	let a: i64 = i64::zero();
	assert_eq!(a, 0 as i64);

	let b = i64::one();
	assert_eq!(b, 1 as i64);
}

#[test]
fn scalar_f32() {
	let zero = f32::zero();
	assert_eq!(zero, 0.0);
	let one = f32::one();
	assert_eq!(one, 1.0);
}

#[test]
fn scalar_f64() {
	let zero = f64::zero();
	assert_eq!(zero, 0.0);
	let one = f64::one();
	assert_eq!(one, 1.0);
}
