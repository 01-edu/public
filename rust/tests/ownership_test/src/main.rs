/*
## ownership

### Instruction

Create a function that takes ownership of a string and returns the
first sub-word in it. It should work for camelCase as well as snake_case
first_subword(camelCase) returns camel
first_subword(snake_case) returns snake

And fix the printing expression so the code works

###
*/

use ownership::*;

fn main() {
	let s1 = String::from("helloWorld");
	let s2 = String::from("snake_case");
	let s3 = String::from("CamelCase");
	let s4 = String::from("just");

	println!("first_subword({}) = {}", s1.clone(), first_subword(s1)); // Must print first_subword(helloWorld) = hello
	println!("first_subword({}) = {}", s2.clone(), first_subword(s2)); // Must print first_subword(snake_case) = snake
	println!("first_subword({}) = {}", s3.clone(), first_subword(s3)); // Must print first_subword(CamelCase) = Camel
	println!("first_subword({}) = {}", s4.clone(), first_subword(s4)); // Must print first_subword(just) = just
}

#[cfg(test)]
mod tests {
	use super::*;

	#[test]
	fn first_subword_test() {
		struct TstString<'a> {
			str: String,
			l: &'a str,
		}

		let o_tsts = vec![
			TstString {
				str: "helloWorld".to_string(),
				l: "hello",
			},
			TstString {
				str: "how_you".to_string(),
				l: "how",
			},
			TstString {
				str: "Changeyou".to_string(),
				l: "Changeyou",
			},
			TstString {
				str: "CamelCase".to_string(),
				l: "Camel",
			},
		];

		for t in o_tsts.iter() {
			assert_eq!(t.l.to_string(), first_subword(t.str.clone()));
		}
	}
}
