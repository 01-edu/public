## organize_garage

### Instructions

Create a structure `Garage` with generic values. It must derive at least `Debug`, `PartialEq`, and `Eq`, and will have the following public fields:

- `left` as `Option<T>`.
- `right` as `Option<T>`.

It will implement the following public methods:

- `move_to_right`: Moves the values from left to right.
- `move_to_left`: Moves the values from right to left.

> The generic type will need to have `Add` and `Copy` traits implemented.

### Usage

Here is a program to test your function.

```rust
use organize_garage::*;

fn main() {
    let mut garage = Garage {
        left: Some(5),
        right: Some(2),
    };

    println!("{:?}", garage);
    garage.move_to_right();
    println!("{:?}", garage);
    garage.move_to_left();
    println!("{:?}", garage);
}
```

And its output

```console
$ cargo run
Garage { left: Some(5), right: Some(2) }
Garage { left: None, right: Some(7) }
Garage { left: Some(7), right: None }
$
```
