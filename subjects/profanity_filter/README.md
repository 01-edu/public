## profanity filter

### Instructions

Sometimes it is more desirable to catch the failure of some parts of a program instead of just calling panic.

For this exercise you will have to create a message blocker, where you must block the word `stupid`.

You will have to create a structure called `Message`, which contains:

- elements:
  - `content`: `String`
  - `user`: `String`
- associated functions:
  - `new`: which initializes the structure.
  - `send_ms`: which only has its implementation type (**self**) as argument. It should return `None` if the content of the message is either **empty** or contains the word **stupid**. It should return the content of the message otherwise.

You will also need to create a **function** named `check_ms` which accepts a reference to a `Message`, and returns a tuple. This function will invoke the `send_ms` function.

- If `send_ms` returns `None`, then your function should return `false` and `"ERROR: illegal"`.
- Else, your function should return `true` and the contents of the message sent.

##### Expected Function

```rust
pub fn check_ms(message: &str) -> Result<&str, &str> {
    todo!()
}
```

### Usage

Here is a program to test your function

```rust
use profanity_filter::*;

fn main() {
    ["hello there", "", "you are stupid", "stupid"]
        .into_iter()
        .for_each(|m| println!("{:?}", profanity_filter::check_ms(m)));}
```

And its output:

```console
$ cargo run
Ok("hello there")
Err("ERROR: illegal")
Err("ERROR: illegal")
Err("ERROR: illegal")
$
```

### Notions

- [Enum Definition](https://doc.rust-lang.org/stable/book/ch06-01-defining-an-enum.html?highlight=option#the-option-enum-and-its-advantages-over-null-values)
