use fibonacci2::*;

fn main() {
	println!(
		"The element in the position {} in fibonacci series is {}",
		2,
		fibonacci(2)
	);
	println!(
		"The element in the position {} in fibonacci series is {}",
		4,
		fibonacci(4)
	);
	println!(
		"The element in the position {} in fibonacci series is {}",
		22,
		fibonacci(22)
	);
	println!(
		"The element in the position {} in fibonacci series is {}",
		20,
		fibonacci(20)
	);
}

#[test]
fn it_works() {
	assert_eq!(fibonacci(0), 0);
	assert_eq!(fibonacci(1), 1);
	assert_eq!(fibonacci(22), 17711);
	assert_eq!(fibonacci(20), 6765);
}
