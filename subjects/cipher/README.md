## cipher

### Instructions

The Atbash cipher is an encryption method in which each letter of a word is replaced by its mirror letter in the alphabet.

Your objective is to create a **function** called `cipher` which must return a `Result` wrapped in an `Option`, this result should return either a `boolean`
or an `Error` based on the structure `CipherError`. This structure should be the error type for the **function** `cipher`.

This function should compare the original `String` with the ciphered `String`. It should return `true` if the cipher is correct. If the cipher is incorrect it should return the error type `CipherErr` with a `boolean` and the expected atbash cipher `String`.

### Notions

- [Module std::fmt](https://doc.rust-lang.org/std/fmt/index.html)

### Expected Function and structure

```rust

#[derive(Debug, Clone, Eq, PartialEq)]
pub struct CipherError {
    // expected public fields
}
impl CipherError {
    pub fn new(validation: bool, expected: String) -> CipherError {

    }
}
pub fn cipher(original: &str, ciphered: &str) -> Option<Result<bool, CipherError>> {

}
```

### Usage

Here is a program to test your function:

```rust
use cipher::*;

fn main() {
    println!("{:?}", cipher("1Hello 2world!", "1Svool 2dliow!"));
    println!("{:?}", cipher("1Hello 2world!", "svool"));
    println!("{:?}", cipher("", "svool"));
}
```

And its output:

```console
$ cargo run
Some(Ok(true))
Some(Err(CipherError { validation: false, expected: "1Svool 2dliow!" }))
None
$
```
