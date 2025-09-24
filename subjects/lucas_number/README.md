## lucas_number

### Instructions

Complete the body of the **function** `lucas_number`.

```rust
pub fn lucas_number(n: u32) -> u32 {
    todo!()
}
```

This function receives a number `n` and returns the `n`th number in the Lucas Numbers where the `n`th number is the sum of the previous two numbers in the series.

The Lucas Numbers start like this: 2, 1, 3, 4, 7, 11, 18, 29, 47, 76, 123, etc...

### Usage

Here is a possible test for your function:

```rust
use lucas_number::*;

fn main() {
    println!(
        "The element in the position {} in Lucas Numbers is {}",
        2,
        lucas_number(2)
    );
    println!(
        "The element in the position {} in Lucas Numbers is {}",
        5,
        lucas_number(5)
    );
    println!(
        "The element in the position {} in Lucas Numbers is {}",
        10,
        lucas_number(10)
    );
    println!(
        "The element in the position {} in Lucas Numbers is {}",
        13,
        lucas_number(13)
    );
}
```

And its output:

```console
$ cargo run
The element in the position 2 in Lucas Numbers is 3
The element in the position 5 in Lucas Numbers is 11
The element in the position 10 in Lucas Numbers is 123
The element in the position 13 in Lucas Numbers is 521
$
```
