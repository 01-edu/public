## box_it

### Instructions

Create the following **functions**:

- `parse_into_boxed`: which accepts a string of numbers separated by spaces. If a number has a `k` as a suffix it should be multiplied by 1000. The function parses these numbers and boxes them into a vector of `Box<u32>`.

- `into_unboxed`: which accepts the value returned from `parse_into_boxed` and unboxes each element into another vector.

### Expected Functions

```rust
pub fn parse_into_boxed(s: String) -> Vec<Box<u32>> {
    todo!()
}

pub fn into_unboxed(a: Vec<Box<u32>>) -> Vec<u32> {
    todo!()
}
```

### Usage

Here is a program to test your functions

```rust
use std::mem;

use box_it::*;

fn main() {
    let s = "5.5k 8.9k 32".to_owned();

    let boxed = parse_into_boxed(s);
    println!("Element value: {:?}", boxed[0]);
    println!("Element size: {:?} bytes", mem::size_of_val(&boxed[0]));

    let unboxed = into_unboxed(boxed);
    println!("Element value: {:?}", unboxed[0]);
    println!("Element size: {:?} bytes", mem::size_of_val(&unboxed[0]));

    // As with everything related to regular Rust memory management, both the `Vec` and the `Box`es will be properly dropped when out of scope and freed, ensuring no leaks
}
```

And its output:

```console
$ cargo run
Element value: 5500
Element size: 8 bytes
Element value: 5500
Element size: 4 bytes
$
```

### Notions

- [smart pointers](https://doc.rust-lang.org/book/ch15-00-smart-pointers.html)
- [using box](https://doc.rust-lang.org/book/ch15-01-box.html)
