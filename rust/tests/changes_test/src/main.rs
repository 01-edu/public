/*
## changes

### Instructions

Imagine you are working in a software to control smart lights in a
house. You have access to an array of all the lights in a house.

Define the associated function `new` to the data structure `Light`
which creates a new light with the alias passed in the arguments and
a brightness of 0.

Define the function `change_brightness` that receives a Vec of lights,
an alias and a u8 value and sets the u8 value as the new brightness of the light
identified by the alias in the Vec of lights.
*/

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

#[test]
fn test_unexistente_alias() {
	let mut lights = Vec::new();
	for i in 0..5 {
		let alias = format!("light-{}", i);
		lights.push(Light::new(&alias));
	}
	let copy = lights.clone();
	change_brightness(&mut lights, "light-6", 100);
	assert_eq!(copy, lights);
}

#[test]
fn test_alias() {
	let mut lights = Vec::new();
	for i in 0..5 {
		let alias = format!("light-{}", i);
		lights.push(Light::new(&alias));
	}
	let alias = "light-3";
	change_brightness(&mut lights, alias, 100);
	assert_eq!(lights[3].brightness, 100);
}
