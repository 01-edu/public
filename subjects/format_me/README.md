## format_me

### Instructions

In this exercise you will implement the trait `Display` for the structure `Park` and the enum `ParkType`.

Here are the public fields and possible variants for the two types:

- `Park`:
  - `name` as `String`
  - `park_type` as `ParkType`
  - `address` as `String`
  - `cap` as `String`
  - `state` as `String`
- `ParkType`:
  - `Garden`
  - `Forest`
  - `Playground`

### Expected Functions and Structures

```rust
use std::fmt;

pub struct Park {}
  
pub enum ParkType {}

impl fmt::Display for Park {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {}
}

impl fmt::Display for ParkType {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {}
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
            name: "Les Tuileries".to_string(),
            park_type: ParkType::Garden,
            address: "Pl. de la Concorde".to_string(),
            cap: "75001".to_string(),
            state: "France".to_string()
        }
    );
    println!(
        "{}",
        Park {
            name: "".to_string(),
            park_type: ParkType::Playground,
            address: "".to_string(),
            cap: "".to_string(),
            state: "".to_string()
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
