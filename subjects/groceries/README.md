## groceries

### Instructions

Create a **function** named `insert`, that inserts a new element at the end of the `Vec`.

Create another **function** named `at_index` that returns the value found at the index passed as an argument.


```rust
pub fn insert(vec: &mut Vec<String>, val: String) {
}

pub fn at_index(vec: &Vec<String>, index: usize) -> String {
}
```

### Usage

Here is a possible program to test your function:

```rust
use groceries::{insert, at_index};

fn main() {
	let mut groceries = vec![
		"yogurt".to_string(),
		"panettone".to_string(),
		"bread".to_string(),
		"cheese".to_string(),
	];
	insert(&mut groceries, String::from("nuts"));
	println!("The groceries list contains {:?}", &groceries);
	println!(
		"The second element of the grocery  list is {:?}",
		at_index(&groceries, 1)
	);
}
```

And its output:

```console
$ cargo run
The groceries list contains ["yogurt", "panettone", "bread", "cheese", "nuts"]
The second element of the grocery  list is "panettone"
$
```

### Notions

- [Common Collections](https://doc.rust-lang.org/stable/book/ch08-00-common-collections.html)
- [Vectors](https://doc.rust-lang.org/stable/book/ch08-01-vectors.html)