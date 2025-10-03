## smallest

### Instructions

Create a function named `smallest` that gets the smallest number in the `HashMap`.

If the `HashMap` is empty, return the maximum `i32`.

### Expected Function

```rust
pub fn smallest(h: HashMap<&str, i32>) -> i32 {
    todo!()
}
```

### Usage

Here is a program to test your function.

```rust
use smallest::smallest;
use std::collections::HashMap;

fn main() {
    let hash = HashMap::from([
        ("Cat", 122),
        ("Dog", 333),
        ("Elephant", 334),
        ("Gorilla", 14),
    ]);

    println!(
        "The smallest of the elements in the HashMap is {}",
        smallest(hash)
    );
}

```

And its output

```console
$ cargo run
The smallest of the elements in the HashMap is 14
$
```

### Notions

- [hash maps](https://doc.rust-lang.org/book/ch08-03-hash-maps.html)
