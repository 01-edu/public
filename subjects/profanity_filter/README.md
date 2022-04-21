## profanity filter

### Instructions

Sometimes it is more desirable to catch the failure of some parts of a program instead of just calling panic.

For this exercise you will have to create a message blocker, where you must block the word `stupid`.

You will have to create a structure called `Message`, this structure
must have the following elements:

- content: String
- user: String

The `struct` must also have an implementation of 2 **functions** associated with it:

- `new`, which initializes the structure
- `send_ms`, which only has its implementation type (**self**) as argument and returns an option:
  - This function must return `None` if the content of the message is either **empty** or contains the word **stupid**. Otherwise it returns the content of the message.

You will have to create one more **function** that is not associated with any structure:

- `check_ms` which:
  - receives as parameters the reference to the structure `Message`
  - and returns a tuple, containing a `bool` and a `string`:
    - This function will execute the function `send_ms` and if the result of the option is `None`, it should return (false, "ERROR: illegal").Otherwise it returns `true` and the content of the message sent.

### Notions

- [Enum Definition](https://doc.rust-lang.org/stable/book/ch06-01-defining-an-enum.html?highlight=option#the-option-enum-and-its-advantages-over-null-values)

### Expected Function

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
