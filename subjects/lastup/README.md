## lastup

### Instructions

Complete the `lastup` function that takes a string and puts the last letter of each word in uppercase and the rest in lowercase.

### Expected Functions

```rust
pub fn lastup(input: &str) -> String {
}
```

### Usage

Here is a program to test your function.

```rust
use lastup::lastup;

fn main() {
    println!("{}", lastup("joe is missing"));
    println!("{}", lastup("jill is leaving A"));
    println!("{}",lastup("heLLo THere!"));
}
```

And its output

```console
$ cargo run
joE iS missinG
jilL iS leavinG A
hellO therE!
$
```
