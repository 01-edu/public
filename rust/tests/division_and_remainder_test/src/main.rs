//Change the body of the function divide to return the value of the
// integer division of the x and y and the remainder of that division
// You're only allowed to change the body of the function
fn main() {
	let x = 9;
	let y = 4;
	let (division, remainder) = divide(x, y);
	println!(
		"\t{}/{}: division = {}, remainder = {}",
		x, y, division, remainder
	);
}

#[cfg(test)]
mod tests {
	use super::*;

	#[test]
	fn it_works() {
		assert_eq!(divide(3, 4), (3 / 4, 3 % 4));
		assert_eq!(divide(73, 4), (73 / 4, 73 % 4));
		assert_eq!(divide(432, 32), (432 / 32, 432 % 32));
		assert_eq!(divide(897, 29), (897 / 29, 897 % 29));
	}
}
