## closures

### Instructions

Using closures and iterators create a **function**, that returns the first 50 even numbers squared in a `Vec<i32>`.

### Expected Functions

```rust
fn first_fifty_even_square() -> Vec<i32> {
    todo!()
}
```

### Usage

Here is a program to test your function.

```rust
use closures::*;

fn main() {
	println!("Hello, world!");
	let v1 = first_fifty_even_square();

	println!("All elements in {:?}, len = {}", v1, v1.len());
}
```

And its output:

```console
$ cargo run
All elements in [4, 16, 36, ..., 10000], len = 50
$
```

### Notions

[Iterators and Closures](https://doc.rust-lang.org/book/ch13-00-functional-features.html)
