## delete_prefix

### Instructions

Define the function `delete_prefix(prefix, s)` which returns the string slice `s` with the `prefix` removed wrapped in `Some`. If `prefix ` is not a prefix of `s` it returns `None`.

### Expected Function

```rust
pub fn delete_prefix(prefix: &str, s: &str) -> Option<&str> {
}
```

### Usage

Here is a program to test your function.

```rust
use delete_prefix::delete_prefix;

fn main() {
	println!("{:?}", delete_prefix("ab", "abcdefghijklmnop"));
	println!("{:?}", delete_prefix("x", "abcdefghijklmnop"));
}
```

And its output:

```console
student@ubuntu:~/delete_prefix/test$ cargo run
Some("cdefghijklmnop")
None
student@ubuntu:~/delete_prefix/test$
```
