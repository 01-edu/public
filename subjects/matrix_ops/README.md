## matrix_ops

### Instructions

In this exercise, you will implement some matrix operations: `Add`, `Sub` and `Mul`.

You will be reusing your `Matrix` and `Scalar` structures defined in the [matrix](../matrix/README.md) and [lalgebra-scalar](../lalgebra-scalar/README.md) exercises.

Additionally, as it isn't possible to implement foreign traits on foreign types, we will write a `Wrapper` struct that will simply wrap our `Matrix` struct and implement these traits on `Wrapper` instead.

> Remember that two matrices can only be added or subtracted if they have the same dimensions. Since we're dealing with arrays with sizes known at compile-time and compile-time generics, we don't need to check for dimensions at runtime. This is part of the magic of Rust's compile time generics!

As with the previous exercise, make sure the generics have the correct constraints.

### Expected Functions and Structure

```rust
#[derive(Debug, Eq, PartialEq, Clone, Copy)]
pub struct Wrapper<W, H, T>(pub Matrix<W, H, T>);

impl<W, H, T> From<[[T; W]; H]> for Wrapper<W, H, T> {

}

impl<W, H, T> Add for Wrapper<W, H, T> {

}

impl<W, H, T> Sub for Wrapper<W, H, T> {

}

impl<S, T> Mul for Wrapper<S, S, T> {

}
```

### Usage

Here is a program to test your function

```rust
use matrix_ops::*;

fn main() {
    let matrix = Wrapper::from([[8, 1], [9, 1]]);
    let matrix_2 = Wrapper::from([[1, 1], [1, 1]]);
    println!("{:?}", matrix + matrix_2);

    let matrix = Wrapper::from([[1, 3], [2, 5]]);
    let matrix_2 = Wrapper::from([[3, 1], [1, 1]]);
    println!("{:?}", matrix - matrix_2);

    let matrix = Wrapper::from([[1, 2], [3, 4]]);
    let matrix_2 = Wrapper::from([[2, 0], [1, 2]]);
    println!("{:?}", matrix * matrix_2);

    // The examples below should give a compile-time error.
    // Because we have correct const generics and arrays with a fixed, known size
    // we can't operate either with matrices of different sizes or with invalid matrices (for instance with rows of different sizes).

    // let matrix = Wrapper::from([[1, 1], [1, 1]]);
    // let matrix_2 = Wrapper::from([[1, 1, 3], [1, 1]]);
    // println!("{:?}", matrix - matrix_2);

    // let matrix = Wrapper::from([[1, 3], [9, 1]]);
    // let matrix_2 = Wrapper::from([[1, 1, 3], [1, 1, 4]]);
    // println!("{:?}", matrix + matrix_2);
}
```

And its output

```console
$ cargo run
Wrapper(Matrix([[9, 2], [10, 2]]))
Wrapper(Matrix([[-2, 2], [1, 4]]))
Wrapper(Matrix([[4, 4], [10, 8]]))
$
```
