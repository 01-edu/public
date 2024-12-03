## division_and_remainder

### Instructions

Create a **function** named `divide` that receives two `i32` and returns a `tuple`. The first element is the result of the integer division between the two numbers, and the second is the remainder of the division.

```rust
pub fn divide(x: i32, y: i32) -> (i32, i32) {
}
```

### Usage

Here is a program to test your function

```rust
use division_and_remainder::*;

fn main() {
    let x = 9;
    let y = 4;
    let (division, remainder) = divide(x, y);
    println!(
        "{}/{}: division = {}, remainder = {}",
        x, y, division, remainder
    );
}
```

And its output

```console
$ cargo run
9/4: division = 2, remainder = 1
$
```

### Notions

- [The Tuple Type](https://doc.rust-lang.org/stable/book/ch03-02-data-types.html?highlight=accessing%20a%20tuple#compound-types)

- [Tuples](https://doc.rust-lang.org/rust-by-example/primitives/tuples.html)

- [Tuple Structs without Named Fields](https://doc.rust-lang.org/stable/book/ch05-01-defining-structs.html?highlight=tuple#using-tuple-structs-without-named-fields-to-create-different-types)
