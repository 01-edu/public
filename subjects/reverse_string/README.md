## reverse_string

### Instructions

Write a **function** `rev_str` that takes a `&str` as a parameter, and returns a `String` with its letters reversed.

### Notions

- [Strings](https://doc.rust-lang.org/rust-by-example/std/str.html)
- [Primitive Type str](https://doc.rust-lang.org/std/primitive.str.html)
- [Primtive Type char](https://doc.rust-lang.org/std/primitive.char.html)
- [Module std::string](https://doc.rust-lang.org/std/string/index.html)

### Expected Functions

```rust
pub fn rev_str(input: &str) -> String {
}
```

### Usage

Here is a possible program to test your function :

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

And its output:

```console
student@ubuntu:~/reverse_string/test$ cargo run
!dlrow ,olleH
namoR si eman ym ,olleH
!rac ecin a evah I
uoY era dlo woH
augá elpmaxe na si siht :xe
student@ubuntu:~/reverse_string/test$
```
