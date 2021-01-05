## changes

### Instructions

Imagine you are working in a software to control smart lights in a house. You have access to an array of all the lights in a house.

Define the associated function `new` to the data structure `Light` which creates a new light with the alias passed in the arguments and a brightness of 0.

Define the function `change_brightness` that receives a Vec of lights, an alias and a u8 value and sets the u8 value as the new brightness of the light identified by the alias in the Vec of lights.

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
student@ubuntu:~/[[ROOT]]/test$ cargo run
brightness = 0
new brightness = 200
student@ubuntu:~/[[ROOT]]/test$
```
