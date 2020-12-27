## generics

### Instructions

Write a functions called identity that calculates the identity of a value (receives any data type and returns the same value)

### Expected Function

```rust
fn identity(v: _) -> _ {
}
```

### Usage

Here is a program to test your function.

```rust
fn main() {
	println!("{}", identity("Hello, world!"));
	println!("{}", identity(3));
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
Hello, world!
3
student@ubuntu:~/[[ROOT]]/test$
```
