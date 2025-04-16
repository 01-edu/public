## scytale_cipher

### Instructions

Create a **function** which creates a scytale cipher (also known as spartan cipher).

In practice, it is represented by a strip wrapped around a cylinder. The message is written across the loops of the strip (not along the strip). The message becomes *coded* if the radius of the cylinder changes, or the strip is removed from the cylinder.

Your function should recreate the scytale cipher, so that the `&str` represents the message, and the `usize` represents the number of times the strip is wrapped around the cylinder.

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
fn scytale_cipher(message: &str, i: usize) -> String {
    todo!()
}
```

### Usage

Here is a program to test your function

```rust
fn main() {
    println!("\"scytale Code\" size=6 -> {:?}", scytale_cipher("scytale Code", 6)));
    println!("\"scytale Code\" size=8 -> {:?}", scytale_cipher("scytale Code", 8)));
}
```

And its output

```console
$ cargo run
"scytale Code" size=6 -> "sec yCtoadle"
"scytale Code" size=8 -> "sCcoydtea l e"
$
```
