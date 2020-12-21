use scalar::*;

#[test]
#[should_panic]
fn test_panic_sum() {
	sum(25, 255);
}

#[test]
#[should_panic]
fn test_panic_diff() {
	diff(-32768, 32766);
}

#[test]
#[should_panic]
fn test_panic_pro() {
	pro(-128, 2);
}

#[test]
fn pass() {
	assert_eq!(sum(1, 2), 3);
	assert_eq!(diff(1, 2), -1);
	assert_eq!(pro(1, 2), 2);
	assert_eq!(quo(1.0, 2.0), 0.5);
	assert_eq!(rem(1.0, 2.0), 1.0);
}
