## arrays

### Instructions

Define a **function** named `thirtytwo_tens` that returns an array with 32 positions filled with only the value `10`, so that `[10, 10, 10, ... 10].len()` is equal to 32.

Write a **function** that takes an array of `i32` and returns the sum of the elements (make it work with the main).

### Expected functions

The type of one of the arguments is missing. Use the example `main` function to determine the correct type.

```rust
pub fn sum(a: _) -> i32 {
	//type of argument missing in the signature here
}

pub fn thirtytwo_tens() -> [i32; 32] {
}
```

### Usage

Here is a program to test your function.

> It's incomplete. Use the output and the other available information to figure out what is missing.

```rust
use arrays::{sum, thirtytwo_tens};

fn main() {
	let a = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
	let a1: Vec<i32> = (1..11).; //missing info here
	let b = [_; 10]; //missing info here

	println!("The Sum of the elements in {:?} = {}", a, sum(a));//missing info here
	println!("The Sum of the elements in {:?} = ", a1, sum(a1));//missing info here
	println!("The Sum of the elements in {:?} = {}", b, sum(b));//missing info here
	println!(
		"Array size {} with only 10's in it {:?}",
		thirtytwo_tens().len(),
		thirtytwo_tens()
	);
}
```

And its output:

```console
$ cargo run
The Sum of the elements in [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] = 55
The Sum of the elements in [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] = 55
The Sum of the elements in [5, 5, 5, 5, 5, 5, 5, 5, 5, 5] = 50
Array size 32 with only 10's in it [10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10]
$
```

### Notions

- [arrays](https://doc.rust-lang.org/std/primitive.array.html)
