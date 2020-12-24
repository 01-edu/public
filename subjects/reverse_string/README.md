## reverse_string

### Instructions

Write a function `rev_str` that takes a `&str` as a parameter, and returns a string with its words reversed.

### Expected Functions

```rust
pub fn rev_str(input: &str) -> String {
}
```

### Usage

```rust
use reverse_string::rev_str;

fn main() {
	println!("{}", rev_str("Hello, world!"));
	println!("{}", rev_str("Hello, my name is Roman"));
	println!("{}", rev_str("I have a nice car!"));
	println!("{}", rev_str("How old are You"));
	println!("{}", rev_str("ex: this is an example água"));
}
```

And it's output:

```console
!dlrow ,olleH
namoR si eman ym ,olleH
!rac ecin a evah I
uoY era dlo woH
augá elpmaxe na si siht :xe
```
