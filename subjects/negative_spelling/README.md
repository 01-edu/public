## negative_spelling

### Instructions

In this exercise, you'll create the function `negative_spell` that will spell a negative number.

Here are some examples of what your function should return:

- `-1` -> `"minus one"`
- `-14` -> `"minus fourteen"`
- `-348` -> `"minus three hundred forty-eight"`
- `-1002` -> `"minus one thousand two"`

> Only negative numbers will be accepted, up to `"minus one million"`.
> If a positive number is passed the function should return `"error: positive number"`.

### Expected function

```rust
pub fn negative_spell(n: i64) -> String {
    todo!()
}
```

### Usage

Here is a program to test your function.

```rust
use negative_spelling::*;

fn main() {
    println!("{}", negative_spell(-1234));
    println!("{}", negative_spell(100));
}
```

And its output:

```console
$ cargo run
minus one thousand two hundred thirty-four
error: positive number
$
```

### Notions

- [patterns](https://doc.rust-lang.org/book/ch18-00-patterns.html)
