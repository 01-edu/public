/*
## doubtful

### Instructions

Write a function called `doubtful` that adds to every string passed
to it a question mark (?)

You have to fix the code to make it compile an for that you can
only modify the code where is indicated
*/

use doubtful::*;

fn main() {
	let mut s = String::from("Hello");

	println!("Before changing the string: {}", s);
	doubtful(&mut s);
	println!("After changing the string: {}", s);
}

#[test]
fn test_function() {
	let mut s = "hello".to_string();
	let s_copy = s.clone();

	doubtful(&mut s);

	assert_eq!(s, s_copy + "?");
}
