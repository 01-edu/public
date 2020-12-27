// Complete the `capitalize_first` function that turns the first letter of a string uppercase.
// Complete the `title_case` function that turns the first letter of each word in a string uppercase.
// Complete the `change_case` function that turns the uppercase letters of a string into lowercase and
// the lowercase letters into uppercase.

use capitalizing::*;

#[allow(dead_code)]
fn main() {
	println!("{}", capitalize_first("joe is missing"));
	println!("{}", title_case("jill is leaving A"));
	println!("{}", change_case("heLLo THere"));
}

#[test]
fn test_success() {
	assert_eq!(capitalize_first("hello"), "Hello");
	assert_eq!(capitalize_first("this is working"), "This is working");
}

#[test]
fn test_titlle_case() {
	assert_eq!(title_case("this is a tittle"), "This Is A Tittle");
	assert_eq!(title_case("hello my name is carl"), "Hello My Name Is Carl");
}

#[test]
fn test_change_case() {
	assert_eq!(change_case("PROgraming"), "proGRAMING");
	assert_eq!(change_case("heLLo THere"), "HEllO thERE");
}

#[test]
fn test_empty() {
	assert_eq!(capitalize_first(""), "");
	assert_eq!(title_case(""), "");
	assert_eq!(change_case(""), "");
}
