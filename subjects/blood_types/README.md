## blood_types

### Instructions

In this exercise you will create a data model of blood types and an API to deal with them.

First, we'll implement the data representation of the blood types.

Take a look at the `BloodType` struct below. It contains two enums which enable us to describe a blood type (e.g. "A-").

- `Antigen`: which can be one of:
  - `A`
  - `B`
  - `AB`
  - `O`
- `RhFactor`: which can be one of:
  - `Positive`
  - `Negative`

To provide a simple way to create blood types, implement the trait `FromStr` for `BloodType`. This will allow us to use the `parse` method and `from_str`, so we can do:

```rust
let a_neg: BloodType = "A-".parse();
```

Implement the following Traits:

- `std::fmt::Debug`: for `BloodType`, allowing us to print a vector such as `[BloodType { antigen: A, rh_factor: Positive }, BloodType { antigen: B, rh_factor: Negative }]` as `"[A+, B-]"` when using format strings `{:?}`.

Create three methods for BloodType:

- `can_receive_from`: which returns whether `self` can receive blood from `other` blood type.
- `donors`: which returns all the blood types that can give blood to `self`.
- `recipients`: which returns all the blood types that can receive blood from `self`.

### Expected Functions and Structures

```rust
use std::{fmt, str::FromStr};

#[derive(PartialEq, Eq, Hash, Clone, Copy)]
pub enum Antigen {
	A,
	AB,
	B,
	O,
}

#[derive(PartialEq, Eq, Hash, Clone, Copy)]
enum RhFactor {
	Positive,
	Negative,
}

#[derive(PartialEq, Eq, Hash, Clone, Copy)]
pub struct BloodType {
	pub antigen: Antigen,
	pub rh_factor: RhFactor,
}

impl FromStr for BloodType {
	type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
		todo!()
    }
}

impl fmt::Debug for BloodType {
	fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
		todo!()
    }
}

impl BloodType {
	pub fn can_receive_from(self, other: Self) -> bool {
		todo!()
	}

	pub fn donors(self) -> Vec<Self> {
		todo!()
	}

	pub fn recipients(self) -> Vec<Self> {
		todo!()
	}
}
```

### Usage

Here is a program to test your function.

```rust
use blood_types::*;

fn main() {
	let blood_type = "O+".parse::<BloodType>().unwrap();
	println!("recipients of O+ {:?}", blood_type.recipients());
	println!("donors of O+ {:?}", blood_type.donors());

	let another_blood_type = "A-".parse::<BloodType>().unwrap();
	println!(
		"donors of O+ can receive from {:?} {:?}",
		another_blood_type,
		blood_type.can_receive_from(another_blood_type)
	);
}
```

And its output

```console
$ cargo run
recipients of O+ [AB+, O+, A+, B+]
donors of O+ [O+, O-]
donors of O+ can receive from A- false
$
```
