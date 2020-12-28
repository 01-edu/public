// # Instructions
// Implement the From<u32> Trait to create a roman number from a u32
// the roman number should be in subtractive notation (the common way to write roman
// number I, II, II, IV, V, VI, VII, VIII, IX, X ...)

// For this start by defining the digits as `RomanDigit` with the values
// I, V, X, L, C, D, M and Nulla for 0

// Next define RomanNumber as a wrapper to a vector of RomanDigit's
// And implement the Trait From<u32>

// Examples:
// RomanNumber::from(32) = [X,X,X,I,I]
// RomanNumber::from(9) = [I,X]
// RomanNumber::from(45) = [X,L,V]
// RomanNumber:;from(0) = [Nulla]

#[allow(unused_imports)]
use roman_numbers::RomanDigit::*;
#[allow(unused_imports)]
use roman_numbers::RomanNumber;

#[allow(dead_code)]
fn main() {
	println!("{:?}", RomanNumber::from(32));
	println!("{:?}", RomanNumber::from(9));
	println!("{:?}", RomanNumber::from(45));
	println!("{:?}", RomanNumber::from(0));
}
#[test]
fn it_works() {
	assert_eq!(RomanNumber::from(3).0, [I, I, I]);
	assert_eq!(RomanNumber::from(6).0, [V, I]);
	assert_eq!(RomanNumber::from(15).0, [X, V]);
	assert_eq!(RomanNumber::from(30).0, [X, X, X]);
	assert_eq!(RomanNumber::from(150).0, [C, L]);
	assert_eq!(RomanNumber::from(200).0, [C, C]);
	assert_eq!(RomanNumber::from(600).0, [D, C]);
	assert_eq!(RomanNumber::from(1500).0, [M, D]);
}

#[test]
fn substractive_notation() {
	assert_eq!(RomanNumber::from(4).0, [I, V]);
	assert_eq!(RomanNumber::from(44).0, [X, L, I, V]);
	assert_eq!(RomanNumber::from(3446).0, [M, M, M, C, D, X, L, V, I]);
	assert_eq!(RomanNumber::from(9).0, [I, X]);
	assert_eq!(RomanNumber::from(94).0, [X, C, I, V]);
}
