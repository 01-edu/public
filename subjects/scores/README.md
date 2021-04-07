## score

### Instructions

Lets play a little!
Create a function `score` that given a string, computes the score for that given string.

Each letter has their value, you just have to sum the values of the letters in the
given string.

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

### Notions

- [patterns](https://doc.rust-lang.org/book/ch18-00-patterns.html)

### Expected functions

```rust
pub fn score(word: &str) -> u64 {

}
```

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
student@ubuntu:~/[[ROOT]]/test$ cargo run
1
0
14
student@ubuntu:~/[[ROOT]]/test$
```
