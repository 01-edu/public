## cipher

### Instructions

The Atbash cipher is an encryption method in which each letter of a word is replaced by its mirror letter in the alphabet.

Your objective is to create a **function** named `cipher` which must return a `Result` wrapped in an `Option`. The `Result` should contain either a `boolean` or an `Error` based on the `CipherError` structure. This structure should be the error type for the **function** `cipher`.

`cipher` should compare the original `String` with the ciphered `String`. It should return `true` if the cipher is correct. If the cipher is incorrect it should return the error type `CipherError` with a `boolean` and the expected atbash cipher `String`.

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
