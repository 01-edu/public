## talking

### Instructions

Build the function `talking` which will allow you to talk with your computer.

Its answers will be created by you following the rules below.

- It answers `"There is no need to yell, calm down!"` if you yell at it. For example `"LEAVE ME ALONE!"`. It considers yelling when all the letters are uppercase.
- It answers `"Sure."` if you ask it something without yelling. For example `"Is everything ok with you?"`.
- It answers `"Quiet, I am thinking!"` if you yell a question at it. For example: `"HOW ARE YOU?"`.
- It says `"Just say something!"` if you don't say anything.
- It answers `"Interesting"` to anything else.

### Expected functions

```rust
pub fn talking(text: &str) -> &str {
    todo!()
}
```

### Usage

Here is a program to test your function.

```rust
use talking::*;

fn main() {
    println!("{:?}", talking("JUST DO IT!"));
    println!("{:?}", talking("Hello how are you?"));
    println!("{:?}", talking("WHAT'S GOING ON?"));
    println!("{:?}", talking("something"));
    println!("{:?}", talking(""));
}
```

And its output:

```console
$ cargo run
"There is no need to yell, calm down!"
"Sure."
"Quiet, I am thinking!"
"Interesting"
"Just say something!"
$
```

### Notions

- [patterns](https://doc.rust-lang.org/book/ch18-00-patterns.html)
