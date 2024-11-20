## nextprime

### Instructions

Create a **function** which returns the first prime number which is greater than or equal to the `u64` passed as an argument.

The function must be optimized, so as to avoid time-outs.

> A prime number is a natural number greater than 1 that is a not a product of two smaller natural numbers.
> 4 is not a prime number (so it's called a composite number) because it can be represented as 2 _ 2. 5 is a prime number as it can only be represented by 5 _ 1 or 1 \* 5.

### Expected function

```rust
pub fn next_prime(nbr: u64) -> u64 {

}
```

### Usage

Here is a possible program to test your function :

```rust
use nextprime::*;

fn main() {
    println!("The next prime after 4 is: {}", next_prime(4));
    println!("The next prime after 11 is: {}", next_prime(11));
}

```

And its output :

```console
$ cargo run
The next prime after 4 is: 5
The next prime after 11 is: 11
$
```
