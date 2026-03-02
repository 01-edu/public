## how many references

### Instructions

Create a structure `Node` which will contain a vector of `Rc`s with the following **methods**:

- `new`: which creates a new `Node` with the given initial state.
- `add_element`: which adds an element to the `ref_list`.
- `rm_all_ref`: which accepts an `Rc<String>` and removes all elements from the vector that are equal to that value. This should only happen if the two `Rc`s point to the same allocation.

Additionally, create the **function** `how_many_references`, which returns how many clones of the same `Rc` passed as argument exist.

### Expected Functions and structures

```rust
use std::rc::Rc;

pub struct Node {
    pub ref_list: Vec<Rc<String>>,
}

impl Node {
    pub fn new(ref_list: Vec<Rc<String>>) -> Self {
        todo!()
    }

    pub fn add_element(&mut self, element: Rc<String>) {
        todo!()
    }

    pub fn rm_all_ref(&mut self, element: Rc<String>) {
        todo!()
    }
}

pub fn how_many_references(ref_list: &Rc<String>) -> usize {
    todo!()
}
```

### Usage

Here is a program to test your functions,

```rust
use how_many_references::*;
use std::rc::Rc;

fn main() {
    let a = Rc::new("a".to_owned());
    let b = Rc::new("b".to_owned());
    let c = Rc::new("c".to_owned());

    let a1 = Rc::new("a".to_owned());

    let mut new_node = Node::new(vec![a.clone()]);
    new_node.add_element(b.clone());
    new_node.add_element(a.clone());
    new_node.add_element(c.clone());
    new_node.add_element(a.clone());

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
