## bigger

### Instructions

Create a function named `bigger` that gets the biggest positive number in the `HashMap`.

### Expected Function

```rust
pub fn bigger(h: HashMap<&str, i32>) -> i32 {
}
```

### Usage

Here is a program to test your function.

```rust
use bigger::*;
use std::collections::HashMap;

fn main() {
    let hash = HashMap::from_iter([
        ("Daniel", 122),
        ("Ashley", 333),
        ("Katie", 334),
        ("Robert", 14),
    ]);

    println!(
        "The biggest of the elements in the HashMap is {}",
        bigger(hash)
    );
}
```

And its output

```console
$ cargo run
The biggest of the elements in the HashMap is 334
$
```

### Notions

- [hash maps](https://doc.rust-lang.org/book/ch08-03-hash-maps.html)
