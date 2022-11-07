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
    let nb_1 = 9;
    let nb_2 = 4;
    let nb_3 = 2;
    let (min, max) = min_and_max(nb_1, nb_2, nb_3);
    println!(
        "The minimum is {}, the maximum is {}",
        min, max
    );
}
```

And its output

```console
$ cargo run
The minimum is 2, the maximum is 9
$
```

### Notions

- [The Tuple Type](https://doc.rust-lang.org/stable/book/ch03-02-data-types.html?highlight=accessing%20a%20tuple#compound-types)

- [Tuples](https://doc.rust-lang.org/rust-by-example/primitives/tuples.html)

- [Tuple Structs without Named Fields](https://doc.rust-lang.org/stable/book/ch05-01-defining-structs.html?highlight=tuple#using-tuple-structs-without-named-fields-to-create-different-types)
