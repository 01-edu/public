## prime_checker

### Instructions

Create a **function** `prime_checker` that takes an `u32` and check if it is a prime number.

The result will be `None` if the argument is less or equal one otherwise, it will return the result with the prime number inside if the number is prime, or it will return an `enum` `PrimeErr` if the `u32` is not prime.

The `enum` `PrimeErr` will be `Even` if the number is even or `Divider(u32)` where the `u32` is the smaller divider of the number if it is not prime.

### Expected Function and structure

```rust
#[derive(PartialEq, Eq, Debug)]
pub enum PrimeErr {
    Even,
    Divider(u32),
}

pub fn prime_checker(nb: u32) -> Option<Result<u32, PrimeErr>> {}
```

### Usage

Here is a program to test your function:

```rust
use prime_checker::*;

fn main() {
    println!(
        "Is {} prime? {}",
        2,
        match prime_checker(2) {
            Some(res) => {
                match res {
                    Ok(_) => "Yes, it is!",
                    Err(_) => "No, it is not!",
                }
            }
            None => "Less than 2",
        }
    );
    println!(
        "Is {} prime? {}",
        14,
        match prime_checker(14) {
            Some(res) => {
                match res {
                    Ok(_) => "Yes, it is!",
                    Err(_) => "No, it is not!",
                }
            }
            None => "Less than 2",
        }
    );
}
```

And its output:

```console
$ cargo run
Is 2 prime? Yes, it is!
Is 14 prime? No, it is not!
$
```
