## smallest

### Instructions

Create a function named `smallest` that gets the smallest number in the `HashMap`.

If the `HashMap` is empty, return the maximum `i32`.

### Expected Function

```rust
pub fn smallest(h: HashMap<&str, i32>) -> i32 {
}
```

### Usage

Here is a program to test your function.

```rust
use std::collections::HashMap;
use smallest::smallest;

fn main() {

    let mut hash = HashMap::new();
    hash.insert("Cat", 122);
    hash.insert("Dog", 333);
    hash.insert("Elephant", 334);
    hash.insert("Gorilla", 14);

    println!("The smallest of the elements in the HashMap is {}", smallest(hash));
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
