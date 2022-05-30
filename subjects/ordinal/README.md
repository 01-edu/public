## ordinal

### Instructions

Complete the function `num_to_ordinal`. It returns the ordinal number for a given cardinal number.

### Expected functions

```rust
pub fn num_to_ordinal(x: u32) -> String {

}
```

### Usage

Here is a program to test your function.

```rust
use ordinal::*;

fn main() {
    println!("{}", num_to_ordinal(1));
    println!("{}", num_to_ordinal(22));
    println!("{}", num_to_ordinal(43));
    println!("{}", num_to_ordinal(47));
}
```

And its output

```console
$ cargo run
1st
22nd
43rd
47th
$
```
