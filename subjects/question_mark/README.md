## question_mark

### Instructions

Create the following structures. Each has one field:

- `One`: `first_layer` with type `Option<Two>`.
- `Two`: `second_layer` with type `Option<Three>`
- `Three`: `third_layer` with type `Option<Four>`
- `Four`: `fourth_layer` with type `Option<u16>`.

You must also create a **function** associated to the structure `One` called `get_fourth_layer`, which should return the `fourth_layer` value in the `Four` structure.

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
    pub fn get_fourth_layer(self) -> Option<u16> {
        todo!()
    }
}
```

### Usage

Here is a program to test your function:

```rust
use question_mark::*;

fn main() {
    let a = One {
        first_layer: Some(Two {
            second_layer: Some(Three {
                third_layer: Some(Four {
                    fourth_layer: Some(1000)
                })
            })
        })
    };

    println!("{:?}", a.get_fourth_layer());
}
```

And its output:

```console
$ cargo run
Some(1000)
$
```

### Notions

- [Unpacking options with ?](https://doc.rust-lang.org/stable/rust-by-example/error/option_unwrap/question_mark.html)
