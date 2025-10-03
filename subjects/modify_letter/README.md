## modify_letter

### Instructions

Create a **function** `remove_letter_sensitive` that returns a string without the letter specified as argument.

Create a **function** `remove_letter_insensitive` that returns a string without the letter specified as argument (ignoring case).

Create a **function** `swap_letter_case` that returns a string swapping the case for the chosen letter.

> Our tests will only use ASCII characters.

### Expected Functions

```rust
pub fn remove_letter_sensitive(s: &str, letter: char) -> String {
    todo!()
}

pub fn remove_letter_insensitive(s: &str, letter: char) -> String {
    todo!()
}

pub fn swap_letter_case(s: &str, letter: char) -> String {
    todo!()
}
```

### Usage

Here is a program to test your functions.

```rust
use modify_letter::*;

fn main() {
    println!("{}", remove_letter_sensitive("Jojhn jis sljeepjjing", 'j'));
    println!(
        "{}",
        remove_letter_insensitive("JaimA ais swiaAmmingA", 'A')
    );
    println!("{}", swap_letter_case("byE bye", 'e'));
}
```

And its output

```console
$ cargo run
John is sleeping
Jim is swimming
bye byE
$
```
