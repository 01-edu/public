// In this exercise you will create a model of and gives an API to
// deal with blood types

// Start creating the data representation of the blood types
// Create the enumerator `Antigen` that has 4 possibilities: A, B, O and AB
// And the enumerator `RhFactor` that has two possible values: Positive
// and Negative

// After, create the struct BloodType that contains two fields with the
// names antigen and rh_factor

// To provide a simple way to create blood types implement the trait
// FromStr for BloodType (which will allow us to use the `parse`
// method and the associated function from_str, so we can do:
// ```rust
//    let a_neg: BloodType = "A-".parse();
// ```
//)

// Implement the std::cmp::Ord trait to make possible to sort a vector
// or array of BloodType's

// Implement the trait std::Debug for BloodType allowing to print a
// vector such as [BloodType { antigen: A, rh_factor: Positive},
// BloodType{ antigen: B, rh_factor: Negative}] as [ A+, A-] using the
// formatting {:?}

// Write three methods for BloodType:
// - can_receive_from(&self, other: BloodType) -> bool {}: which
// returns true if self can receive blood from `other` blood type
// - donors(&self) -> Vec<BloodType>: which returns
// all the blood types that can give blood to self
// - recipients(&self) -> Vec<BloodType>: which returns all the blood
// types that can receive blood from self

#[allow(unused_imports)]
use blood_types::{Antigen, BloodType, RhFactor};

fn main() {
	let blood_type: BloodType = "O+".parse().unwrap();
	println!("recipients of O+ {:?}", blood_type.recipients());
	println!("donors of O+ {:?}", blood_type.donors());
	let another_blood_type: BloodType = "A-".parse().unwrap();
	println!(
		"donors of O+ can receive from {:?} {:?}",
		&another_blood_type,
		blood_type.can_receive_from(&another_blood_type)
	);
}

#[test]
fn compatible_ab_neg_with_a_pos() {
	let blood_type: BloodType = "AB-".parse().unwrap();
	let other_bt: BloodType = "A+".parse().unwrap();
	assert!(!blood_type.can_receive_from(&other_bt));
}

#[test]
fn compatible_a_neg_with_a_pos() {
	let blood_type: BloodType = "A-".parse().unwrap();
	let other_bt: BloodType = "A+".parse().unwrap();
	assert!(!blood_type.can_receive_from(&other_bt));
}

#[test]
fn compatible_a_neg_with_ab_neg() {
	let blood_type: BloodType = "AB-".parse().unwrap();
	let other_bt: BloodType = "A-".parse().unwrap();
	assert!(blood_type.can_receive_from(&other_bt));
}

#[test]
fn compatible_ab_neg_with_o_pos() {
	let blood_type: BloodType = "AB-".parse().unwrap();
	let other_bt: BloodType = "O+".parse().unwrap();
	assert!(!blood_type.can_receive_from(&other_bt));
}

#[test]
fn compatible_ab_pos_with_o_pos() {
	let blood_type: BloodType = "AB+".parse().unwrap();
	let other_bt: BloodType = "O+".parse().unwrap();
	assert!(blood_type.can_receive_from(&other_bt));
}

#[test]
fn test_compatible_ab_neg_with_o_neg() {
	let blood_type: BloodType = "AB-".parse().unwrap();
	let other_bt: BloodType = "O-".parse().unwrap();
	assert!(blood_type.can_receive_from(&other_bt));
}

#[test]
fn test_antigen_ab_from_str() {
	let blood = "AB+";
	let blood_type: BloodType = blood.parse().unwrap();
	assert_eq!(blood_type.antigen, Antigen::AB);
	assert_eq!(blood_type.rh_factor, RhFactor::Positive);
}

#[test]
fn test_antigen_a_from_str() {
	let blood = "A-";
	let blood_type = blood.parse::<BloodType>().unwrap();
	assert_eq!(blood_type.antigen, Antigen::A);
	assert_eq!(blood_type.rh_factor, RhFactor::Negative);
}

#[test]
#[should_panic]
fn test_unexistent_blood_type() {
	let _blood_type: BloodType = "AO-".parse().unwrap();
}

#[test]
fn test_donors() {
	let mut givers = "AB+".parse::<BloodType>().unwrap().donors();
	println!("Before sorting {:?}", &givers);
	givers.sort();
	println!("{:?}", &givers);
	let mut expected = vec![
		"AB-".parse::<BloodType>().unwrap(),
		"A-".parse().unwrap(),
		"B-".parse().unwrap(),
		"O-".parse().unwrap(),
		"AB+".parse().unwrap(),
		"A+".parse().unwrap(),
		"B+".parse().unwrap(),
		"O+".parse().unwrap(),
	];
	expected.sort();
	assert_eq!(givers, expected);
}

#[test]
fn test_a_neg_donors() {
	let mut givers = "A-".parse::<BloodType>().unwrap().donors();
	givers.sort();
	let mut expected: Vec<BloodType> = vec!["A-".parse().unwrap(), "O-".parse().unwrap()];
	expected.sort();
	assert_eq!(givers, expected);
}

#[test]
fn test_o_neg_donors() {
	let mut givers = "O-".parse::<BloodType>().unwrap().donors();
	givers.sort();
	let mut expected: Vec<BloodType> = vec!["O-".parse().unwrap()];
	expected.sort();
	assert_eq!(givers, expected);
}

#[test]
fn test_ab_pos_recipients() {
	let mut recipients: Vec<BloodType> = "AB+".parse::<BloodType>().unwrap().recipients();
	recipients.sort();
	let mut expected: Vec<BloodType> = vec!["AB+".parse().unwrap()];
	expected.sort();
	assert_eq!(recipients, expected);
}

#[test]
fn test_a_neg_recipients() {
	let mut recipients = "A-".parse::<BloodType>().unwrap().recipients();
	recipients.sort();
	let mut expected: Vec<BloodType> = vec![
		"A-".parse().unwrap(),
		"AB+".parse().unwrap(),
		"A+".parse().unwrap(),
		"AB-".parse().unwrap(),
	];
	expected.sort();
	assert_eq!(recipients, expected);
}

#[test]
fn test_output() {
	let blood_type: BloodType = "O+".parse().unwrap();
	println!("recipients of O+ {:?}", blood_type.recipients());
	println!("donors of O+ {:?}", blood_type.donors());
	let another_blood_type: BloodType = "A-".parse().unwrap();
	println!(
		"donors of O+ can receive from {:?} {:?}",
		&another_blood_type,
		blood_type.can_receive_from(&another_blood_type)
	);
}
