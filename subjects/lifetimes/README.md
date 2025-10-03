## lifetimes

### Instructions

Complete the `Person` struct with the fields and associated function described below. `new` should set the `age` to `0`.

### Expected Functions and Data Structures (Both need to be completed)

```rust
#[derive(Debug, Clone, Copy, PartialEq)]
pub struct Person {
    pub name: &str,
    pub age: u32,
}

impl Person {
    pub fn new(name: &str) -> Self {
        todo!()
    }
}
```

### Usage

Here is a program to test your function.

```rust
use lifetimes::*;

fn main() {
    let person = Person::new("Leo");

    println!("Person = {:?}", person);
}
```

And its output:

```console
$ cargo run
Person = Person { name: "Leo", age: 0 }
$
```

### Notions

- [lifetimes](https://doc.rust-lang.org/book/ch10-03-lifetime-syntax.html)
