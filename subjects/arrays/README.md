## arrays

### Instructions

Define a function call thirtytwo_tens that returns an array with 32 positions fill with only the value 10: [10, 10, 10, ... 10].len()
= 32

Write a function that takes an array of i32 and returns the sum of the elements (make it work with the main)

### Expected functions

The type of one of the arguments is missing use the example `main` function to determine the correct type.

```rust
fn sum(a: _) -> i32 {
}

fn thirtytwo_tens() -> [i32; 32] {
}
```

### Usage

Here is a program to test your function.

There is things missing in this program use the output and the other information that you have available to determine what is missing.

```rust
use arrays::{sum, thirtytwo_tens};

fn main() {
	let a = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
	let a1: Vec<i32> = (1..11).;
	let b = [_; 10];

	println!("The Sum of the elements in {:?} = {}", a, sum(a));
	println!("The Sum of the elements in {:?} = ", a1, sum(a1));
	println!("The Sum of the elements in {:?} = {}", b, sum(b));
	println!(
		"Array size {} with only 10's in it {:?}",
		thirtytwo_tens().len(),
		thirtytwo_tens()
	);
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
The Sum of the elements in [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] = 55
The Sum of the elements in [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] = 55
The Sum of the elements in [5, 5, 5, 5, 5, 5, 5, 5, 5, 5] = 50
Array size 32 with only 10's in it [10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10]
student@ubuntu:~/[[ROOT]]/test$
```
