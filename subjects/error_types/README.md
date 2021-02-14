## error types

### Instructions

For this exercise, you will have to implement an **error type**.

The main objective is to create a form validator, where you must implement an
error type. This must validate the password and the first name. The
first name must not be empty and the password must have at least 8 char and a combination of alphanumeric and none alphanumeric ASCII characters

ex: "asDd123=%" => good
"asgfD" => error
"asdsdf2" => error
"sad\_#$" => error

Create a structure called `Form` that will have the following fields:

- `first_name`, that will be a string
- `last_name`, that will be a string
- `birth`, of type `NaiveDate` that will convert a string "2015-09-05" to a date of that format
- `sex`, SexType that must be a `enum` with the fields `Male` and `Female`
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
pub enum SexType {
    // expected public fields
}
#[derive(Debug, Eq, PartialEq)]
pub struct Form {
    // expected public fields
}

impl Form {
    pub fn new(first_name: String,
           last_name: String,
           birth: NaiveDate,
           sex: SexType,
           birth_location: String,
           password: String) -> Form {}
    pub fn validate(&self) -> Result<Vec<&str>, FErr> {}
}
```

### Usage

Here is a program to test your function:

```rust
use error_types::*;

fn main() {
    let mut form_output = Form::new(
        String::from("Alice"),
        String::from("Bear"),
        create_date("2015-09-05"),
        SexType::Male,
        String::from("Africa"),
        String::from("qwqwsa1dty_"));

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
student@ubuntu:~/[[ROOT]]/test$ cargo run
Form { first_name: "Lee", last_name: "Silva", birth: 2015-09-05, sex: Male, birth_location: "Africa", password: "qwqwsa1dty_" }
["Valid first name", "Valid password"]
FErr { form_values: ("first_name", ""), date: "2020-12-28 13:29:11", err: "No user name" }
FErr { form_values: ("password", "dty_1"), date: "2020-12-28 13:29:11", err: "At least 8 characters" }
FErr { form_values: ("password", "asdasASd(_"), date: "2020-12-28 13:29:11", err: "Combination of different ASCII character types (numbers, letters and none alphanumeric characters)" }
FErr { form_values: ("password", "asdasASd123SA"), date: "2020-12-28 13:29:11", err: "Combination of different ASCII character types (numbers, letters and none alphanumeric characters)" }
student@ubuntu:~/[[ROOT]]/test$
```
