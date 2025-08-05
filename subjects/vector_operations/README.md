## vector_operations

### Instructions

Define the structure `ThreeDVector`, that represents a 3 dimensional vector.

In physics, these vectors are represented as `ai`, `bj` and `ck`. `a`, `b` and `c` are real numbers. `i`, `j` and `k` represent the direction in the Cartesian plane, along the axises `x`, `y` and `z` respectively. Therefore, we use the fields `i`, `j` and `k` in the structure.

Take a look how the operations `Addition` and `Subtraction` work for a 3 dimensional vector, and implement them by using the `std::ops::Add` and `std::ops::Sub` traits.

### Expected Functions and Structures

```rust
use std::ops::{Add, Sub};

#[derive(Debug, Copy, Clone, PartialEq)]
pub struct ThreeDVector<T> {
	pub i: T,
	pub j: T,
	pub k: T,
}

impl Add for ThreeDVector {
}

impl Sub for ThreeDVector {
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
