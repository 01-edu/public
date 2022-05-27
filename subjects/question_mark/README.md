## question_mark

### Instructions

Create the following structures. Each has one element:

- `One`: `first_layer` as type `Option<Two>`.
- `Two`: `second_layer` as type `Option<Three>`
- `Three`: `third_layer` as type `Option<Four>`
- `Four`: `fourth_layer` as type `Option<u16>`.

Beside the structures, you must create a **function** named `get_fourth_layer`, and associate it to the `One` structure. This **function** should return the `Option` value in the `Four` structure.

### Expected Function and structures

```rust
pub struct One {
    // expected public fields
}
pub struct Two {
    // expected public fields
}
pub struct Three {
    // expected public fields
}
pub struct Four {
    // expected public fields
}

impl One {
    pub fn get_fourth_layer(&self) -> Option<u16> {}
}
```

### Usage

Here is a program to test your function:

```rust
use question_mark::*;

fn main() {
    let a = One {
        first_layer : Some(Two {
            second_layer: Some(Three {
                third_layer: Some(Four {
                    fourth_layer: Some(1000)
                })
            })
        })
    };

    // output: 1000
    println!("{:?}", match a.get_fourth_layer() {
        Some(e) => e,
        None => 0
    })
}
```

And its output:

```console
$ cargo run
1000
$
```

### Notions

- [Unpacking options with ?](https://doc.rust-lang.org/stable/rust-by-example/error/option_unwrap/question_mark.html)
