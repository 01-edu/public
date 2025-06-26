## find_factorial

### Instructions

Create a **function** named `factorial` which returns the factorial of a given number.

```rust
pub fn factorial(num: u64) -> u64 {
}
```

As a reminder, the factorial of a number is the product of all the integers from 1 to that number.

Example: the factorial of 6 (written 6!) is 1 \* 2 \* 3 \* 4 \* 5 \* 6 = 720.

> Do not forget the rules for 0 and 1.

### Usage

Here is a possible program to test your function:

```rust
use find_factorial::*;

fn main() {
    println!("The factorial of 0 = {}", factorial(0));
    println!("The factorial of 1 = {}", factorial(1));
    println!("The factorial of 5 = {}", factorial(5));
    println!("The factorial of 10 = {}", factorial(10));
    println!("The factorial of 19 = {}", factorial(19));
}
```

And its output:

```console
$ cargo run
The factorial of 0 = 1
The factorial of 1 = 1
The factorial of 5 = 120
The factorial of 10 = 3628800
The factorial of 19 = 121645100408832000
$
```
