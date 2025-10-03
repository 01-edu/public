## scytale_decoder

### Instructions

Create a **function** which **decodes** a scytale cipher (also known as spartan cipher).

Your function will receive a `String` representing the ciphered message, and a `usize` representing the number of letters per turn of the strip around the cylinder.

> If the ciphered message is empty or the letters per turn are 0 the function will return `None`.

### Example

**letters_per_turn 2:** `"sec yCtoadle"` -> `"scytale Code"`

```console
--------------------------------
  |s|  |c|  |y|  |t|  |a|  |l|
  |e|  | |  |C|  |o|  |d|  |e|
--------------------------------
```

**letters_per_turn 4:** `"steoca dylCe"` -> `"scytale Code"`

```console
------------------------------------------
  |s|  |c|  |y|
  |t|  |a|  |l|
  |e|  | |  |C|
  |o|  |d|  |e|
------------------------------------------
```

### Expected Functions

```rust
pub fn scytale_decoder(s: String, letters_per_turn: usize) -> Option<String> {
    todo!()
}
```

### Usage

Here is a program to test your function:

```rust
use scytale_decoder::*;

fn main() {
    println!(
        "\"sec yCtoadle\" size=2 -> {:?}",
        scytale_decoder("sec yCtoadle".to_owned(), 2)
    );

    println!(
        "\"steoca dylCe\" size=4 -> {:?}",
        scytale_decoder("steoca dylCe".to_owned(), 4)
    );
}
```

And its output

```console
$ cargo run
"sec yCtoadle" size=2 -> Some("scytale Code")
"steoca dylCe" size=4 -> Some("scytale Code")
$
```
