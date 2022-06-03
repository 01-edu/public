## how many references

### Instructions

Create the following **functions**:

- `add_ele`: which adds an element to the value in the `Node`.
- `how_many_references`: which returns how many times the value is referenced in the code.
- `rm_all_ref`: which accepts an `Rc<String>` and removes all elements from the vector that are equal to that value. This should only happen if the two `Rc`s point to the same allocation.

### Expected Functions and structures

```rust
pub use std::rc::Rc;

pub struct Node {
    pub value: Vec<Rc<String>>,
}

impl Node {
    pub fn new(value: Vec<Rc<String>>) -> Node {
        Node { value: value }
    }
    pub fn add_ele(&mut self, v: Rc<String>) {}
    pub fn rm_all_ref(&mut self, v: Rc<String>) {}
}

pub fn how_many_references(value: &Rc<String>) -> usize {}
```

### Usage

Here is a program to test your functions,

```rust
use how_many_references::*;

fn main() {
    let a = Rc::new(String::from("a"));
    let b = Rc::new(String::from("b"));
    let c = Rc::new(String::from("c"));

    let a1 = Rc::new(String::from("a"));

    let mut new_node = Node::new(vec![a.clone()]);
    new_node.add_ele(b.clone());
    new_node.add_ele(a.clone());
    new_node.add_ele(c.clone());
    new_node.add_ele(a.clone());

    println!("a: {:?}", how_many_references(&a));
    println!("b: {:?}", how_many_references(&b));
    println!("c: {:?}", how_many_references(&c));
    new_node.rm_all_ref(a1.clone());
    new_node.rm_all_ref(a.clone());

    println!("a: {:?}", how_many_references(&a));
    println!("b: {:?}", how_many_references(&b));
    println!("c: {:?}", how_many_references(&c));
}
```

And its output:

```console
$ cargo run
a: 4
b: 2
c: 2
a: 1
b: 2
c: 2
$
```

### Notions

- [Reference Counted Smart Pointer](https://doc.rust-lang.org/book/ch15-04-rc.html)
- [Struct std::rc::Rc](https://doc.rust-lang.org/std/rc/struct.Rc.html)
