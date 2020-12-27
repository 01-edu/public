// Write a function that receives a string slice and returns the
// length of character of the string

use strings::*;

fn main() {
	println!("lenght of {} = {}", "❤", "❤".len());
	println!("lenght of {} = {}", "❤", char_lenght("❤"));
	println!("lenght of {} = {}", "形声字", char_lenght("形聲字"));
	println!("lenght of {} = {}", "形声字", "形聲字".len());
	println!("lenght of {} = {}", "change", "change".len());
	println!("lenght of {} = {}", "change", char_lenght("change"));
	println!("char lenght of {} = {}", "😍", char_lenght("😍"));
}

// fn char_lenght(s: &str) -> usize {
// 	let mut chars = 0;
// 	for _ in s.chars() {
// 		chars += 1;
// 	}
// 	chars
// }

#[test]
fn test_ascii() {
	let s = "ascii";
	assert_eq!(char_lenght(s), 5);
}

#[test]
fn test_emoji() {
	let s = "❤😍";
	assert_eq!(char_lenght(s), 2);
}
#[test]
fn test_chinese_char() {
	let s = "形声字";
	assert_eq!(char_lenght(s), 3);
}
