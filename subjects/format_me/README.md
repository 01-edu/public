## format_me

### Instructions

In this exercise you will implement the trait `Display` for the structure `Park` and the enum `ParkType`.

Here are the public fields and possible variants for the two types:

- `Park`:
  - `name` as `Option<String>`
  - `park_type` as `ParkType`
  - `address` as `Option<String>`
  - `cap` as `Option<String>`
  - `state` as `Option<String>`
- `ParkType`:
  - `Garden`
  - `Forest`
  - `Playground`

On implementing the `Display` trait for `Park`, you should display the park's fields using the following format: `{park_type} - {name}, {address}, {cap} - {state}`. If any of the fields `name`, `address`, `cap` or `state` are missing, display `No name`, `No address`, `No cap`, or `No state` respectively instead.

On implementing the `Display` trait for `ParkType`, you should display the park type as a string, in all lowercase.

### Expected Functions and Structures

```rust
use std::fmt;

pub struct Park {}

pub enum ParkType {}

impl fmt::Display for Park {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        todo!()
    }
}

impl fmt::Display for ParkType {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        todo!()
    }
}
```

### Usage

Here is a program to test your function:

```rust
use format_me::*;

fn main() {
    println!(
        "{}",
        Park {
            name: Some("Les Tuileries".to_owned()),
            park_type: ParkType::Garden,
            address: Some("Pl. de la Concorde".to_owned()),
            cap: Some("75001".to_owned()),
            state: Some("France".to_owned())
        }
    );
    println!(
        "{}",
        Park {
            name: None,
            park_type: ParkType::Playground,
            address: None,
            cap: None,
            state: None
        }
    );
}
```

And its output

```console
$ cargo run
garden - Les Tuileries, Pl. de la Concorde, 75001 - France
playground - No name, No address, No cap - No state
$
```
