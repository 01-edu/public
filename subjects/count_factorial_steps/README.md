## count_factorial_steps

### Instructions

Create a **function** named `count_factorial_steps` that receives a factorial number and counts how many multiplications are necessary to have this number.

If the argument is not a factorial, or it is equal 0 or 1, then the function should return 0.

```rust
pub fn count_factorial_steps(factorial: u64) -> u64 {
    todo!()
}
```

As a reminder, the factorial of a number is the product of all the integers from 1 to that number.

Example: the factorial of 6 (written 6!) is 1 \* 2 \* 3 \* 4 \* 5 \* 6 = 720. As such, the factorial steps of 720 are 6.

### Usage

Here is a possible program to test your function:

```rust
use count_factorial_steps::*;

fn main() {
    println!(
        "The factorial steps of 720 = {}",
        count_factorial_steps(720)
    );
    println!("The factorial steps of 13 = {}", count_factorial_steps(13));
    println!("The factorial steps of 6 = {}", count_factorial_steps(6));
}

```

And its output:

```console
$ cargo run
The factorial steps of 720 = 6
The factorial steps of 13 = 0
The factorial steps of 6 = 3
```
