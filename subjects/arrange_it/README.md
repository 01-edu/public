## arrange_it

### Instructions

Create a **function** named `arrange_phrase`, that takes a string literal, _sorts_ the words and returns it. Each word will contain a number that indicates the position of that word.

### Expected Functions

```rust
pub fn arrange_phrase(phrase: &str) -> String {
}
```

> Your heap allocations will be monitored to ensure that you do not make too many allocations, and that your allocations are reasonably sized.

### Usage

Here is a program to test your function

```rust
use arrange_it::*;

fn main() {
    println!("{}", arrange_phrase("is2 Thi1s T4est 3a"));
}
```

And its output

```console
$ cargo run
This is a Test
$
```

### Notions

- [stack and heap](https://doc.rust-lang.org/1.22.0/book/first-edition/the-stack-and-the-heap.html)
- [str](https://doc.rust-lang.org/std/primitive.str.html)
