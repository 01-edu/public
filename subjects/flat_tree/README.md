## flat_tree

### Instructions

Create the `flatten_tree` **function** which receives a `std::collections::BTreeSet` and returns a new `Vec` with the elements of the binary tree in order.

### Expected function

```rust
pub fn flatten_tree<T: ToOwned<Owned = T>>(tree: &BTreeSet<T>) -> Vec<T> {
    todo!()
}
```

### Usage

Here is a possible program to test your function:

```rust
use flat_tree::*;
use std::collections::BTreeSet;

fn main() {
    let tree = BTreeSet::from([34, 0, 9, 30]);
    println!("{:?}", flatten_tree(&tree));

    let tree = BTreeSet::from(["Slow", "kill", "will", "Horses"]);
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
