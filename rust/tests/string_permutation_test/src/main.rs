// Define the function `is_permutation` that returns true if the
// string `s1` is a permutation of `s2`, otherwise it returns false
// `s1` is a permutation of `s2` if all the elements in `s1` appear the
// same number of times in `s2` and all the characters in `s1` appear in
// `s2` even if in different order)
use string_permutation::*;

#[allow(dead_code)]
fn main() {
	let word = "thought";
	let word1 = "thougth";
	println!(
		"Is `{}` a permutation of `{}`? = {}",
		word,
		word1,
		is_permutation(word, word1)
	);
}

#[test]
fn permutation_ascii() {
	assert!(is_permutation("abcde", "edbca"));
	assert!(!is_permutation("avcde", "edbca"));
	assert!(!is_permutation("cde", "edbca"));
	assert!(is_permutation("code", "deco"));
	assert!(!is_permutation("code", "deeco"));
	assert!(!is_permutation("codde", "deeco"));
}

#[test]
fn permutation_unicode() {
	assert!(is_permutation("hello♥", "♥oelhl"));
	assert!(!is_permutation("♥", "♥♥"));
}
