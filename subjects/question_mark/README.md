## question_mark

### Instructions

You will have to create 3 structures:

- `One`, that contains one element called `first_layer` it should be an `Option` for the structure `Two`.
- `Two`, that contains one element called `second_layer` it should be an `Option` for the structure `Three`.
- `Three`, that contains one element called `third_layer` it should be an `Option` for the structure `Four`.
- `Four`, that contains one element called `fourth_layer` it should be an `u16` that is an `Option`.

Beside the structure you must create a function named `get_fourth_layer` that is associated to the `One` structure.
This function should return the `Option` value in the `Four` structure.

### Expected Function

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

Here is a program to test your function

```rust
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
student@ubuntu:~/[[ROOT]]/test$ cargo run
1000
student@ubuntu:~/[[ROOT]]/test$
```

### Notions

- https://doc.rust-lang.org/stable/rust-by-example/error/option_unwrap/question_mark.html
