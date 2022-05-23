## borrow_me_the_reference

### Instructions

Ownership is Rust's most unique feature. It enables Rust to make memory safety guarantees without needing a garbage collector. You must have a good understanding of ownership in rust.

Create the following functions:

- `delete_and_backspace`: which receives a borrowed string, and processes it. `-` represents the backspace key and `+` represents the delete key, so that `"helll-o"` and `"he+lllo"` are both converted to `"hello"`. The `-` and `+` characters should be removed from the string.

- `is_correct`: which borrows a Vector of string literals representing simple addition and subtraction equations. The function should replace the correct equations with `✔`, and the wrong ones with `✘`. It should return a `usize` with the percentage of correct equations.

### Expected Functions
```rust
pub fn delete_and_backspace(s: &mut String) {
}

pub fn is_correct(v: &mut Vec<&str>) -> usize {
}
```



### Dependencies

meval = "0.2"

### Usage

Here is a program to test your function

```rust
use borrow_me_the_reference::{delete_and_backspace, is_correct};

fn main() {
	let mut a = String::from("bpp--o+er+++sskroi-++lcw");
	let mut b: Vec<&str> = vec!["2+2=4", "3+2=5", "10-3=3", "5+5=10"];

	// - If a value does **not implement Copy**, it must be **borrowed** and so will be passed by **reference**.
	delete_and_backspace(&mut a); // the reference of  the value
	let per = is_correct(&mut b); // the reference of  the value

	println!("{:?}", (a, b, per));
	// output: ("borrow", ["✔", "✔", "✘", "✔"], 75)
}
```

And its output

```console
$ cargo run
("borrow", ["✔", "✔", "✘", "✔"], 75)
$
```

### Notions

- [ownership](https://doc.rust-lang.org/book/ch04-00-understanding-ownership.html)
- [meval](https://docs.rs/meval/0.2.0/meval/)
