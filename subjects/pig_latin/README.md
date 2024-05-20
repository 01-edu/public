## pig_latin

### Instructions

Create a **function** which transforms the string passed as an argument into Pig Latin:

- If a word begins with a vowel, just add "ay" to the end.
- If it begins with a consonant, then we take all consonants before the first vowel, move them to the end of the word, and then add "ay" at the end.
- If a word starts with a consonant followed by "qu", move it to the end of the word, and then add an "ay" at the end.
- Only the latin vowels will be considered as vowels (aeiou).

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
    println!("{}", pig_latin(&String::from("queen")));
}
```

And its output:

```console
$ cargo run
iglooay
appleay
ellohay
aresquay
enonxay
airchay
ueenqay
$
```

### Notions

- [patterns](https://doc.rust-lang.org/book/ch18-00-patterns.html)
