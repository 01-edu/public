## nextprime

### Instructions

Write a **function** which returns the first prime number which is equal or superior to the `u64` passed as parameter.

The function must be optimized in order to avoid time-outs with the tester.

(We consider that only positive numbers can be prime numbers)

### Expected function

```rust
pub fn next_prime(nbr: u64) -> u64 {

}
```

### Usage

Here is a possible program to test your function :

```rust
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
