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
pub struct Message {

}

impl Message {
  pub fn new(ms: String, u: String) -> Message {

  }
  pub fn send_ms(&self) -> Option<&str> {

  }
}

pub fn check_ms(ms: &Message) -> (bool, &str) {

}
```

### Usage

Here is a program to test your function

```rust
use profanity_filter::*;

fn main() {
  let m0 = Message::new("hello there".to_string(), "toby".to_string());
  println!("{:?}", check_ms(&m0));

  let m1 = Message::new("".to_string(), "toby".to_string());
  println!("{:?}", check_ms(&m1));

  let m2 = Message::new("you are stupid".to_string(), "toby".to_string());
  println!("{:?}", check_ms(&m2));

  let m3 = Message::new("stupid".to_string(), "toby".to_string());
  println!("{:?}", check_ms(&m3));
}
```

And its output:

```console
$ cargo run
(true, "hello there")
(false, "ERROR: illegal")
(false, "ERROR: illegal")
(false, "ERROR: illegal")
$
```

# Notions

- [Enum Definition](https://doc.rust-lang.org/stable/book/ch06-01-defining-an-enum.html?highlight=option#the-option-enum-and-its-advantages-over-null-values)
