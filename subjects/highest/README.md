## Highest

### Instructions

In this exercise a `Numbers` struct will be given.

These methods have to be written:

- `new` create a new instance of Number.
- `List` that returns an `array` with every number in the struct.
- `Latest` that returns an `Option<u32>` with the last added number.
- `Highest` that return an `Option<u32>` with the highest number from the list.
- `Highest_Three` that returns a `Vec<u32>` with the three highest numbers.

### Notions

- [Trait iterator](https://doc.rust-lang.org/std/iter/trait.Iterator.html)

### Expected functions

```rust
pub fn new(&[u32]) -> Self {}

pub fn List(&self) -> &[u32] {}

pub fn Latest(&self) -> Option<u32> {}

pub fn Highest(&self) -> Option<u32> {}

pub fn Highest_Three(&self) -> Vec<u32> {}
```

### Usage

Here is a program to test your function.

```rust
use highest::*;

#[derive(Debug)]
struct Numbers<'a> {
    numbers: &'a [u32],
}

fn main() {
    let expected = [30, 500, 20, 70];
    let n = Numbers::new(&expected);
    println!("{:?}", n.List());
    println!("{:?}", n.Highest());
    println!("{:?}", n.Latest());
    println!("{:?}", n.Highest_Three());
}
```

And its output:

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
[30, 500, 20, 70]
Some(500)
Some(70)
[500, 70, 30]
student@ubuntu:~/[[ROOT]]/test$
```
