## collect

### Instructions

Implement the **function** `bubble_sort`, which receives a `Vec<i32>` and returns the same vector but in increasing order using the bubble sort algorithm.

### Expected Function

```rust
pub fn bubble_sort(vec: &mut Vec<i32>) {
}
```

### Usage

Here is a program to test your function.

```rust
use collect::*;

fn main() {
	let ref mut v = vec![3, 2, 4, 5, 1, 7];
	let mut b = v.clone();
	bubble_sort(v);
	println!("{:?}", v);

	b.sort();
	println!("{:?}", b);
}
```

And its output:

```console
$ cargo run
[1, 2, 3, 4, 5, 7]
[1, 2, 3, 4, 5, 7]
$
```
