## prev_prime

### Instructions

Write a function that returns the first prime number that is equal or inferior to the `int` passed as parameter.

If there are no primes inferior to the `int` passed as parameter the function should return 0.

### Expected function

```rust
pub fn prev_prime(nbr: u64) -> u64  {

}
```

### Usage

Here is a possible program to test your function :

```rust
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
