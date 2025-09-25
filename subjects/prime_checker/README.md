## prime_checker

### Instructions

Create a **function** `prime_checker` that takes an `usize` and check if it is a prime number.

The result will be `None` if the argument is less than or equal to one, otherwise it will return a `Result`.
If the `usize` is prime, the function will return an `Ok(usize)`. For any other case it will return an `Err(PrimeErr)`.
The `enum` `PrimeErr` will be `Even` if the number is a multiple of two, or `Divider(usize)` where the `usize` is the smallest divider of the number.

> Your solution should be optimized to a certain degree.

### Expected Function and structure

```rust
#[derive(PartialEq, Eq, Debug)]
pub enum PrimeErr {
    Even,
    Divider(usize),
}

pub fn prime_checker(nb: usize) ->  /* Implement return type here */ {
    todo!()
}
```

### Usage

Here is a program to test your function:

```rust
use prime_checker::*;

fn main() {
    println!("Is {} prime? {:?}", 2, prime_checker(2));
    println!("Is {} prime? {:?}", 14, prime_checker(14));
    println!("Is {} prime? {:?}", 2147483647, prime_checker(2147483647));
}
```

And its output:

```console
$ cargo run
Is 2 prime? Some(Ok(2))
Is 14 prime? Some(Err(Even))
Is 2147483647 prime? Some(Ok(2147483647))
$
```
