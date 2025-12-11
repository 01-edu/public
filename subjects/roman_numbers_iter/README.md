## roman_numbers_iter

### Instructions

Implement the `Iterator` trait for the `RomanNumber` type. You should use the code from the previous exercise `roman_numbers`.

### Notions

- [Trait Iterator](https://doc.rust-lang.org/std/iter/trait.Iterator.html)

### Expected Functions

```rust
//...

impl Iterator for RomanNumber {}
```

### Usage

Here is a program to test your function.

```rust
use roman_numbers_iter::RomanNumber;

fn main() {
    let mut number = RomanNumber::from(15);

    println!("{:?}", number);
    println!("{:?}", number.next());
}
```

And its output

```console
$ cargo run
RomanNumber([X, V])
Some(RomanNumber([X, V, I]))
$
```
