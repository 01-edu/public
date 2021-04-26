## borrow_me_the_reference

### Instructions

Ownership is Rust's most unique feature, and it enables Rust to make memory safety guarantees without
needing garbage collector. Therefore you must understand ownership in rust.

Create the following functions :

  - `delete_and_backspace`, imagine that the `-` represents the `backspace key` and the `+` represents the `delete key`, this function must receive a borrowed string and turn this string into the string that applies the `backspaces` and the `deletes`.
  - For example:
	- "helll-o" turns into "hello"
  
  	- "he+lllo" turns into "hello"

  - `is_correct` that borrows a Vector of string literals with some correct and incorrect math equations and replaces the correct equations with `✔` and the wrong with `✘` and returns a `usize` with the percentage of correct equations.
	
### Expected Functions	

```rust
pub fn delete_and_backspace(s: &mut String) {
}

pub fn is_correct(v: &mut Vec<&str>) -> usize {
}
```

### Notions

- https://doc.rust-lang.org/book/ch04-00-understanding-ownership.html
- https://docs.rs/meval/0.2.0/meval/


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
student@ubuntu:~/borrow_me_the_reference/test$ cargo run
("borrow", ["✔", "✔", "✘", "✔"], 75)
student@ubuntu:~/borrow_me_the_reference/test$
```
