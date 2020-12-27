// Create the function `contain` that checks a `HashMap` to see if it contains the given key.
// Create the function `remove` that removes a given key from the `HashMap`.

use std::collections::HashMap;

use simple_hash::*;

fn main() {
	let mut hash: HashMap<&str, i32> = HashMap::new();
	hash.insert("Daniel", 122);
	hash.insert("Ashley", 333);
	hash.insert("Katie", 334);
	hash.insert("Robert", 14);

	println!(
		"Does the HashMap contains the name Roman? => {}",
		contain(hash.clone(), "Roman")
	);
	println!(
		"Does the HashMap contains the name Katie? => {}",
		contain(hash.clone(), "Katie")
	);
	println!("Removing Robert {:?}", remove(hash.clone(), "Robert"));
	println!("Hash {:?}", hash);
}

#[test]
fn test_contains() {
	let mut s = HashMap::new();

	s.insert("Pedro", 43);
	s.insert("Ralph", 12);
	s.insert("Johnny", 546);
	s.insert("Albert", 12323214);

	assert_eq!(true, contain(s.clone(), "Pedro"));
	assert_eq!(true, contain(s.clone(), "Ralph"));
	assert_eq!(true, contain(s.clone(), "Johnny"));
	assert_eq!(true, contain(s.clone(), "Albert"));
	assert_eq!(false, contain(s.clone(), "Marco"));
	assert_eq!(false, contain(s.clone(), "Joan"));
	assert_eq!(false, contain(s.clone(), "Louise"));
}

#[test]
fn test_remove() {
	let mut n = HashMap::new();
	n.insert("Dani Sordo", 37);
	n.insert("Sébastien Loeb", 46);
	n.insert("Ott Tanak", 32);
	n.insert("Thierry Neuville", 32);

	remove(n.clone(), "Dani Sordo");
	assert_eq!(true, contain(n.clone(), "Ott Tanak"));
	assert_eq!(false, contain(n.clone(), "Dani Ŝordo"))
}
