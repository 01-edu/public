## arrays

### Instructions

Define a **function** named `thirtytwo_tens` that returns an array with 32 positions filled with only the value `10`, so that `[10, 10, 10, ... 10].len()` is equal to 32.

Write a **function** that takes an array of `i32` and returns the sum of the elements (make it work with the main).

### Expected functions

The type of one of the arguments is missing. Use the example `main` function to determine the correct type.

```rust
pub fn sum(a: _) -> i32 {
}

pub fn thirtytwo_tens() -> [i32; 32] {
}
```

### Usage

Here is a program to test your function.

> It's incomplete. Use the output and the other available information to figure out what is missing.

```rust
use arrays::*;

fn main() {
    let a = (1..=10)._;
    let b = [_];

    println!("The sum of the elements in {:?} is {}", a, sum(&a));
    println!("The sum of the elements in {:?} is {}", b, sum(&b));
    println!(
        "Array of {} elements filled with 10 = {:?}",
        thirtytwo_tens().len(),
        thirtytwo_tens()
    );
}
```

And its output:

```console
$ cargo run
The sum of the elements in [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] is 55
The sum of the elements in [5, 5, 5, 5, 5, 5, 5, 5, 5, 5] is 50
Array of 32 elements filled with 10 = [10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10]
$
```

### Notions

- [arrays](https://doc.rust-lang.org/std/primitive.array.html)
