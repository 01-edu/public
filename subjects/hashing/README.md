## hashing

### Instructions

Given a list of integers (`Vec<i32>`) write three **functions**.

- `mean`: that calculates the mean (the average value) of all the values in the list.

- `median`: that calculates the median (for a sorted list, it is the value in the middle). If there is an even amount of numbers in the list, the middle pair must be determined, added together, and divided by two to find the median value.

- `mode` that calculates the mode (the value
that appears more often).

### Expected Functions

```rust
pub fn mean(list: &Vec<i32>) -> f64 {
}

pub fn median(list: &Vec<i32>) -> i32 {
}

pub fn mode(list: &Vec<i32>) -> i32 {
}
```

### Usage

Here is a program to test your function.

```rust
use hashing::*;

fn main() {
	println!("Hello, world!");
	let v = vec![4, 7, 5, 2, 5, 1, 3];
	println!("mean {}", hashing::mean(&v));
	println!("median {}", hashing::median(&v));
	println!("mode {}", hashing::mode(&v));
}
```

And its output;

```console
$ cargo run
mean 3.857142857142857
median 4
mode 5
$
```

### Notions

- [hash maps](https://doc.rust-lang.org/book/ch08-03-hash-maps.html)
