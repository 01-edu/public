## matrix

### Instructions

Define a data structure to represent a matrix of any size **known at compile-time**.

We will consider a matrix as a rectangular arrangement of scalars. It can be represented as a 2 dimensional array. You will use the definition of `Scalar` from the [lalgebra-scalar](../lalgebra-scalar/README.md) exercise.

Implement the following associated functions:

- `zero`: which returns a matrix of size `W` x `H` with all the positions filled by `T::zero()`.
- `identity`: which returns the identity matrix of size `S`.

> Remember that arrays must have a size known at compile-time. The expected structure below gives some hints but you _will_ have to make some modifications to the types for the exercise to be correct.

### Expected Functions and Structure

```rust
#[derive(Debug, Eq, PartialEq, Clone, Copy)]
pub struct Matrix<W, H, T>(pub [[T; W]; H]);

impl<W, H, T> Matrix<W, H, T> {
    pub fn zero() -> Self {
        todo!()
    }
}

impl<S, T> Matrix<S, S, T> {
    pub fn identity() -> Self {
        todo!()
    }
}
```

### Usage

Here is a program to test your function.

```rust
use matrix::*;

fn main() {
    let m = Matrix([[0; 4]; 3]);
    println!("{:?}", m);
    println!("{:?}", Matrix::<4, 4, u32>::identity());
    println!("{:?}", Matrix::<3, 4, f64>::zero());
}
```

And its output:

```console
$ cargo run
Matrix([[0, 0, 0, 0], [0, 0, 0, 0], [0, 0, 0, 0]])
Matrix([[1, 0, 0, 0], [0, 1, 0, 0], [0, 0, 1, 0], [0, 0, 0, 1]])
Matrix([[0.0, 0.0, 0.0], [0.0, 0.0, 0.0], [0.0, 0.0, 0.0], [0.0, 0.0, 0.0]])
$
```

### Notions

- [Traits](https://doc.rust-lang.org/book/ch19-03-advanced-traits.html)
- [Constant generics](https://rust-lang.github.io/rfcs/2000-const-generics.html)
