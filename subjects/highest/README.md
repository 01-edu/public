## Highest

### Instructions

In this exercise, a `Numbers` struct will be given. This struct will contain an array of `u32` numbers.

These methods have to be written for it:

- `new`: create a new instance of `Numbers`.
- `inner`: which returns the inner array as a slice.
- `last`: which returns an `Option<u32>` with the last added number.
- `highest`: which returns an `Option<u32>` with the biggest number from the list.
- `highest_three`: which returns an `Option<[u32; 3]>` with the three biggest numbers. These numbers should be ordered. If there aren't three numbers in the array, return `None`.

### Expected functions

```rust
#[derive(Debug)]
pub struct Numbers {
    numbers: [u32; _],
}

impl Numbers {
    pub fn new(numbers: [u32; _]) -> Self {
        todo!()
    }

    pub fn inner(&self) -> &[u32] {
        todo!()
    }

    pub fn latest(&self) -> Option<u32> {
        todo!()
    }

    pub fn highest(&self) -> Option<u32> {
        todo!()
    }

    pub fn highest_three(&self) -> Option<[u32; 3]> {
        todo!()
    }
}
```

### Usage

Here is a program to test your function.

```rust
use highest::*;

fn main() {
    let expected = [30, 500, 20, 70];
    let n = Numbers::new(expected);
    println!("{:?}", n.inner());
    println!("{:?}", n.highest());
    println!("{:?}", n.latest());
    println!("{:?}", n.highest_three());
}
```

And its output:

```console
$ cargo run
[30, 500, 20, 70]
Some(500)
Some(70)
Some([500, 70, 30])
$
```

### Notions

- [Trait iterator](https://doc.rust-lang.org/std/iter/trait.Iterator.html)
