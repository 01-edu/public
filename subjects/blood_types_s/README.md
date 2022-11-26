## blood_types_s

### Instructions

Use the following table to define the methods asked:

| Blood Types | Donate Blood to  | Receive Blood From |
| ----------- | ---------------- | ------------------ |
| A+          | A+, AB+          | A+, A-, O+, O-     |
| O+          | O+, A+, B+, AB+  | O+, O-             |
| B+          | B+, AB+          | B+, B-, O+, O-     |
| AB+         | AB+              | Everyone           |
| A-          | A+, A-, AB+, AB- | A-, O-             |
| O-          | Everyone         | O-                 |
| B-          | B+, B-, AB+, AB- | B-, O-             |
| AB-         | AB+, AB-         | AB-, A-, B-, O-    |

Implement three methods for `BloodType`:

- `can_receive_from`: returns `true` if `self` **can** receive blood from `other` blood type.
- `donors`: which returns all the blood types that can give blood to `self`.
- `recipients`: which returns all the blood types that can receive blood from `self`.

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
pub enum RhFactor {
	Positive,
	Negative,
}

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Clone)]
pub struct BloodType {
	pub antigen: Antigen,
	pub rh_factor: RhFactor,
}

impl BloodType {
	pub fn can_receive_from(&self, other: &Self) -> bool {
	}

	pub fn donors(&self) -> Vec<Self> {
	}

	pub fn recipients(&self) -> Vec<Self> {
	}
}
```

### Usage

Here is a program to test your function.

```rust
use blood_types_s::{Antigen, BloodType, RhFactor};

fn main() {
	let blood_type = BloodType {
		antigen: Antigen::O,
		rh_factor: RhFactor::Positive,
	};
	println!("recipients of O+ {:?}", blood_type.recipients());
	println!("donors of O+ {:?}", blood_type.donors());
	let another_blood_type = BloodType {
		antigen: Antigen::O,
		rh_factor: RhFactor::Positive,
	};
	println!(
		"donors of O+ can receive from {:?} {:?}",
		&another_blood_type,
		blood_type.can_receive_from(&another_blood_type)
	);
}
```

And its output

```console
$ cargo run
recipients of O+ [BloodType { antigen: AB, rh_factor: Positive }, BloodType { antigen: O, rh_factor: Positive }, BloodType { antigen: A, rh_factor: Positive }, BloodType { antigen: B, rh_factor: Positive }]
donors of O+ [BloodType { antigen: O, rh_factor: Positive }, BloodType { antigen: O, rh_factor: Negative }]
donors of O+ can receive from BloodType { antigen: O, rh_factor: Positive } true
$
```
