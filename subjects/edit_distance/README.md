## edit_distance

### Instructions

Create a function call `edit_distance` that calculates the minimum number of changes (insertion, deletions and substitutions) that need to be made to a string `source` to arrive to another `target` string

### Expected Function

```rust
pub fn edit_distance(source: &str, target: &str) -> usize {
}
```

### Notions

For more information and examples https://en.wikipedia.org/wiki/Edit_distance

### Usage

Here is a program to test your function.

```rust
fn main() {
	let source = "alignment";
	let target = "assignment";
	println!(
		"It's necessary to make {} change(s) to {}, to get {}",
		edit_distance(source, target),
		source,
		target
	);
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
It's necessary to make 2 change(s) to alignment, to get assignment
student@ubuntu:~/[[ROOT]]/test$
```
