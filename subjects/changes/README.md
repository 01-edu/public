## changes

### Instructions

Imagine you are working on some software to control smart lights in a house. You have access to an array of all the lights in that house.

Define the associated **function** `new`, and add it to the data structure `Light`. It should create a new light with the alias passed as an argument, with a brightness of 0.

Define the **function** `change_brightness`, which receives a `Vec` of lights, an `alias` and a `u8`value. It should find the light in the `Vec` by its alias, and set the value of the brightness.

### Expected Functions and Structure

```rust
#[derive(Debug, Eq, PartialEq, Clone)]
pub struct Light {
	pub alias: String,
	pub brightness: u8,
}

impl Light {
	pub fn new(alias: &str) -> Self {
	}
}

pub fn change_brightness(lights: &mut Vec<Light>, alias: &str, value: u8) {
}
```

### Usage

Here is an incomplete program to test your function

```rust
use changes::*;

fn main() {
	// bedroom
	let mut lights = vec![
		Light::new("living_room"),
		Light::new("bedroom"),
		Light::new("rest_room"),
	];
	println!("brightness = {}", lights[0].brightness);
	change_brightness(&mut lights, "living_room", 200);
	println!("new brightness = {}", lights[0].brightness);
}
```

And its expected output

```console
$ cargo run
brightness = 0
new brightness = 200
$
```

### Notions

- [Example of Structs](https://doc.rust-lang.org/book/ch05-02-example-structs.html)
- [Keyword Self](https://doc.rust-lang.org/std/keyword.self.html)
