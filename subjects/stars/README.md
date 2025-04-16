## stars

### Instructions

Create a function named `stars` that takes a `u32` as an argument and returns a `String` of stars (asterisks). The number of stars returned is 2^n (2 to the `n`th power).

### Expected functions

```rust
pub fn stars(n: u32) -> String {

}
```

### Usage

Here is a program to test your function.

```rust
use stars::stars;

fn main() {
    println!("{}", stars(1));
    println!("{}", stars(4));
    println!("{}", stars(5));
}
```

And its output:

```console
$ cargo run
**
****************
********************************
$
```
