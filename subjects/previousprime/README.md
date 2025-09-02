## prev_prime

### Instructions

Create a **function** which returns the biggest prime number which is smaller than the `u64` passed as an argument.

If there are no smaller primes, the function should return `0`.

> A prime number is a natural number greater than 1 that is a not a product of two smaller natural numbers.
> 4 is not a prime number (so it's called a composite number) because it can be represented as 2 \* 2. 5 is a prime number as it can only be represented by 5 \* 1 or 1 \* 5.

### Expected function

```rust
pub fn prev_prime(nbr: u64) -> u64  {
    todo!()
}
```

### Usage

Here is a possible program to test your function:

```rust
use previousprime::*;

fn main() {
    println!("The previous prime number before 34 is: {}", prev_prime(34));
}
```

And its output:

```console
$ cargo run
The previous prime number before 34 is: 31
$
```
