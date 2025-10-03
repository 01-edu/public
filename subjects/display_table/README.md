## display_table

### Instructions

Define the `Table` struct below, and implement the associated functions `new` and `add_row`. You can see how they should work from the usage.

Implement the `std::fmt::Display` trait for the `Table` structure so that the table is printed like in the usage. The length of each column must adjust to the longest element of the column, and the element must be centered in the "cell" when possible. If the element cannot be exactly centered, it must be offset one byte to the left of the center, as can be seen in the example output.

If the table is empty `println!` must not print anything.

> Our tests will always assume that, when calling `add_row`, the row has the same correct size as the headers.

### Expected functions and Structures

```rust
#[derive(Clone, Debug, PartialEq)]
pub struct Table {
    pub headers: Vec<String>,
    pub body: Vec<Vec<String>>,
}

impl fmt::Display for Table {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        todo!()
    }
}

impl Table {
    pub fn new() -> Table {
        todo!()
    }

    pub fn add_row(&mut self, row: &[String]) {
        todo!()
    }
}
```

### Usage

Here is a possible program to test your function:

```rust
use display_table::*;

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
