## lalgebra_scalar

### Instructions

Define a `Scalar` trait which implements the operations `Add`, `Sub`, `Mul`, `Div` and any other restrictions you may need. Use a trait inheritance (supertraits)

Another condition for a number to be a scalar is to have a `zero` (as the neutral element in the addition), and a `one` (as the neutral element in the multiplication). Therefore the `Scalar` trait will require these two functions (described below).

After finishing completing the declaration of the trait, implement the `Scalar` trait for `u32`, `u64`, `i32`, `i64`, `f32` and `f64`.

### Expected Function (The signature must be completed)

> You need add the impl for each cases asked in the subject

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
use lalgebra_scalar::*;

fn main() {
	println!("{:?}", f64::zero());
	println!("{:?}", i32::zero());
	println!("{:?}", f64::one());
	println!("{:?}", i32::one());
}
```

And its output:

```console
$ cargo run
0.0
0
1.0
1
$
```

### Notions

- [Module std::ops](https://doc.rust-lang.org/std/ops/index.html)
[Module std::marker::Sized](https://doc.rust-lang.org/std/marker/trait.Sized.html)
