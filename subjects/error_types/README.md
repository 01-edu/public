## error types

### Instructions

For this exercise, you will have to implement an **error type** for a form validator. We will validate the password and the first name.

The first name must not be empty and the password must have **at least 8 characters**, and a combination of ASCII **alphanumeric characters** and **symbols** (`<`, `&`, `/` ...).

Examples:

- `"asDd123=%"`: **good**.
- `"asgfD"`: **error** as it only contains alphabetic characters.
- `"asdsdf2"`: **error** as it is missing symbols.
- `"sad\_#$"`: **error** as it is missing numeric characters.

Create a structure named `Form` that will have the following fields:

- `name`: `String`
- `password`: `String`

You must implement the **associated functions** `new` and `validate` that will validate the form.

For the error type you must create a `struct` named `FormError`. It must have the fields:

- `form_values`: represents the invalid input in a key-value pair. For example: `("password", "asdaSD\_")` or `("first_name", "someone")`
- `date`: represents when the error occurred in the format `"YYYY-MM-DD HH:MM:SS"`
- `err`: the error description, which should be either of the following values:
  - `"Username is empty"`
  - `"Password should be at least 8 characters long"`
  - `"Password should be a combination of ASCII numbers, letters and symbols"`

### Expected Function

```rust
// this will be the structure that wil handle the errors
#[derive(Debug, Eq, PartialEq)]
pub struct FormError {
    // expected public fields
}

impl FormError {
    pub fn new(field_name: &'static str, field_value: String, err: &'static str) -> Self {
        todo!()
    }
}

#[derive(Debug, Eq, PartialEq)]
pub struct Form {
    // expected public fields
}

impl Form {
    pub fn validate(&self) -> Result<(), FormError> {
        todo!()
    }
}
```

### Usage

Here is a program to test your function:

```rust
use error_types::*;

fn main() {
    let mut form_output = Form {
        name: "Lee".to_owned(),
        password: "qwqwsa1dty_".to_owned(),
    };

    println!("{:?}", form_output);
    println!("{:?}", form_output.validate());

    form_output.name = "".to_owned();
    println!("{:?}", form_output.validate());

    form_output.name = "as".to_owned();
    form_output.password = "dty_1".to_owned();
    println!("{:?}", form_output.validate());

    form_output.password = "asdasASd(_".to_owned();
    println!("{:?}", form_output.validate());

    form_output.password = "asdasASd123SA".to_owned();
    println!("{:?}", form_output.validate());
}
```

And its output:

```console
$ cargo run
Form { name: "Lee", password: "qwqwsa1dty_" }
Ok(())
Err(FormError { form_values: ("name", ""), date: "2022-10-17 12:09:25", err: "Username is empty" })
Err(FormError { form_values: ("password", "dty_1"), date: "2022-10-17 12:09:25", err: "Password should be at least 8 characters long" })
Err(FormError { form_values: ("password", "asdasASd(_"), date: "2022-10-17 12:09:25", err: "Password should be a combination of ASCII numbers, letters and symbols" })
Err(FormError { form_values: ("password", "asdasASd123SA"), date: "2022-10-17 12:09:25", err: "Password should be a combination of ASCII numbers, letters and symbols" })
$
```

### Notions

- [Error types](https://doc.rust-lang.org/rust-by-example/error/multiple_error_types/define_error_type.html)
- [chrono](https://docs.rs/chrono/latest/chrono/index.html)
