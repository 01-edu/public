## armstrong_number

### Instructions

Create a function which checks if the number is a valid Armstrong number. The function will return the number if it is a valid one and `None` otherwise.

An Armstrong number is a number where the sum of its own digits each raised to the power of the number of digits is equal to itself.

As an example 153 is an Armstrong number because:
`153 = 1^3 + 5^3 + 3^3`

### Expected Function

```rust
pub fn is_armstrong_number(nb: u32) -> Option<u32> {
}
```

### Usage

Here is a possible program to test your function:

```rust
use armstrong_number::*;

fn main() {
    println!("{:?}", is_armstrong_number(0));
    println!("{:?}", is_armstrong_number(1));
    println!("{:?}", is_armstrong_number(153));
    println!("{:?}", is_armstrong_number(370));
    println!("{:?}", is_armstrong_number(371));
    println!("{:?}", is_armstrong_number(407));
    println!("{:?}", is_armstrong_number(400));
    println!("{:?}", is_armstrong_number(198));
}
```

And its output:

```console
$ cargo run
Some(0)
Some(1)
Some(153)
Some(370)
Some(371)
Some(407)
None
None
$
```
