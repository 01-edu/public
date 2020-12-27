// Write a functions called identity that calculates the identity of a
// value (receives any data type and returns the same value)
use generics::*;

fn main() {
	println!("{}", identity("Hello, world!"));
	println!("{}", identity(3));
}

#[derive(PartialEq, Debug)]
struct Point {
	x: i32,
	y: i32,
}

#[test]
fn test_with_int() {
	assert_eq!(identity(3), 3);
}

#[test]
fn test_with_float() {
	assert_eq!(identity(1.0), 1.0);
}

#[test]
fn test_with_str() {
	assert_eq!(identity("you"), "you");
}

#[test]
fn test_with_struct() {
	let s = Point { x: 1, y: 2 };
	assert_eq!(identity(&s), &s);
}
