## dress_code

### Instructions

Create a function called `choose_outfit` that receives the following input:

- A `formality_level` as an `Option<u32>`.
- An `invitation_message` as a `Result<&str, &str>`.

The function will return a struct `Outfit` which contains:

- `jacket`, an `enum` `Jacket` that contains `Black`, `White` and `Flowers`.
- `hat`, an `enum` `Hat` that contains `Snapback`, `Baseball`, `Fedora`.

```rust
#[derive(Debug, PartialEq, Eq)]
pub struct Outfit {
    pub jacket: Jacket,
    pub hat: Hat,
}
```

For the `jacket`:

- The jacket should be `Flowers` when the `formality_level` is `None`.
- The jacket should be `White` when the `formality_level` is more than 0.
- Otherwise, it should be `Black`.

For the `hat`:

- If the `invitation_message` is `Ok()` it should be `Fedora`.
- Otherwise, it should be `Snapback`.

In the specific case where `formality_level` is `None` and `invitation_message` is not `Ok()` then the `jacket` should be `Flowers` and the `hat` should be `Baseball`.

Remember that all the `enum` and `struct` used must be `pub`.

### Expected functions

```rust
pub fn choose_outfit(formality_level: Option<u32>, invitation_message: Result<&str, &str>) -> Outfit {
    todo!()
}
```

### Usage

Here is a program to test your function.

```rust
use dress_code::*;

fn main() {
    println!(
        "My outfit will be: {:?}",
        choose_outfit(Some(0), Ok("Dear friend, ..."))
    );
}
```

And its output:

```console
$ cargo run
My outfit will be: Outfit { jacket: Black, hat: Fedora }
$
```

### Notions

- [patterns](https://doc.rust-lang.org/book/ch18-00-patterns.html)
