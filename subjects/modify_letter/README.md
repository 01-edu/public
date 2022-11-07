## modify_letter

### Instructions

Create a **function** `remove_letter_sensitive` that returns a string without the letter specified as argument

Create a **function** `remove_letter_insensitive` that returns a string without the letter specified as argument (ignoring case)

Create a **function** `swap_letter_case` that returns a string swapping the case for the chosen letter.

If there is an empty character in the arguments, the function should return the string as being passed.

If there is an empty string, it should return an empty string.

### Expected Functions

```rust
pub fn remove_letter_sensitive(s: &str, letter: char) -> String {
}

pub fn remove_letter_insensitive(s: &str, letter: char) -> String {
}

pub fn swap_letter_case(s: &str, letter: char) -> String {
}
```

### Usage

Here is a program to test your functions.

```rust
use modify_letter::*modify_letter*;

fn main() {
    println!("{}", remove_letter_sensitive("Joje jis mijssjing", 'j'));
    println!("{}", remove_letter_insensitive("JaillA ais swiaAmmingA", 'A'));
    println!("{}", swap_letter_case("hEllo therE", 'e'));
}
```

And its output

```console
$ cargo run
Joe is missing
Jill is swimming
hello thEre
$
```
