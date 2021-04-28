## spelling

### Instructions

In this exercise a number between 0 and 1000000 will be generated.
Your purpose is to create the function `spell` that will spell the numbers generated.

So, if the program generates the number:

- 1 your function will return the string "one"
- 14 your function will return the string "fourteen".
- 96 your function will return the string "ninety-six"
- 100 your function will return the string "one hundred".
- 101 your function will return the string "one hundred one"
- 348 your function will return the string "one hundred twenty-three"
- 1002 your function will return the string "one thousand two".
- 1000000 your function will return the string "one million"

Only positive numbers will be tested. (Up to a million).

### Notions

- [patterns](https://doc.rust-lang.org/book/ch18-00-patterns.html)

### Expected function

```rust
pub fn spell(n: u64) -> String {

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
