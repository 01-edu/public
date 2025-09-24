## moving_targets

### Instructions

You will have a linked list of `Target` named `Field`.

You will handle recursive types and ownership and implement the following **associated functions**:

- `new`: which will initialize the `Field` with `head` set to `None`.
- `push`: which receives a `Target` and add it as a `Node` at the head of the list.
- `pop`: which returns the last added `Target` wrapped in an `Option` and removes it from the list.
- `peek`: which returns the last added `Target` as a reference wrapped in an `Option` but does not remove it from the list.
- `peek_mut`: which returns the last added `Target` as a mutable reference wrapped in an `Option` and also does not remove it from the list.

You must also implement a type named `Link`. This will be the connection between the `Field` and `Target` structures. This will be a recursion type, and it must point to `None` if there is no `Target` to point to.

### Expected Functions and structures

```rust
#[derive(Debug, PartialEq, Eq)]
pub struct Target {
    pub size: u32,
    pub xp: u32,
}

pub struct Field {
    head: Link,
}

type Link = // To be implemented

struct Node {
    elem: Target,
    next: Link,
}

impl Field {
    pub fn new() -> Self {
        todo!()
    }

    pub fn push(&mut self, target: Target) {
        todo!()
    }

    pub fn pop(&mut self) -> Option<Target> {
        todo!()
    }

    pub fn peek(&self) -> Option<&Target> {
        todo!()
    }

    pub fn peek_mut(&mut self) -> Option<&mut Target> {
        todo!()
    }
}
```

### Usage

Here is a program to test your function:

```rust
use moving_targets::*;

fn main() {
    let mut field = Field::new();

    println!("{:?}", field.pop());
    field.push(Target { size: 12, xp: 2 });
    println!("{:?}", *field.peek().unwrap());
    field.push(Target { size: 24, xp: 4 });
    println!("{:?}", field.pop());
    let last_target = field.peek_mut().unwrap();
    *last_target = Target { size: 2, xp: 0 };
    println!("{:?}", field.pop());
}
```

And its output:

```console
$ cargo run
None
Target { size: 12, xp: 2 }
Some(Target { size: 24, xp: 4 })
Some(Target { size: 2, xp: 0 })
$
```

### Notions

- [Box\<T\>](https://doc.rust-lang.org/book/ch15-01-box.html)
- [Linked lists in Rust](https://rust-unofficial.github.io/too-many-lists/index.html)
