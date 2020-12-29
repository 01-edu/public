// # Instructions:
// Create a function that borrows two slices (&[T]) and returns a hashmap where
// the first slice represents the keys and the second represents the values.

// The signature of the function is the following
// fn slices_to_map<'a, T: Hash + Eq, U>(keys: &'a [T], values: &'a [U]) -> HashMap<&'a T, &'a U> {

// # Example:
// for the slices &["hello", "how", "are", "you"] &[1, 3, 5, 8]
// returns the hashmap ["hello": 1, "how": 3, "are": 5, "you":8]

use slices_to_map::*;
#[allow(unused_imports)]
use std::collections::HashMap;

#[allow(dead_code)]
fn main() {
	let keys = ["Olivia", "Liam", "Emma", "Noah", "James"];
	let values = [1, 3, 23, 5, 2];
	println!("{:?}", slices_to_map(&keys, &values));
}

#[cfg(test)]
mod tests {
	use super::*;

	#[test]
	fn test_same_length() {
		let keys = ["Olivia", "Liam", "Emma", "Noah", "James"];
		let values = [1, 3, 23, 5, 2];
		let mut expected = HashMap::new();
		expected.insert(&"Olivia", &1);
		expected.insert(&"Liam", &3);
		expected.insert(&"Emma", &23);
		expected.insert(&"Noah", &5);
		expected.insert(&"James", &2);
		assert_eq!(slices_to_map(&keys, &values), expected);
	}

	#[test]
	fn test_different_length() {
		let keys = ["Olivia", "Liam", "Emma", "Noah", "James"];
		let values = [1, 3, 23, 5, 2, 9];
		let mut expected = HashMap::new();
		expected.insert(&"Olivia", &1);
		expected.insert(&"Liam", &3);
		expected.insert(&"Emma", &23);
		expected.insert(&"Noah", &5);
		expected.insert(&"James", &2);
		assert_eq!(slices_to_map(&keys, &values), expected);

		let keys = ["Olivia", "Liam", "Emma", "Noah", "James", "Isabella"];
		let values = [1, 3, 23, 5, 2];
		assert_eq!(slices_to_map(&keys, &values), expected);
	}

	#[test]
	fn it_works_for_vecs() {
		let mut expected = HashMap::new();
		expected.insert(&"Olivia", &1);
		expected.insert(&"Liam", &3);
		expected.insert(&"Emma", &23);
		expected.insert(&"Noah", &5);
		expected.insert(&"James", &2);

		let keys = ["Olivia", "Liam", "Emma", "Noah", "James"];
		let values = [1, 3, 23, 5, 2];

		assert_eq!(slices_to_map(&keys, &values), expected);
		let keys = vec!["Olivia", "Liam", "Emma", "Noah", "James"];
		let values = vec![1, 3, 23, 5, 2, 9];
		assert_eq!(slices_to_map(&keys, &values), expected);

		let keys = vec!["Olivia", "Liam", "Emma", "Noah", "James", "Isabella"];
		let values = vec![1, 3, 23, 5, 2];
		assert_eq!(slices_to_map(&keys, &values), expected);
	}
}
