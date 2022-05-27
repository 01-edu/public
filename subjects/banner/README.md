## banner

### Instructions

"`Result` is a better version of the `Option` type that describes a possible `error` instead of possible `absence`".

Create a structure called `Flag` which has the following elements:

- `short_hand`: `String`
- `long_hand`: `String`
- `desc`: `String`

This structure must have a **function** called `opt_flag` which initializes the structure.
This **function** receives two string references and returns a structure `Flag`. Here is an example of its usage:

```rust
    let d = Flag::opt_flag("diff", "gives the difference between two numbers");

    println!("short hand: {}, long hand: {}, description: {}", d.short_hand, d.long_hand, d.desc);
    // output: "short hand: -d, long hand: --diff, description: gives the difference between two numbers"
```

A second structure named `FlagsHandler` will be given which just has one element: `flags: HashMap<(String, String), Callback>`. You'll need to implement the following **associated functions**" (methods) associated with `FlagsHandler` are for you to complete:"

- `add_flag`, which adds the flag and callback function to the HashMap.
- `exec_func`, which executes the function using the flag provided and returns the result. The result will be either a string with the value from the callback or an error.

A `type` called `Callback` will also be provided. It is a function which is going to be used in the structure and functions above. This function will be the callback for the flag associated to it.

You will have to create the following callback functions:

- `div`: which converts the reference strings to `float`s and returns the `Result`, as the division of the `float`s or the standard (std) error: `ParseFloatError`.
- `rem`: which converts the reference strings to `float`s and returns the `Result`, as the remainder of the division of the `float`s or the standard (std) error `ParseFloatError`.

### Expected Function

```rust
use std::collections::HashMap;

pub struct Flag {
    // expected public fields
}

impl Flag {
    pub fn opt_flag(l_h: &str, d: &str) -> Flag {

    }
}

pub type Callback = fn(&str, &str) -> Result<String, ParseFloatError>;

pub struct FlagsHandler {
    pub flags: HashMap<(String, String), Callback>,
}

impl FlagsHandler {
    pub fn add_flag(&mut self, flag: (String, String), func: Callback) {

    }
    pub fn exec_func(&mut self, flag: (String, String), argv: &[&str]) -> String {

    }
}

pub fn div(a: &str, b: &str) -> Result<String, ParseFloatError> {

}
pub fn rem(a: &str, b: &str) -> Result<String, ParseFloatError> {

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

    handler.add_flag((d.short_hand, d.long_hand), div);
    handler.add_flag((r.short_hand, r.long_hand), rem);

    println!("{:?}", handler.exec_func(("-d".to_string(), "--division".to_string()), &["1.0", "2.0"]));

    println!("{:?}",handler.exec_func(("-r".to_string(), "--remainder".to_string()), &["2.0", "2.0"]));

    println!("{:?}",handler.exec_func(("-d".to_string(), "--division".to_string()), &["a", "2.0"]));

    println!("{:?}",handler.exec_func(("-r".to_string(), "--remainder".to_string()), &["2.0", "fd"]));
}
```

And its output:

```console
$ cargo run
"0.5"
"0"
"invalid float literal"
"invalid float literal"
$
```

### Notions

- [Result](https://doc.rust-lang.org/rust-by-example/error/result.html)
- [Method optflag](https://docs.rs/getopts/0.2.18/getopts/struct.Options.html#method.optflag)
