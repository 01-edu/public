## score

### Instructions

Lets play a little.

Create a function named `score` that given a `&str`, computes the score for that given string as a `u64`.

Each letter has a value, you just have to sum the values of the letters in the given string.

You will need these:

| Letter                       | Value |
| ---------------------------- | :---: |
| A, E, I, O, U, L, N, R, S, T |   1   |
| D, G                         |   2   |
| B, C, M, P                   |   3   |
| F, H, V, W, Y                |   4   |
| K                            |   5   |
| J, X                         |   8   |
| Q, Z                         |  10   |

### Expected functions

> You'll need to work out the function signature for yourself.

### Usage

Here is a program to test your function.

```rust
use scores::*;

fn main() {
    println!("{}", score("a"));
    println!("{}", score("ã ê Á?"));
    println!("{}", score("ThiS is A Test"));
}
```

And its output

```console
$ cargo run
1
0
14
$
```

### Notions

- [patterns](https://doc.rust-lang.org/book/ch18-00-patterns.html)
