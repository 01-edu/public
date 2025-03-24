## expected_variable

### Instructions

Create a **function** named `expected_variable` that receives a string to compare and an expected string. It should return an `Option`. Every comparison should be case insensitive.

If the compared string is **not** in camel case or snake case, `expected_variable` returns `None`. You can use the `case` crate to figure that out. Otherwise, the compared string should be compared to the expected string using the `edit_distance` **function** that you have already created.

If the result of `edit_distance` has more than 50% alikeness to the expected string, considering the length of the expected string and the result of `edit_distance`, the function should return that value with a `'%'` symbol after the number.
Otherwise `expected_value` should return `None`.

### Expected Function

> You'll need to work out the function signature for yourself.

### Usage

Here is a program to test your function:

```rs
use expected_variable::*;

fn main() {
    println!(
        "{} close to it",
        expected_variable("On_Point", "on_point").unwrap()
    );
    println!(
        "{} close to it",
        expected_variable("soClose", "so_close").unwrap()
    );
    println!(
        "{:?}",
        expected_variable("something", "something_completely_different")
    );
    println!(
        "{} close to it",
        expected_variable("BenedictCumberbatch", "BeneficialCucumbersnatch").unwrap()
    );
}
```

And its output:

```sh
$ cargo run
100% close to it
88% close to it
None
67% close to it
$
```

### Notions

- [Crate convert_case](https://crates.io/crates/convert_case)
- [Crate heck](https://crates.io/crates/heck)
- [Specifying Dependencies](https://doc.rust-lang.org/cargo/reference/specifying-dependencies.html)
