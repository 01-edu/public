## error types

### Instructions

For this exercise, you will have to implement an **error type**.

The main objective is to create a form validator, where you must implement an
error type. This must validate the password and the first name. The
first name must not be empty and the password must have **at least 8 char**, a **combination of alphanumeric**, **none-alphanumeric** (ex: <, & or /, ...)

ex: "asDd123=%" => good (as 8 char, alphanumeric and non-alphanumeric)\
"asgfD" => error (does only have alphabetic letters)\
"asdsdf2" => error (missing none-alphanumeric)\
"sad\_#$" => error (does not have a combination of alphanumeric)

Create a structure called `Form` that will have the following fields:

- `first_name`, that will be a string
- `last_name`, that will be a string
- `birth`, of type `NaiveDate` that will convert a string "2015-09-05" to a date of that format
- `fav_colour`, of type Color that must be a `enum` with the fields `Red`, `Blue`  and `Green`
- `birth_location`, that will be a string
- `password`, that will be a string

You must also implement for this structure a function to initialize the structure, `new` and a function called
`validate` that will validate the form

For the error type you must create a type struct called `FErr`, that will be the type for the error
It must have the fields:

- `form_values`, this will be a tuple of strings that will save the value that the user inserted into the form

ex: ("password", "asdaSD\_")
("first_name", "someone")

- `date`, that will have the date that the error occurred in the format "2020-12-14 09:33:41"
- `err`, that will have the error description:
  - "No user name"
  - "At least 8 characters"
  - "Combination of different ASCII character types (numbers, letters and none alphanumeric characters)"

### Notions

- [Error types](https://doc.rust-lang.org/rust-by-example/error/multiple_error_types/define_error_type.html)
- [Struct NaiveDate](https://docs.rs/chrono/0.4.19/chrono/naive/struct.NaiveDate.html)


### Dependencies

chrono = "0.4"


### Expected Function

```rust
pub use chrono::{Utc, NaiveDate};

// this will be the structure that wil handle the errors
#[derive(Debug, Eq, PartialEq)]
pub struct FErr {
    // expected public fields
}
impl FErr {
    pub fn new(name: String, error: String, err: String) -> FErr {}
}

#[derive(Debug, Eq, PartialEq)]
pub enum Color {
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
        fav_colour: Color,
        birth_location: String,
        password: String,
    ) -> Form {}
    
    pub fn validate(&self) -> Result<Vec<&str>, FErr> {}
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
        Color::Red,
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
Form { first_name: "Lee", last_name: "Silva", birth: 2015-09-05, fav_colour: Red, birth_location: "Africa", password: "qwqwsa1dty_" }
["Valid first name", "Valid password"]
FErr { form_values: ("first_name", ""), date: "2022-04-21 09:18:12", err: "No user name" }
FErr { form_values: ("password", "dty_1"), date: "2022-04-21 09:18:12", err: "At least 8 characters" }
FErr { form_values: ("password", "asdasASd(_"), date: "2022-04-21 09:18:12", err: "Combination of different ASCII character types (numbers, letters and none alphanumeric characters)" }
FErr { form_values: ("password", "asdasASd123SA"), date: "2022-04-21 09:18:12", err: "Combination of different ASCII character types (numbers, letters and none alphanumeric characters)" }
$
```
