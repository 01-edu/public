// Define a tuple to represent a student
// Each student is identified by an id number of type i32
// Its name and its last name as a string
// Print the content of the tuple to stdout

use tuples::*;

fn main() {
	let student = Student(20, "Pedro".to_string(), "Domingos".to_string());
	println!("Student: {:?}", student);
	// After print only the first name
	println!("Student first name: {}", first_name(&student));
	// After print only the last name
	println!("Student last name: {}", last_name(&student));
	// After print only the id
	println!("Student Id: {}", id(&student));
}

#[test]
fn test_id() {
	let student = Student(20, String::from("Pedro"), String::from("Domingos"));
	assert_eq!(id(&student), 20);
}

#[test]
fn test_first_name() {
	let student = Student(20, String::from("Pedro"), String::from("Domingos"));
	assert_eq!(first_name(&student), "Pedro".to_string());
}

#[test]
fn test_last_name() {
	let student = Student(20, String::from("Pedro"), String::from("Domingos"));
	assert_eq!(last_name(&student), "Domingos".to_string());
}
