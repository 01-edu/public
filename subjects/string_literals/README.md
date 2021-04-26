## string literals

### Instructions

Create the following functions:

- `is_empty`, that returns true if a string is empty
- `is_ascii`, that returns true if all characters of a given string is in ASCII range
- `contains`, that returns true if the string contains a pattern given
- `split_at`, that divides a string in two returning a tuple
- `find', that returns the index if the first character of a given string that matches the pattern

> This exercise will test how many times the **heap is going to be allocated**!\
> So try your best to allocate the minimum data on the heap!

### Notions

- https://doc.rust-lang.org/1.22.0/book/first-edition/the-stack-and-the-heap.html
- https://doc.rust-lang.org/rust-by-example/primitives/literals.html

### Expected Functions

```rust
pub fn is_empty(v: &str) -> bool {
}

pub fn is_ascii(v: &str) -> bool {
}

pub fn contains(v: &str, pat: &str) -> bool {
}

pub fn split_at(v: &str, index: usize) -> (&str, &str) {
}

pub fn find(v: &str, pat: char) -> usize {
}
```

### Usage

Here is a program to test your function

```rust
use string_literals::*;

fn main() {
    println!("{}", is_empty(""));
    println!("{}", is_ascii("rust"));
    println!("{}", contains("rust", "ru"));
    println!("{:?}", split_at("rust", 2));
    println!("{}", find("rust", 'u'));
}
```

And its output

```console
student@ubuntu:~/string_literals/test$ cargo run
true
true
true
("ru", "st")
1
student@ubuntu:~/string_literals/test$
```
