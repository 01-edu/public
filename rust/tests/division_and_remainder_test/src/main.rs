use division_and_remainder::*;

fn main() {
	let x = 9;
	let y = 4;
	let (division, remainder) = divide(x, y);
	println!(
		"\t{}/{}: division = {}, remainder = {}",
		x, y, division, remainder
	);
}

#[test]
#[should_panic]
fn divide_by_0() {
	divide(40, 0);
}

#[test]
fn test_divide() {
	let (div, rem) = divide(40, 3);

	assert_eq!(div, 13);
	assert_eq!(rem, 1);

	let (div, rem) = divide(389, 39);

	assert_eq!(div, 9);
	assert_eq!(rem, 38);

	let (div, rem) = divide(29, 332);

	assert_eq!(div, 0);
	assert_eq!(rem, 29);
}
