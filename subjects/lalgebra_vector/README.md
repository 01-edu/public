## lalgebra_vector

### Instructions

A vector in linear algebra is defined as "anything that can be added and that can be multiplied by a scalar"

Define the associated function `dot` that calculates the dot product between two vectors

The dot product between two vectors of different length it's not defined

Note: `Vector` must implement Debug, Clone, Eq, PartialEq.

#### Important

Note that in the addition of vector of the example the `Rhs` is not `Self`

### Expected Functions and Structure

```rust
pub struct Vector<T: Scalar>(pub Vec<T>);

use std::ops::Add;

impl Add for Vector<T> {
}

impl Vector<T> {
	pub fn new() -> Self {
	}

	pub fn dot(&self, other: &Self) -> Option<T> {
}
```

### Usage

Here is a program to test your function.

```rust
use lalgebra_vector::*;

fn main() {
	let vector_1: Vector<i64> = Vector(vec![1, 3, -5]);
	let vector_2: Vector<i64> = Vector(vec![4, -2, -1]);
	println!("{:?}", vector_1.dot(&vector_2));
	println!("{:?}", vector_1 + &vector_2);
}
```

And its output:

```console
student@ubuntu:~/lalgebra_vector/test$ cargo run
Some(3)
Some(Vector([5, 1, -6]))
student@ubuntu:~/lalgebra_vector/test$
```
