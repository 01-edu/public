## lalgebra_scalar

### Instructions

A scalar type must implement the operations Addition, Subtraction, Multiplication and Division (you might also have to use more restrictions). For this use a trait inheritance (supertraits)

Another condition for a number to be a scalar is to have a zero (neutral element in the addition) and a one (neutral element in the multiplication). Therefore the Scalar trait will require 2 functions zero() and one()

After finishing implement the Scalar trait for u32, u64, i32, i64, f32, f64

### Expected Function

```rust
pub trait Scalar: _ {
	type Item;
	fn zero() -> Self::Item;
	fn one() -> Self::Item;
}
```

### Usage

Here is a program to test your function.

```rust
fn main() {
	println!("{:?}", f64::zero());
	println!("{:?}", i32::zero());
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
0.0
0
student@ubuntu:~/[[ROOT]]/test$
```
