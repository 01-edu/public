## spelling

### Instructions

In this exercise, you'll create the function `spell` that will spell out a generated number.

Here are some examples of what your function should return:

- `1` -> `"one"`
- `14` -> `"fourteen"`
- `96` -> `"ninety-six"`
- `100` -> `"one hundred"`
- `101` -> `"one hundred one"`
- `348` -> `"three hundred forty-eight"`
- `1002` -> `"one thousand two"`
- `1000000` -> `"one million"`

> Only positive numbers will be tested, up to `1_000_000` (one million).

### Expected function

```rust
pub fn spell(n: u64) -> String {
    todo!()
}
```

### Usage

Here is a program to test your function.

```rust
use spelling::*;

fn main() {
    println!("{}", spell(348));
    println!("{}", spell(9996));
}
```

And its output:

```console
$ cargo run
three hundred forty-eight
nine thousand nine hundred ninety-six
$
```

### Notions

- [patterns](https://doc.rust-lang.org/book/ch18-00-patterns.html)
