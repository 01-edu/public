## error types

### Instructions

For this exercise, you will have to implement an **error type** for a form validator. This must validate the password and the first name.

The first name must not be empty and the password must have **at least 8 characters**, and a combination of **alphabetic**, **numeric** and **none-alphanumeric** (`<`, `&`,  `/` ...).
Additionally, you need to create a helper function create_date that converts a string like "2015-09-05" into a NaiveDate format.

Examples:
- `"asDd123=%"`: **good**.
- `"asgfD"`: **error** as it only contains alphabetic characters.
- `"asdsdf2"`: **error** as it is missing none-alphanumeric characters.
- `"sad\_#$"`: **error** as it is missing numeric characters.

Create a structure named `Form` that will have the following fields:
- `first_name`: `String`
- `last_name`: `String`
- `birth`: `NaiveDate` (converted from a string using the create_date function).
- `birth_location`: `String`
- `password`: `String`

You must implement the **associated functions** `new` and
`validate` that will validate the form.

For the error type you must create a `struct` named `FormError`. It must have the fields:
- `form_values`: this will be a tuple of strings representing the invalid input. For example: `("password", "asdaSD\_")` or `("first_name", "someone")`

- `date`: that will have the date that the error occurred in the format `"2020-12-14 09:33:41"`
- `err`: the error description:
  - `"No user name"`
  - `"At least 8 characters"`
  - `"Combination of different ASCII character types (numbers, letters and none alphanumeric characters)"`


### Dependencies

chrono = "0.4"


### Expected Function

```rust
pub use chrono::{Utc, NaiveDate};

// this will be the structure that wil handle the errors
#[derive(Debug, Eq, PartialEq)]
pub struct FormError {
    // expected public fields
}
impl FormError {
    pub fn new(field_name: String, field_value: String, err: String) -> FormError {}
}

#[derive(Debug, Eq, PartialEq)]
pub struct Form {
    // expected public fields
}

impl Form {
    pub fn new(
        first_name: String,
        last_name: String,
        birth: NaiveDate,
        birth_location: String,
        password: String,
    ) -> Form {}
    
    pub fn validate(&self) -> Result<Vec<&str>, FormError> {}
}
```

### Usage

Here is a program to test your function:

```rust
use error_types::*;

fn main() {
    let mut form_output = Form::new(
        String::from("Lee"),
        String::from("Silva"),
        create_date("2015-09-05"),
        String::from("Africa"),
        String::from("qwqwsa1dty_"),
    );

    println!("{:?}", form_output);
    println!("{:?}", form_output.validate().unwrap());

    form_output.first_name = String::from("");
    println!("{:?}", form_output.validate().unwrap_err());

    form_output.first_name = String::from("as");
    form_output.password = String::from("dty_1");
    println!("{:?}", form_output.validate().unwrap_err());

    form_output.password = String::from("asdasASd(_");
    println!("{:?}", form_output.validate().unwrap_err());

    form_output.password = String::from("asdasASd123SA");
    println!("{:?}", form_output.validate().unwrap_err());
}
```

And its output:

```console
$ cargo run
Form { first_name: "Lee", last_name: "Silva", birth: 2015-09-05, birth_location: "Africa", password: "qwqwsa1dty_" }
["Valid first name", "Valid password"]
FormError { form_values: ("first_name", ""), date: "2022-10-17 12:09:25", err: "No user name" }
FormError { form_values: ("password", "dty_1"), date: "2022-10-17 12:09:25", err: "At least 8 characters" }
FormError { form_values: ("password", "asdasASd(_"), date: "2022-10-17 12:09:25", err: "Combination of different ASCII character types (numbers, letters and none alphanumeric characters)" }
FormError { form_values: ("password", "asdasASd123SA"), date: "2022-10-17 12:09:25", err: "Combination of different ASCII character types (numbers, letters and none alphanumeric characters)" }
$
```

### Notions

- [Error types](https://doc.rust-lang.org/rust-by-example/error/multiple_error_types/define_error_type.html)
- [Struct NaiveDate](https://docs.rs/chrono/0.4.19/chrono/naive/struct.NaiveDate.html)
