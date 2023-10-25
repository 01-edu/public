## flat_tree

### Instructions

Create the `flatten_tree` **function** which receives a `std::collections::BTreeSet` and returns a new `Vec` with the elements of the binary tree in order.

### Expected function

```rust
pub fn flatten_tree<T: ToOwned<Owned = T>>(tree: &BTreeSet<T>) -> Vec<T> {

}
```

### Usage

Here is a possible program to test your function:

```rust
use flat_tree::*;
use std::collections::BTreeSet;

fn main() {
    let mut tree = BTreeSet::new();
    tree.insert(34);
    tree.insert(0);
    tree.insert(9);
    tree.insert(30);
    println!("{:?}", flatten_tree(&tree));

    let mut tree = BTreeSet::new();
    tree.insert("Slow");
    tree.insert("kill");
    tree.insert("will");
    tree.insert("Horses");
    println!("{:?}", flatten_tree(&tree));
}

```

And its output:

```console
$ cargo run
[0, 9, 30, 34]
["Horses", "Slow", "kill", "will"]
$
```
