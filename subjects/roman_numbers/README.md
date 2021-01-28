## roman_numbers

### Instructions
Implement the From<u32> Trait to create a roman number from a u32 the roman number should be in subtractive notation (the common way to write roman number I, II, III, IV, V, VI, VII, VIII, IX, X ...)

For this start by defining the digits as `RomanDigit` with the values I, V, X, L, C, D, M and Nulla for 0

Next define RomanNumber as a wrapper to a vector of RomanDigit's And implement the Trait From<u32>

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

#[derive(Debug, Clone, Debug, PartialEq, Eq)]
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

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
RomanNumber([X, X, X, I, I])
RomanNumber([I, X])
RomanNumber([X, L, V])
RomanNumber([Nulla])
student@ubuntu:~/[[ROOT]]/test$
```
