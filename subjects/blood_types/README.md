## blood_types

### Instructions

In this exercise you will create a data model of blood types and an API to deal with them.

Start by copying the data representation of the blood types:

- Create the enum `Antigen` that has 4 possibilities: A, B, O and AB And the enum `RhFactor` that has two possible values: Positive and Negative

- After, copy the struct `BloodType` that contains two fields with the names antigen and rh_factor

- To provide a simple way to create blood types implement the trait FromStr for BloodType (which will allow us to use the `parse` method and the associated function from_str, so we can do:

```rust
   let a_neg: BloodType = "A-".parse();
```

- Implement the std::cmp::Ord trait to make possible to sort a vector or array of BloodType's

- Implement the trait std::Debug for BloodType allowing to print a vector such as [BloodType { antigen: A, rh_factor: Positive}, BloodType{ antigen: B, rh_factor: Negative}] as [ A+, A-] using the formatting {:?}

- Lastly, write three methods for BloodType:

  - `can_receive_from`: which returns true if `self` can receive blood from `other` blood type
  - `donors`: which returns all the blood types that can give blood to `self`
  - `recipients`: which returns all the blood types that can receive blood from `self`

### Expected Functions and Structures

```rust
#[derive(Debug, PartialEq, Eq, Clone, PartialOrd, Ord)]
pub enum Antigen {
	A,
	AB,
	B,
	O,
}

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Clone)]
enum RhFactor {
	Positive,
	Negative,
}

#[derive(PartialEq, Eq, PartialOrd)]
pub struct BloodType {
	pub antigen: Antigen,
	pub rh_factor: RhFactor,
}

use std::cmp::{Ord, Ordering};

use std::str::FromStr;

impl FromStr for Antigen {
}

impl FromStr for RhFactor {
}

impl Ord for BloodType {
}

impl FromStr for BloodType {
}

use std::fmt::{self, Debug};

impl Debug for BloodType {
}

impl BloodType {
	pub fn can_receive_from(&self, other: &Self) -> bool {
	}

	pub fn donors(&self) -> Vec<Self> {
	}

	pub fn recipients(&self) -> Vec<BloodType> {
}
```

### Usage

Here is a program to test your function.

```rust
use blood_types::*;

fn main() {
	let blood_type: BloodType = "O+".parse().unwrap();
	println!("recipients of O+ {:?}", blood_type.recipients());
	println!("donors of O+ {:?}", blood_type.donors());
	let another_blood_type: BloodType = "A-".parse().unwrap();
	println!(
		"donors of O+ can receive from {:?} {:?}",
		&another_blood_type,
		blood_type.can_receive_from(&another_blood_type)
	);
}
```

And its output

```console
student@ubuntu:~/blood_types/test$ cargo run
recipients of O+ [AB+, O+, A+, B+]
donors of O+ [O+, O-]
donors of O+ can receive from A- false
student@ubuntu:~/blood_types/test$
```
