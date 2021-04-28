## arrange_it

### Instructions

Create a **function** called `arrange_phrase` that takes a string literal as a phrase and returns it organized
Each word will have a number that indicates the position of that word.

> This exercise will test how many times the **heap is going to be allocated**!\
> So try your best to allocate the minimum data on the heap!

### Expected Function

```rust
pub fn arrange_phrase(phrase: &str) -> String {
}
```

### Notions

- [stack and heap](https://doc.rust-lang.org/1.22.0/book/first-edition/the-stack-and-the-heap.html)
- [str](https://doc.rust-lang.org/std/primitive.str.html)

### Usage

Here is a program to test your function

```rust
use arrange_it::*;

fn main() {
    println!("{:?}", arrange_phrase("is2 Thi1s T4est 3a"));
}
```

And its output

```console
$ cargo run
"This is a Test"
$
```
