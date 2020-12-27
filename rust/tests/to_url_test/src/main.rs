// Define a function called `to_url` that takes a string and
// substitutes every white-space with '%20'

use to_url::*;

fn main() {
	let s = "Hello, world!";
	println!("{} to be use as an url is {}", s, to_url(s));
}

#[cfg(test)]
mod test {
	use super::*;

	#[test]
	fn test_url() {
		assert_eq!(to_url("this is my search"), "this%20is%20my%20search");
	}
}
