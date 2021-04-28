## roman_numbers_iter

### Instructions

Implement the `IntoIterator` trait for the `RomanNumber` type to enable using a for loop notation. This implementation must allow taking ownership, borrowing and borrowing mutably.

1. Taking ownership (this consumes the RomanNumber)

```rust
for digit in number {
	...
}
```

2. Borrowing immutably (this preserves the RomanNumber)

```rust
	for digit in &number {

	}
```

3. Borrowing mutably (this allow you to modify the RomanNumber without having to return the ownership)

```rust
	for digit in &mut number {

	}
```

### Notions

- https://doc.rust-lang.org/std/iter/trait.IntoIterator.html
- https://doc.rust-lang.org/std/iter/index.html

### Expected Functions

```rust
use roman_numbers::{RomanDigit, RomanNumber};

impl IntoIterator for &RomanNumber {
}

impl IntoIterator for &mut RomanNumber {
}

impl IntoIterator for RomanNumber {
}
```

### Usage

Here is a program to test your function.

```rust
use roman_numbers::RomanNumber;

fn main() {
	let number = RomanNumber::from(15);

	for digit in &number {
		println!("{:?}", digit);
	}
	println!("{:?}", number);
}
```

And its output

```console
$ cargo run
RomanNumber([X, X, X, I, I])
RomanNumber([I, X])
RomanNumber([X, L, V])
RomanNumber([Nulla])
$
```
