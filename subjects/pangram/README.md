## pangram

### Instructions

Create a `function` is_pangram which will determine whether or not a string is a pangram.

A pangram is a sentence which uses every letter of the alphabet at least once.

Example:

"The quick brown fox jumps over the lazy dog."

### Notions

- [patterns](https://doc.rust-lang.org/book/ch18-00-patterns.html)

### Expected functions

```rust
pub fn is_pangram(s: &str) -> bool {

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
student@ubuntu:~/pangram/test$ cargo run
true
false
student@ubuntu:~/pangram/test$
```
