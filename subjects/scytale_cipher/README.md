## scytale_cipher

### Instructions

Create a function called `scytale_cipher` that takes a `string` and an `integer` and returns a `string`.
This function should create a scytale cipher also known as spartan cipher. In practice its a cylinder with a
strip strapped around it on which is written a message, when removed the strip the message is coded.
Depending on the size of the cylinder the message would change. So the only way to decipher the coded message is
by using a cylinder of the same size.

You function should recreate the scytale cipher, the string being the message and the size being the number of
straps in the cylinder.

### Example

message : "scytale Code"
size : 6

```console
['s', 'e']
['c', ' ']
['y', 'C']
['t', 'o']
['a', 'd']
['l', 'e']
```

output : sec yCtoadle
size : 8

```console
['s', 'C']
['c', 'o']
['y', 'd']
['t', 'e']
['a', ' ']
['l', ' ']
['e', ' ']
[' ', ' ']
```

output : sCcoydtea l e

### Expected Functions

```rust
fn scytale_cipher(message: String, i: u32) -> String {}
```

### Usage

Here is a program to test your function

```rust
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
