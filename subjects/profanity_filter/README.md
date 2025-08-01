## profanity filter

### Instructions

Generally it is more desirable to catch the failure of some parts of a program instead of panicking.

For this exercise you will have to create a message blocker, where you must block the word `"stupid"`.

You will need to create a **function** named `check_ms` which accepts a string and returns a `Result`. It should return an error with the message `"ERROR: illegal"` if the message is either empty or contains the word stupid. Otherwise, it should return the content of the message.

##### Expected Function

```rust
pub fn check_ms(message: &str) -> Result<&str, &str> {
    todo!()
}
```

### Usage

Here is a program to test your function

```rust
fn main() {
    ["hello there", "", "you are stupid", "stupid"]
        .into_iter()
        .for_each(|m| println!("{:?}", profanity_filter::check_ms(m)));
}
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

- [Result Definition](https://doc.rust-lang.org/stable/book/ch09-02-recoverable-errors-with-result.html)
