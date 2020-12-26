/*
## error types

### Instructions

Make this code compile

*/

use changes::*;

fn main() {
	let mut a = String::from("Hello");

	println!("the value of `a` before changing {}", a);

	add_excitement(&mut a);

	println!("The value of `a` after changing {}", a);
}

// fn add_excitement(s: &mut String) {
// 	s.push_str("!");
// }

#[test]
fn test_ascii() {
	let mut expected = "hello".to_string();
	add_excitement(&mut expected);
	assert_eq!("hello!", &expected);
	let mut expected = "go on".to_string();
	add_excitement(&mut expected);
	assert_eq!("go on!", &expected);
}

#[test]
fn test_unicode() {
	let mut expected = "↕".to_string();
	add_excitement(&mut expected);
	assert_eq!(expected, "↕!");
}
