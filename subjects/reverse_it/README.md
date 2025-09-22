## reverse_it

### Instructions

Create a function named `reverse_it`, that takes a number. It should return a string with the number reversed, followed by the original number. If the number is negative, a `-` should be added to the beginning of the string.

### Expected Functions

```rust
pub fn reverse_it(v: i32) -> String {
    todo!()
}
```

### Usage

Here is a program to test your function:

```rust
use reverse_it::*;

fn main() {
    println!("{}", reverse_it(123));
    println!("{}", reverse_it(-123));
}
```

And its output:

```console
$ cargo run
321123
-321123
$
```
