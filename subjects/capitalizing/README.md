## capitalizing

### Instructions

Complete the `capitalize_first` **function** which converts the first letter of the string to uppercase.

Complete the `title_case` **function** which converts the first letter of each word in a string to uppercase.

Complete the `change_case` **function** which converts the uppercase letters of a string into lowercase, and the lowercase to uppercase.

### Expected Functions

```rust
pub fn capitalize_first(input: &str) -> String {
}

pub fn title_case(input: &str) -> String {
}

pub fn change_case(input: &str) -> String {
}
```

### Usage

Here is a program to test your functions.

```rust
use capitalizing::*;

fn main() {
    println!("{}", capitalize_first("joe is missing"));
    println!("{}", title_case("jill is leaving A"));
    println!("{}", change_case("heLLo THere"));
}
```

And its output

```console
$ cargo run
Joe is missing
Jill Is Leaving A
HEllO thERE
$
```
