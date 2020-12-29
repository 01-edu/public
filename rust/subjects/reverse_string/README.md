## reverse_string

### Instructions

Complete the function `rev_str` that takes a `&str` as a parameter, and returns a string with its words reversed.

### Expected functions

```rust
pub fn rev_str(input: &str) -> String {}
```

### Usage

Here is a program to test your function.

```rust
fn main() {
    println!("{}", rev_str("Hello, world!"));
    println!("{}", rev_str("Hello, my name is Roman"));
    println!("{}", rev_str("I have a nice car!"));
    println!("{}", rev_str("How old are You"));
    println!("{}", rev_str("ex: this is an example água"));
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
!dlrow ,olleH
namoR si eman ym ,olleH
!rac ecin a evah I
uoY era dlo woH
augá elpmaxe na si siht :xe
student@ubuntu:~/[[ROOT]]/test$
```
