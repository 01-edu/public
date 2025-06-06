## borrow_me_the_reference

### Instructions

> Ownership is arguably Rust's most unique feature. It enables Rust to make memory safety guarantees without needing a garbage collector.

Understanding ownership is essential to take full advantage of Rust's capabilities, as it influences almost all aspects of the language.

Create the following functions:

- `delete_and_backspace`: which receives a borrowed string, and processes it. `-` represents the backspace key and `+` represents the delete key, so that `"helll-o"` and `"he+lllo"` are both converted to `"hello"`. The `-` and `+` characters should be removed from the string.

- `do_operations`: which borrows a vector of string literals representing simple addition and subtraction equations. The function should replace the operation with the result.

### Expected Functions

```rust
pub fn delete_and_backspace(s: &mut String) {
}

pub fn do_operations(v: &mut [String]) {
}
```

### Usage

Here is a program to test your function

```rust
use borrow_me_the_reference::*;

fn main() {
    let mut a = "bpp--o+er+++sskroi-++lcw".to_owned();
    let mut b = [
        "2+2".to_owned(),
        "3+2".to_owned(),
        "10-3".to_owned(),
        "5+5".to_owned(),
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
