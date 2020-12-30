## highest

### Instructions

In this exercise you will be given a `Numbers` struct.

Your task is to write these methods:

- `list` that returns an `array` with every number in the struct
- `latest` that returns an `Option<u32>` with the last added number
- `highest` that return an `Option<u32>` with the highest number from the list,
- `highest_three` that returns a `Vec<u32>` with the three highest numbers.

### Notions

- https://doc.rust-lang.org/std/iter/trait.Iterator.html

### Expected functions

```rust
pub fn list(&self) -> &[u32] {}

pub fn latest(&self) -> Option<u32> {}

pub fn highest(&self) -> Option<u32> {}

pub fn highest_three(&self) -> Vec<u32> {}
```

### Usage

Here is a program to test your function.

```rust
use highest::highest;

#[derive(Debug)]
struct Numbers<'a> {
    numbers: &'a [u32],
}

fn main() {
    let expected = [30, 500, 20, 70];
    let n = Numbers::new(&expected);
    println!("{:?}", n.list());
    println!("{:?}", n.highest());
    println!("{:?}", n.latest());
    println!("{:?}", n.highest_three());
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
[30, 500, 20, 70]
Some(500)
Some(70)
[500, 70, 30]
student@ubuntu:~/[[ROOT]]/test$
```
