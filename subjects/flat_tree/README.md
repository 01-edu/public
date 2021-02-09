## flat_rust

### Instructions

- Define the functions `flatten_tree` that receives a std::collections::BTreeSet and returns a new `Vec` with the elements in the binary tree in order.

### Expected function

```rust
pub fn flatten_tree<T: ToOwned<Owned = T>>(tree: &BTreeSet<T>) -> Vec<T> {
}
```

### Usage

Here is a possible test for your function:

```rust
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
student@ubuntu:~/[[ROOT]]/test$ cargo run
[0, 9, 30, 34]
["Horses", "Slow", "kill", "will"]
student@ubuntu:~/[[ROOT]]/test$
```
