## step_iterator

### Instructions

- Create an `Iterator` (by implementing the `std::iter::Iterator` trait) that iterates through the values from `beg` to `end` (including end) by `steps`.
  - The name of your iterator will be `StepIterator` and it must be generic enough so that you can use any integer value or floating point number
  - If the steps do not allow to attain the end of the sequence, only the last value inferior to the end of the series will be returned (See Usage)

- Define the associated function `new` which creates a new Step iterator

### Expected Functions and Structures

```rust
pub struct StepIterator<T> {}

impl StepIterator<T> {
    pub fn new(beg: T, end: T, step: T) -> Self {
        todo!()
    }
}

impl Iterator for StepIterator<T> {}
```

### Usage

Here is a program to test your function.

```rust
use step_iterator::*;

fn main() {
    println!("{:?}", StepIterator::new(0, 100, 10).collect::<Vec<_>>());
    println!("{:?}", StepIterator::new(0, 100, 12).collect::<Vec<_>>());
}
```

And its output:

```console
$ cargo run
[0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100]
[0, 12, 24, 36, 48, 60, 72, 84, 96]
$
```
