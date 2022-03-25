## expected_variable

### Instructions

Create a **function** `expected_variable` that receives two strings: one to be evaluated and the other to be compared to (expected) and returns an Option. Every comparison should be case insensitive.

If the evaluated string is not in camel case or in snake case (use the `case` crate for this evaluation) `expected_variable` returns `None`.
Otherwise the evaluated string should be compared to the expected string using the `edit_distance` **function** that you did in one of the previous quests.

If the result of `edit_distance` has more than 50% of 'alikeness' to the expected string, considering the length of the expected string and the result of `edit_distance`, the function should return that same percentage with a '%' symbol in front of the number.
Otherwise `expected_value` should return `None`.

### Notions

- [Crate case](https://crates.io/crates/case)
- [Specifying Dependencies](https://doc.rust-lang.org/cargo/reference/specifying-dependencies.html)

### Expected Function

#### For this exercise the signature of the function has to be found out.

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
        expected_variable("soClose", "So_Close").unwrap()
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
