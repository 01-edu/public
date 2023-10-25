## scytale_cipher

### Instructions

Create a **function** which decode a scytale cipher (also known as spartan cipher).

In practice, it is represented by a strip wrapped around a cylinder. The message is written across the loops of the strip (not along the strip). The message becomes _coded_ if the radius of the cylinder changes, or the strip is removed from the cylinder.

Your function will receive a `String` representing the ciphered message, and a `u32` representing the number of letters by turn of the strip around the cylinder.

> If the ciphered message is empty or the letters per turn are 0 the function will return `None`.

### Example

**letters_per_turn 2:** `"scytale Code"` -> `"sec yCtoadle"`

```console
--------------------------------
  |s|  |c|  |y|  |t|  |a|  |l|
  |e|  | |  |C|  |o|  |d|  |e|
--------------------------------
```

**letters_per_turn 4:** `"scytale Code"` -> `"steoca dylCe"`

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
pub fn scytale_decoder(s: String, letters_per_turn: u32) -> Option<String> {
}
```

### Usage

Here is a program to test your function

```rust
use scytale_decoder::*;

fn main() {
    println!("\"sec yCtoadle\" size=2 -> {:?}",
        scytale_decoder("sec yCtoadle".to_string(), 2));

    println!("\"steoca dylCe\" size=4 -> {:?}",
        scytale_decoder("steoca dylCe".to_string(), 4));
}
```

And its output

```console
$ cargo run
"sec yCtoadle" size=2 -> Some("scytale Code")
"steoca dylCe" size=4 -> Some("scytale Code")
$
```
