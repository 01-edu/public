## banner

### Instructions

"`Result` is a version of the `Option` type that describes a possible `Err` instead of `None`".

Create a structure called `Flag` which has the following elements:

- `short_hand: String`
- `long_hand: String`
- `desc: String`

This structure must have an associated **function** called `opt_flag` which initializes the structure.
This **function** receives two string references and returns a structure `Flag`. Here is an example of its usage:

```rust
let d = Flag::opt_flag("diff", "gives the difference between two numbers");

println!("short hand: {}, long hand: {}, description: {}", d.short_hand, d.long_hand, d.desc);
// output: "short hand: -d, long hand: --diff, description: gives the difference between two numbers"
```

An associated **type** called `Callback` will also be provided. It should represent a function pointer which is going to be used in the structure and functions below. This function will represent the callback for the flag associated to it.

A second structure named `FlagsHandler` will be given which just has one element: `flags: HashMap<(String, String), Callback>`. You'll also need to implement the following associated **functions**:

- `add_flag`, which adds the flag and callback function to the HashMap.
- `exec_func`, which executes the function using the flag provided and returns the result. Return either the result from the callback or the error stringified.

You will have to create the following callback functions:
- `div`: which converts the reference strings to `f64`s and returns the `Result`, as the division of these floats or the error `ParseFloatError`.
- `rem`: which converts the reference strings to `f64`s and returns the `Result`, as the remainder of the division of these floats or the error `ParseFloatError`.

### Expected Function

```rust
use std::{collections::HashMap, num::ParseFloatError};

pub struct Flag {
    // expected public fields
}

impl<'a> Flag<'a> {
    pub fn opt_flag(name: &'a str, d: &'a str) -> Self {
        todo!()
    }
}

pub type Callback = fn(&str, &str) -> Result<String, ParseFloatError>;

pub struct FlagsHandler {
    pub flags: HashMap<String, Callback>,
}

impl FlagsHandler {
    pub fn add_flag(&mut self, flag: Flag, func: Callback) {
        todo!()
    }

    pub fn exec_func(&self, input: &str, argv: &[&str]) -> Result<String, String> {
        todo!()
    }
}

pub fn div(a: &str, b: &str) -> Result<String, ParseFloatError> {
    todo!()
}

pub fn rem(a: &str, b: &str) -> Result<String, ParseFloatError> {
    todo!()
}

```

### Usage

Here is a program to test your function:

```rust
use banner::*;
use std::collections::HashMap;

fn main() {
    let mut handler = FlagsHandler { flags: HashMap::new() };

    let d = Flag::opt_flag("division", "divides the values, formula (a / b)");
    let r = Flag::opt_flag(
        "remainder",
        "remainder of the division between two values, formula (a % b)",
    );

    handler.add_flag(d, div);
    handler.add_flag(r, rem);

    println!("{:?}", handler.exec_func("-d", &["1.0", "2.0"]));

    println!("{:?}", handler.exec_func("-r", &["2.0", "2.0"]));

    println!("{:?}", handler.exec_func("--division", &["a", "2.0"]));

    println!("{:?}", handler.exec_func("--remainder", &["2.0", "fd"]));
}
```

And its output:

```console
$ cargo run
Ok("0.5")
Ok("0")
Err("invalid float literal")
Err("invalid float literal")
$
```

### Notions

- [Result](https://doc.rust-lang.org/rust-by-example/error/result.html)
- [Method optflag](https://docs.rs/getopts/0.2.18/getopts/struct.Options.html#method.optflag)
