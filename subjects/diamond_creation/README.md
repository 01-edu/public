## diamond_creation

### Instructions

Build the **function** `make_diamond` which takes a letter as an input, and returns a diamond.

> Assume the input is always a valid uppercase letter.

Rules:

- The first and last row contain one 'A'.
- The given letter has to be at the widest point.
- All rows, except the first and last, have exactly two identical letters.
- All rows have as many trailing spaces as leading spaces. This may be 0.
- The diamond is both vertically and horizontally symmetrical.
- The width and height of the diamond are equal.
- The top half has letters in ascending order (abcd).
- The bottom half has letters in descending order (dcba).

### Expected functions

```rust
pub fn get_diamond(c: char) -> Vec<String> {
    todo!()
}
```

### Usage

Here is a program to test your function.

```rust
use diamond_creation::*;

fn main() {
    println!("{:?}", get_diamond('A'));
    println!("{:?}", get_diamond('C'));
    for line in get_diamond('C') {
        println!("{}", line);
    }
}
```

And its output:

```console
$ cargo run
["A"]
["  A  ", " B B ", "C   C", " B B ", "  A  "]
  A  
 B B 
C   C
 B B 
  A  
$
```

### Notions

- [pattern syntax](https://doc.rust-lang.org/book/ch18-03-pattern-syntax.html)
