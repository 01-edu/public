## question_mark

### Instructions

3 structure have to be created:

- `One`:
  - which contains one element called `first_layer` which should be an `Option` of the structure `Two`.
- `Two`:
  - which contains one element called `second_layer` which should be an `Option` of the structure `Three`.
- `Three`:
  - which contains one element called `third_layer` which should be an `Option` of the structure `Four`.
- `Four`:
  - which contains one element called `fourth_layer` which is an `Option<u16>`.

Beside the structure you must create a **function** named `get_fourth_layer` which is associated to the `One` structure (a method).
This **function** should return the `Option` value in the `Four` structure.

### Notions

- [Unpacking options with ?](https://doc.rust-lang.org/stable/rust-by-example/error/option_unwrap/question_mark.html)

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
