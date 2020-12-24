## groceries

### Instructions

Create a function called `insert` that inserts a new element at the end of the Vec

And another function `at_index` that returns the value found at the index passed as an argument

### Expected Functions

```rust
pub fn insert(vec: &mut Vec<String>, val: String) {
}

pub fn at_index(vec: &Vec<String>, index: ) -> String {
}
```

### Usage

```rust
use groceries::{insert, at_index}

fn main() {
	let mut groceries = vec![
		"yogurt".to_string(),
		"panetone".to_string(),
		"bread".to_string(),
		"cheese".to_string(),
	];
	insert(&mut groceries, String::from("nuts"));
	println!("The groceries list now = {:?}", &groceries);
	println!(
		"The second element of the grocery  list is {:?}",
		at_index(&groceries, 1)
	);
}
```

And its output:

```rust
The groceries list now = ["yogurt", "panetone", "bread", "cheese", "nuts"]
The second element of the grocery  list is "panetone"
```
