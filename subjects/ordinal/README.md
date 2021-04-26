## ordinal

### Instructions

Complete the function "num_to_ordinal" that receives a cardinal number and returns its ordinal number.

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
student@ubuntu:~/ordinal/test$ cargo run
1st
22nd
43rd
47th
student@ubuntu:~/ordinal/test$
```
