## delete_prefix

### Instructions

Define the function `delete_prefix(prefix, s)` that returns the string slice `s` with the `prefix` removed wrapped in Some. If `prefix ` is not contained in `s` return None

### Expected Function

```rust
fn delete_prefix(prefix: &str, s: &str) -> Option<&str> {
}
```

### Usage

Here is a program to test your function.

```rust
fn main() {
	println!("{:?}", delete_prefix("ab", "abcdefghijklmnop"));
	println!("{:?}", delete_prefix("x", "abcdefghijklmnop"));
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
Some("cdefghijklmnop")
None
student@ubuntu:~/[[ROOT]]/test$
```
