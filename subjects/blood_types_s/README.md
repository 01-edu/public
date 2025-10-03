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
#[derive(Debug, PartialEq, Eq, Hash, Clone, Copy)]
pub enum Antigen {
    A,
    AB,
    B,
    O,
}

#[derive(Debug, PartialEq, Eq, Hash, Clone, Copy)]
pub enum RhFactor {
    Positive,
    Negative,
}

#[derive(Debug, PartialEq, Eq, Hash, Clone, Copy)]
pub struct BloodType {
    pub antigen: Antigen,
    pub rh_factor: RhFactor,
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
use blood_types_s::*;

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
        "donors of O+ can receive from {:?} {}",
        another_blood_type,
        blood_type.can_receive_from(another_blood_type)
    );
}
```

And its output

```console
$ cargo run
recipients of O+ [BloodType { rh_factor: Positive, antigen: A }, BloodType { rh_factor: Positive, antigen: AB }, BloodType { rh_factor: Positive, antigen: B }, BloodType { rh_factor: Positive, antigen: O }]
donors of O+ [BloodType { rh_factor: Positive, antigen: O }, BloodType { rh_factor: Negative, antigen: O }]
donors of O+ can receive from BloodType { rh_factor: Positive, antigen: O } true
$
```
