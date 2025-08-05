## lalgebra_vector

### Instructions

A vector in linear algebra is defined as "anything that can be added, and that can be multiplied by a scalar".

Define the associated function `dot`, that calculates the dot product between two vectors. If the vectors are of different lengths, return `None`.

Note: `Vector` must implement at least `Debug` and `PartialEq`.

### Expected Functions and Structure

```rust
use std::ops::Add;

#[derive(Debug, PartialEq)]
pub struct Vector<T: Scalar>(pub Vec<T>);

impl Add<_> for Vector<_> {
}

impl Vector<_> {
	pub fn dot(self, rhs: Self) -> Option<T> {
		todo!()
	}
}
```

### Usage

Here is a program to test your function.

```rust
use lalgebra_vector::*;

fn main() {
	println!("{:?}", Vector(vec![1, 3, -5]).dot(Vector(vec![4, -2, -1])));
	println!("{:?}", Vector(vec![1, 3, -5]) + Vector(vec![4, -2, -1]));
}
```

And its output:

```console
$ cargo run
Some(3)
Some(Vector([5, 1, -6]))
$
```
