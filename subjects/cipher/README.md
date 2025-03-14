## cipher

### Instructions

The Atbash cipher is an encryption method in which each letter of a word is replaced by its mirror letter in the alphabet.

Your objective is to create a **function** named `cipher` which must return a `Result`.

`cipher` should compare the original string with a ciphered string. It should return an empty value (`()`) if the cipher is correct. If the cipher is incorrect it should return the error type `CipherError` with the expected cipher (`expected: String`).

### Expected Function and structure

```rust
#[derive(Debug, PartialEq)]
pub struct CipherError {
    // expected public fields
}

pub fn cipher(original: &str, ciphered: &str) -> Result<(), CipherError> {
    todo!()
}
```

### Usage

Here is a program to test your function:

```rust
use cipher::*;

fn main() {
    println!("{:?}", cipher("1Hello 2world!", "1Svool 2dliow!"));
    println!("{:?}", cipher("1Hello 2world!", "svool"));
}
```

And its output:

```console
$ cargo run
Ok(())
Err(CipherError { expected: "1Svool 2dliow!" })
$
```
