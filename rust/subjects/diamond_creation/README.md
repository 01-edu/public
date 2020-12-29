## diamond_creation

### Instructions

Complete the function "make_diamond" that takes a letter as input, and outputs it in a diamond shape.

Rules:

- The first and last row contain one 'A'.
- The given letter has to be at the widest point.
- All rows, except the first and last, have exactly two identical letters.
- All rows have as many trailing spaces as leading spaces. (This might be 0).
- The diamond is vertically and horizontally symmetric.
- The diamond width equals the height.
- The top half has the letters in ascending order. (abcd)
- The bottom half has the letters in descending order. (dcba)

### Notions

- https://doc.rust-lang.org/book/ch18-03-pattern-syntax.html

### Expected functions

```rust
fn get_diamond(c: char) -> Vec<String> {}
```

### Usage

Here is a program to test your function.

```rust
fn main() {
    println!("{:?}", make_diamond('A'));
    println!("{:?}", make_diamond('C'));
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
["A"]
["  A  ", " B B ", "C   C", " B B ", "  A  "]
student@ubuntu:~/[[ROOT]]/test$
```
