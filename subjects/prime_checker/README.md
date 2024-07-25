## prime_checker

### Instructions

Create a **function** `prime_checker` that takes an `u32` and check if it is a prime number.

The result will be `None` if the argument is less or equal one, otherwise it will return a `Result`.
If the `u32` is prime, the function will return an`Ok(u32)`. For any other case it will return an `enum` `PrimeErr`.
The `enum` `PrimeErr` will be `Even` if the number is a multiple of two or `Divider(u32)` where the `u32` is the smallest divider of the number.

> [!TIP]
> As a reminder, a `prime number` is a whole number greater than 1 that cannot be exactly divided by any whole number other than itself and 1.

> We consider that only positive numbers can be prime numbers.

> Your solution should be optimized to a certain degree.

### Expected Function and structure

```rust
#[derive(PartialEq, Eq, Debug)]
pub enum PrimeErr {
    Even,
    Divider(u32),
}

pub fn prime_checker(nb: u32) ->  /* Implement return type here */ {}
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
