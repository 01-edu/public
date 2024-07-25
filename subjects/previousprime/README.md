## prev_prime

### Instructions

Create a **function** which returns the first prime number which is less than or equal to the `u64` passed as an argument.

If there are no smaller primes, the function should return `0`.

> [!TIP]
> As a reminder, a `prime number` is a whole number greater than 1 that cannot be exactly divided by any whole number other than itself and 1.

> We consider that only positive numbers can be prime numbers.

### Expected function

```rust
pub fn prev_prime(nbr: u64) -> u64  {

}
```

### Usage

Here is a possible program to test your function :

```rust
use previousprime::*;

fn main() {
    println!("The previous prime number before 34 is: {}", prev_prime(34));
}
```

And its output :

```console
$ cargo run
The previous prime number before 34 is: 31
$
```
