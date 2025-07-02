## pangram

### Instructions

Create a **function** named `is_pangram` which returns `true` if the given string is a pangram.

A pangram is a sentence which uses every letter of the alphabet at least once.

Example: "The quick brown fox jumps over the lazy dog."

### Expected functions

```rust
pub fn is_pangram(s: &str) -> bool {
    todo!()
}
```

### Usage

Here is a program to test your function.

```rust
use pangram::*;

fn main() {
    println!(
        "{}",
        is_pangram("the quick brown fox jumps over the lazy dog!")
    );
    println!("{}", is_pangram("this is not a pangram!"));
}
```

And its output:

```console
$ cargo run
true
false
$
```

### Notions

- [patterns](https://doc.rust-lang.org/book/ch18-00-patterns.html)
