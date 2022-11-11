## min_and_max

### Instructions

Create a **function** named `min_and_max` that receives three `i32` and returns a `tuple` with the minimum and the maximum number received as input.

```rust
pub fn min_and_max(nb_1: i32, nb_2: i32, nb_3: i32) -> (i32, i32) {
}
```

### Usage

Here is a program to test your function

```rust
use min_and_max::min_and_max;

fn main() {
    println!(
        "Minimum and maximum are: {:?}",
        min_and_max(9, 2, 4)
    );
}
```

And its output

```console
$ cargo run
Minimum and maximum are: (2, 9)
$
```

### Notions

- [The Tuple Type](https://doc.rust-lang.org/stable/book/ch03-02-data-types.html)

- [Tuples](https://doc.rust-lang.org/rust-by-example/primitives/tuples.html)

- [Tuple Structs without Named Fields](https://doc.rust-lang.org/stable/book/ch05-01-defining-structs.html)
