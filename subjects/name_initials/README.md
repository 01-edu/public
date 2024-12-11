## name_initials

### Instructions

Create a **function** named `initials`. This function will receive a vector of string literals with names, and return a vector of Strings with the initials of each name.

### Expected Functions

```rust
pub fn initials(names: Vec<&str>) -> Vec<String> {
}
```

> Your heap allocations will be monitored to ensure that you do not make too many allocations, and that your allocations are reasonably sized.

### Usage

Here is a program to test your function:

```rust
use name_initials::*;

fn main() {
    let names = vec!["Harry Potter", "Someone Else", "J. L.", "Barack Obama"];
    println!("{:?}", initials(names));
}
```

And its output

```console
$ cargo run
["H. P.", "S. E.", "J. L.", "B. O."]
$
```

### Notions

- [stack and heap](https://doc.rust-lang.org/1.22.0/book/first-edition/the-stack-and-the-heap.html)
