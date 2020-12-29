## step_iterator

### Instructions

- Create an Iterator (by implementing the `std::iter::Iterator` trait) that iterates through the values from `beg` to `end` (including end) in the indicated `steps`.

  - The name of you're iterator will be `StepIterator` and it must be generic so you can use any integer value: i8,..,i64, u8,..,u64 or floating point number f32,..,f64

  - If the steps don't allow to arrive until the end of the sequence only the last value inferior to the end of the series will be returned (See Usage)

- Define the associated function: `new` that creates a new Step iterator:

### Expected Functions and Structures

```rust
pub struct StepIterator<T> {
...
}

use std::ops::Add;
impl StepIterator<T> {
	pub fn new(beg: T, end: T, step: T) -> Self {
	}
}

impl std::iter::Iterator for StepIterator<T> {
}
```

### Usage

Here is a program to test your function.

```rust
fn main() {
	for v in StepIterator::new(0, 100, 10) {
		print!("{},", v);
	}
	println!();

	for v in StepIterator::new(0, 100, 12) {
		print!("{},", v)
	}
	println!();
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
0,10,20,30,40,50,60,70,80,90,100,
0,12,24,36,48,60,72,84,96,
student@ubuntu:~/[[ROOT]]/test$
```
