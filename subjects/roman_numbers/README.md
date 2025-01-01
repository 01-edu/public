## roman_numbers

### Instructions

Implement the `From<u32>` trait to create a roman number from a number. The roman number should be in subtractive notation (the common way to write roman number I, II, III, IV, V, VI, VII, VIII, IX, X ...).

Start by defining the digits as `RomanDigit` (`Nulla` is 0).

Next define `RomanNumber` as a wrapper to a vector of `RomanDigit`, and implement the Trait `From<u32>`.

### Expected Functions and Data Structures

```rust
use crate::RomanDigit::*;

#[derive(Copy, Clone, Debug, PartialEq, Eq)]
pub enum RomanDigit {
	Nulla,
	I,
	V,
	X,
	L,
	C,
	D,
	M,
}

#[derive(Clone, Debug, PartialEq, Eq)]
pub struct RomanNumber(pub Vec<RomanDigit>);

impl From<u32> for RomanDigit {
}

impl From<u32> for RomanNumber {
}
```

### Usage

Here is a program to test your function.

```rust
use roman_numbers::RomanNumber;

fn main() {
	println!("{:?}", RomanNumber::from(32));
	println!("{:?}", RomanNumber::from(9));
	println!("{:?}", RomanNumber::from(45));
	println!("{:?}", RomanNumber::from(0));
}
```

And its output:

```console
$ cargo run
RomanNumber([X, X, X, I, I])
RomanNumber([I, X])
RomanNumber([X, L, V])
RomanNumber([Nulla])
$
```
