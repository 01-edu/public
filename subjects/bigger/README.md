## bigger

### Instructions

Create the function `bigger` that gets the biggest positive number in the `HashMap`.

### Notions

[hash maps](https://doc.rust-lang.org/book/ch08-03-hash-maps.html)

### Dependencies

rand = "0.7"

### Expected Function

```rust
pub fn bigger(h: HashMap<&str, i32>) -> i32 {
}
```

### Usage

Here is a program to test your function.

```rust
use std::collections::HashMap;
use bigger::bigger;

fn main() {

    let mut hash = HashMap::new();
    hash.insert("Daniel", 122);
    hash.insert("Ashley", 333);
    hash.insert("Katie", 334);
    hash.insert("Robert", 14);

    println!("The biggest of the elements in the HashMap is {}", bigger(hash));
}
```

And its output

```console
$ cargo run
The biggest of the elements in the HashMap is 334
$
```
