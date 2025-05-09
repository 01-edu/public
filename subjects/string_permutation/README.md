## string_permutation

### Instructions

Define the **function** `is_permutation`, that returns `true` if the string `s1` is a permutation of `s2`.

`s1` is a permutation of `s2` if all the elements in `s1` appear the same number of times in `s2`, and all the characters in `s1` appear in `s2` even if they are in different order.

### Expected Function

```rust
pub fn is_permutation(s1: &str, s2: &str) -> bool {
}
```

### Usage

Here is a program to test your function.

```rust
use string_permutation::*;

fn main() {
    let word = "thought";
    let word1 = "thougth";

    println!(
        "Is {:?} a permutation of {:?}? = {}",
        word,
        word1,
        is_permutation(word, word1)
    );
}
```

And its output

```console
$ cargo run
Is "thought" a permutation of "thougth"? = true
$
```

### Notions

- [hash maps](https://doc.rust-lang.org/book/ch08-03-hash-maps.html)
