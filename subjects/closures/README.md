## closures

### Instructions

Using closures and iterators create a function, `first_fifty_even_square` that returns the first 50

### Expected Functions

```rust
fn first_fifty_even_square() -> Vec<i32> {
}
```

### Usage

Here is a program to test your function.

```rust
fn main() {
	println!("Hello, world!");
	let v1 = first_fifty_even_square();

	println!("All elements in {:?}, len = {}", v1, v1.len());
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
All elements in [4, 16, 36, ..., 10000], len = 50
student@ubuntu:~/[[ROOT]]/test$
```
