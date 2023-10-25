## nextprime

### Instructions

Create a **function** which returns the first prime number which is greater than or equal to the `u64` passed as an argument.

The function must be optimized, so as to avoid time-outs.

> We consider that only positive numbers can be prime numbers.

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
