// Create a struct called Person that has two fields: name of type
// string slice (&str) and age of type u8
// and create the associated function new which creates a new person
// with age 0 and with the name given

use lifetimes::Person;

fn main() {
	let person = Person::new("Leo");

	println!("Person = {:?}", person);
}

#[cfg(test)]
mod tests {
	use super::*;

	#[test]
	fn fields() {
		let person = Person {
			name: "Dijkstra",
			age: 10,
		};
		assert_eq!(person.age, 10);
		assert_eq!(person.name, "Dijkstra");
	}

	#[test]
	fn create_person() {
		let person = Person::new("Leo");
		assert_eq!(person.age, 0);
		assert_eq!(person.name, "Leo");
	}
}
