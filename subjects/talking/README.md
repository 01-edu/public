## talking

### Instructions

Build the function `talking` which will allow you to talk with your computer.

His answers will be created by you following the rules below.

- He answers "There is no need to yell, calm down!" if you yell at him, for example "LEAVE ME ALONE!"
  (it is considered yelling when the sentence is all written in capital letters).
- He answers "Sure." if you ask him something without yelling, for example "Is everything ok with you?"
- He answers "Quiet, I am thinking!" if you yell a question at him. "HOW ARE YOU?"
- He says "Just say something!" if you address him without actually saying anything.
- He answers "Interesting" to anything else.

### Notions

- [patterns](https://doc.rust-lang.org/book/ch18-00-patterns.html)

### Expected functions

```rust
pub fn talking(text: &str) -> &str {

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
