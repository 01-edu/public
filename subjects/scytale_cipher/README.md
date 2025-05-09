## scytale_cipher

### Instructions

Create a **function** which creates a scytale cipher (also known as spartan cipher).

In practice, it is represented by a strip wrapped around a cylinder. The message is written across the loops of the strip (not along the strip). The message becomes _coded_ if the radius of the cylinder changes, or the strip is removed from the cylinder.

Your function should recreate the scytale cipher, so that the `String` represents the message, and the `u32` represents the number of times the strip is wrapped around the cylinder.

### Example

**size 6:** `"scytale Code"` -> `"sec yCtoadle"`

```console
--------------------------------
  |s|  |c|  |y|  |t|  |a|  |l|
  |e|  | |  |C|  |o|  |d|  |e|
--------------------------------
```

**size 8:** `"scytale Code"` -> `"sCcoydtea l e"`

```console
------------------------------------------
  |s|  |c|  |y|  |t|  |a|  |l|  |e|  | |
  |C|  |o|  |d|  |e|  | |  | |  | |  | |
------------------------------------------
```

### Expected Functions

```rust
pub fn scytale_cipher(message: String, i: u32) -> String {
    todo!()
}
```

### Usage

Here is a program to test your function

```rust
use scytale_cipher::scytale_cipher;

fn main() {
    println!("\"scytale Code\" size=6 -> {:?}", scytale_cipher(String::from("scytale Code"), 6)));
    println!("\"scytale Code\" size=8 -> {:?}", scytale_cipher(String::from("scytale Code"), 8)));
}
```

And its output

```console
$ cargo run
"scytale Code" size=6 -> "sec yCtoadle"
"scytale Code" size=8 -> "sCcoydtea l e"
$
```
