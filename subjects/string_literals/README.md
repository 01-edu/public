## string literal

### Instructions

Create the following functions:

- `is_empty`, that returns true if a string is empty
- `is_ascii`, that returns true if all characters of a given string is in ASCII range
- `contains`, that returns true if the string contains a pattern given
- `split_at`, that divides a string in two returning a tuple
- `find', that returns the index if the first character of a given string that matches the pattern

> This exercise will test the **heap allocation** of your function!
> So try your best to allocate the minimum data on the heap! (hit: &str)

### Notions

- https://doc.rust-lang.org/1.22.0/book/first-edition/the-stack-and-the-heap.html
- https://doc.rust-lang.org/rust-by-example/primitives/literals.html

### Expected Functions

```rust
fn is_empty(v: &str) -> bool {
}

fn is_ascii(v: &str) -> bool {
}

fn contains(v: &str, pat: &str) -> bool {
}

fn split_at(v: &str, index: usize) -> (&str, &str) {
}

fn find(v: &str, pat: char) -> usize {
}
```

### Usage

Here is a program to test your function

```rust

```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
student@ubuntu:~/[[ROOT]]/test$
```
