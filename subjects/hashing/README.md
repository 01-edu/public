## hashing

### Instructions

Given a list of integers (Vec<i32>) write three functions

Write a function called `mean` that calculates the `mean` (the average value) of all the values in the list

Write a function called `median` that calculates the `median` (for a sorted list is the value in the middle)

Write a function called `mode` that calculates the mode (the value
that appears more often)

### Expected Functions

```rust
fn mean(list: &Vec<i32>) -> f64 {
}

fn median(list: &Vec<i32>) -> i32 {
}

fn mode(list: &Vec<i32>) -> i32 {
}
```

### Usage

Here is a program to test your function.

```rust
use hashing;

fn main() {
	println!("Hello, world!");
	let v = vec![4, 7, 5, 2, 5, 1, 3];
	println!("mean {}", hashing::mean(&v));
	println!("median {}", hashing::median(&v));
	println!("mode {}", hashing::mode(&v));
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
mean 3.857142857142857
median 4
mode 5
student@ubuntu:~/[[ROOT]]/test$
```
