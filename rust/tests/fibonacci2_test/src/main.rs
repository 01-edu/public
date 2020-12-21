use fibonacci::*;

#[test]
fn it_works() {
	assert_eq!(fibonacci(0), 0);
	assert_eq!(fibonacci(1), 1);
	assert_eq!(fibonacci(22), 17711);
	assert_eq!(fibonacci(20), 6765);
}
