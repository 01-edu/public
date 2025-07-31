## generics

### Instructions

Create a **function** named `identity` which receives an argument of any type and returns it.

### Expected Function (signature to be completed)

```rust
pub fn identity(v: _) -> _ {
	todo!()
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
$ cargo run
Hello, world!
3
$
```

### Notions

- [Generics](https://doc.rust-lang.org/book/ch10-01-syntax.html)
