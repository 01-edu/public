## display_table

### Instructions

- Implement the `std::fmt::Display` trait for the structure table so that the table is printed like in the **[Usage](#usage)**. The length of each column must adjust to the longest element of the column and the element must be centered in the "cell" when possible. If the length of the element does not allow to center exactly, it must be offset slightly to the left.

  - Note: If the table is empty (does not have a header), `println!` must not print anything.

- Define the associated function `new` which creates a new empty table.

- Define the method function `add_row` which adds a new row to the table created from a slice of strings.

### Expected functions and Structures

```rust
#[derive(Clone, Debug, PartialEq)]
pub struct Table {
	pub headers: Vec<String>,
	pub body: Vec<Vec<String>>,
}

impl fmt::Display for Table {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {

    }
}

impl Table {
	pub fn new() -> Table {

	}
	pub fn add_row(&mut self, row: &[String]) {

	}
}
```

### Usage

Here is a possible program to test your function:

```rust
fn main() {
	let mut table = Table::new();
	println!("{}", table);
	table.headers = vec![
		String::from("Model"),
		String::from("Piece N°"),
		String::from("In Stock"),
		String::from("Description"),
	];
	table.add_row(&[
		String::from("model 1"),
		String::from("43-EWQE304"),
		String::from("30"),
		String::from("Piece for x"),
	]);
	table.add_row(&[
		String::from("model 2"),
		String::from("98-QCVX5433"),
		String::from("100000000"),
		String::from("-"),
	]);
	table.add_row(&[
		String::from("model y"),
		String::from("78-NMNH"),
		String::from("60"),
		String::from("nothing"),
	]);
	println!("{}", table);
}
```

And its output:

```console
$ cargo run
|  Model  |  Piece N°   | In Stock  | Description |
|---------+-------------+-----------+-------------|
| model 1 | 43-EWQE304  |    30     | Piece for x |
| model 2 | 98-QCVX5433 | 100000000 |      -      |
| model y |   78-NMNH   |    60     |   nothing   |
$
```
