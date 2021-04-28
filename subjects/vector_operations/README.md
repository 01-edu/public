## vector_operations

### Instructions

Define the structure `ThreeDvector` that represents a 3 dimensional vector in (for convention in physics the vector are represented as ai + bj + ck where a, b, and c are real numbers and i, j and k represent the direction x,y and z respectively in the Cartesian plane there for we use the names i, j and k for the fields in the ThreeDVector structure

Look how the operations Addition and Subtraction work for a 3 dimensional vector and implement them by implementing the std::ops::Add and std::ops::Sub traits

### Expected Functions and Structures

```rust
#[derive(Debug, Copy, Clone, PartialEq)]
pub struct ThreeDVector<T> {
	pub i: T,
	pub j: T,
	pub k: T,
}

use std::ops::{Add, Sub};

impl Add for ThreeDVector<T> {
}

impl Sub for ThreeDVector<T> {
}
```

### Usage

Here is a program to test your function.

```rust
use vector_operations::ThreeDVector;

fn main() {
	let a = ThreeDVector { i: 3, j: 5, k: 2 };
	let b = ThreeDVector { i: 2, j: 7, k: 4 };
	println!("{:?}", a + b);
}
```

And its output

```console
$ cargo run
ThreeDVector { i: 5, j: 12, k: 6 }
$
```
