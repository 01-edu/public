## delete_prefix

### Instructions

Define the function `delete_prefix` which returns the string slice `s` with the `prefix` removed. It should be wrapped in `Some`. If `prefix` is not a prefix of `s`, then `delete_prefix` returns `None`.

### Expected Function

```rust
pub fn delete_prefix(prefix: &str, s: &str) -> Option<&str> {
    todo!()
}
```

### Usage

Here is a program to test your function.

```rust
use delete_prefix::*;

fn main() {
    println!("{:?}", delete_prefix("ab", "abcdefghijklmnop"));
    println!("{:?}", delete_prefix("x", "abcdefghijklmnop"));
}
```

And its output:

```console
$ cargo run
Some("cdefghijklmnop")
None
$
```
