## pig_latin

### Instructions

Write a **function** that transforms a string passed as argument in its `Pig Latin` version.

The rules used by Pig Latin are the following:

- If a word begins with a vowel, just add "ay" to the end.
- If it begins with a consonant, then we take all consonants before the first vowel and we put them at the end of the word and add "ay" at the end.
- If a word starts with a consonant followed by "qu", move it to the end of the word, and then add an "ay" at the end.

### Notions

- [patterns](https://doc.rust-lang.org/book/ch18-00-patterns.html)

### Expected functions

```rust
pub fn pig_latin(text: &str) -> String {

}
```

### Usage

Here is a program to test your function.

```rust
use pig_latin::*;

fn main() {
    println!("{}", pig_latin(&String::from("igloo")));
    println!("{}", pig_latin(&String::from("apple")));
    println!("{}", pig_latin(&String::from("hello")));
    println!("{}", pig_latin(&String::from("square")));
    println!("{}", pig_latin(&String::from("xenon")));
    println!("{}", pig_latin(&String::from("chair")));
}
```

And its output:

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
iglooay
appleay
ellohay
aresquay
enonxay
airchay
student@ubuntu:~/[[ROOT]]/test$
```
