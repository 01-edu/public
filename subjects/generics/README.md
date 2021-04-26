## generics

### Instructions

Write a **function** called `identity` which calculates the identity of a value (receives any data type and returns the same value).

### Expected Function (signature to be completed)

```rust
pub fn identity(v: _) -> _ {
}
```

### Usage

Here is a program to test your function.

```rust
use generics::*;

fn main() {
	println!("{}", identity("Hello, world!"));
	println!("{}", identity(3));
}
```

And its output:

```console
student@ubuntu:~/generics/test$ cargo run
Hello, world!
3
student@ubuntu:~/generics/test$
```
