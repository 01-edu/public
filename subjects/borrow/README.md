## borrow

### Instructions

Complete the signature and the body of the `str_len` **function** that receives a string or a string literal and returns its length without taking ownership of the value (i.e, borrowing the value).

### Expected Function (The signature needs to be completed)

```rust
pub fn str_len(s: ) -> usize {
}
```

### Notions

- [Primitive Type str](https://doc.rust-lang.org/std/primitive.str.html)

### Usage

Here is a possible program to test your function :

```rust
use borrow::*;

fn main() {
	let s = "hello";
	let s1 = "camelCase".to_string();

	println!("\tstr_len(\"{}\") = {}", s, str_len(s));
	println!("\tstr_len(\"{}\") = {}", s1, str_len(&s1));
}
```

And its output:

```console
$ cargo run
str_len("hello") = 5
str_len("camelCase") = 9
$
```
