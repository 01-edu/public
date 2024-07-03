## borrow_me_the_reference

### Instructions

> Ownership is Rust's most unique feature. It enables Rust to make memory safety guarantees without needing a garbage collector.

Understanding ownership is essential to take full advantage of Rust capabilities, it influences almost all aspects of the language.

Create the following functions:

- `delete_and_backspace`: which receives a borrowed string, and processes it. `-` represents the backspace key and `+` represents the delete key, so that `"helll-o"` and `"he+lllo"` are both converted to `"hello"`. The `-` and `+` characters should be removed from the string.

- `do_operations`: which borrows a Vector of string literals representing simple addition and subtraction equations. The function should replace the operation with the result.

### Expected Functions

```rust
pub fn delete_and_backspace(s: &mut String) {
}

pub fn do_operations(v: &mut Vec<String>) {
}
```

### Usage

Here is a program to test your function

```rust
use borrow_me_the_reference::{delete_and_backspace, do_operations};

fn main() {
	let mut a = String::from("bpp--o+er+++sskroi-++lcw");
    let mut b: Vec<String> = vec![
        "2+2".to_string(),
        "3+2".to_string(),
        "10-3".to_string(),
        "5+5".to_string(),
    ];

	delete_and_backspace(&mut a);
	do_operations(&mut b);

	println!("{:?}", (a, b));
}
```

And its output

```console
$ cargo run
("borrow", ["4", "5", "7", "10"])
$
```

### Notions

- [Ownership](https://doc.rust-lang.org/book/ch04-00-understanding-ownership.html)
