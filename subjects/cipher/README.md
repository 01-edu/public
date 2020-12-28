## cipher

### Instructions

The Atbash cipher is a encryption method in which each letter of a word is replaced with its mirror letter in the alphabet

Your objective is to create a function called `cipher` this must return a Result wrapped in an Option, this result should return either a boolean
or an Error being the structure `CipherError`. This structure should be the error type for the function `cipher`

This function should compare the original string wih the ciphered string. returning true if the cipher is correct otherwise the error type
CipherErr with the a true or false if it is validated and the proper atbash cipher.

### Expected Function

```rust

#[derive(Debug, Clone, Eq, PartialEq)]
pub struct CipherError {
    // expected public fields
}
impl CipherError {
    pub fn new(validation: bool, expected: String) -> CipherError {}
}
pub fn cipher(original: &str, ciphered: &str) -> Option<Result<bool, CipherError>> {}
```

### Usage

Here is a program to test your function

```rust
fn main() {
    println!("{:?}", cipher("1Hello 2world!", "1Svool 2dliow!"));
    println!("{:?}", cipher("1Hello 2world!", "svool"));
    println!("{:?}", cipher("", "svool"));
}
```

And its output:

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
Some(Ok(true))
Some(Err(CipherError { validation: false, expected: "1Svool 2dliow!" }))
None
student@ubuntu:~/[[ROOT]]/test$
```
