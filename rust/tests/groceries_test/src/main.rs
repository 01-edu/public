// # Instructions

// Create a function called `insert`
// fn insert(vec: &mut Vec<String>, val: String) that inserts a new element at the end of the Vec

#[allow(unused_imports)]
use groceries::*;

#[allow(dead_code)]
fn main() {
	let mut groceries = vec![
		"yogurt".to_string(),
		"panetone".to_string(),
		"bread".to_string(),
		"cheese".to_string(),
	];
	insert(&mut groceries, String::from("nuts"));
	println!("The groceries list now = {:?}", &groceries);
	println!(
		"The second element of the grocery  list is {:?}",
		at_index(&groceries, 1)
	);
}

#[test]
fn test_insertions() {
	let mut groceries = Vec::new();
	insert(&mut groceries, "milk".to_string());
	assert_eq!(groceries, ["milk"]);
	insert(&mut groceries, "bread".to_string());
	assert_eq!(groceries, ["milk", "bread"]);
}

#[test]
fn test_index() {
	let groceries: Vec<String> = vec![
		"milk".to_string(),
		"bread".to_string(),
		"water".to_string(),
		"wine".to_string(),
	];
	assert_eq!(at_index(&groceries, 0), "milk");
}
