## matrix_display

### Instructions

Complete the `Matrix` struct below.

You will need to create the `new` associated function which initializes the struct.

You will also need to implement the `std::fmt::Display` trait, so that it prints like the example in the usage.

### Expected Functions and Struct

```rust
use std::fmt;

#[derive(Debug, Clone)]
pub struct Matrix(pub Vec<Vec<i32>>);

impl Matrix {
    pub fn new(slice: &[&[i32]]) -> Self {
        todo!()
    }
}

impl fmt::Display for Matrix {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        todo!()
    }
}
```

### Usage

Here is a possible program to test your function:

```rust
use matrix_display::*;

fn main() {
    let matrix = Matrix::new(&[&[1, 2, 3], &[4, 5, 6], &[7, 8, 9]]);
    println!("{}", matrix);
}
```

And its output:

```console
$ cargo run
(1 2 3)
(4 5 6)
(7 8 9)
$
```
