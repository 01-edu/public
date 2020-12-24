use temperature_conv::*;

use std::f64::EPSILON;

fn eql(a: f64, b: f64) -> bool {
	(b - a).abs() < EPSILON
}

#[test]
fn test_f_to_c() {
	println!("{}", fahrenheit_to_celsius(83.0));
	assert!(eql(fahrenheit_to_celsius(20.0), -6.666666666666666));
	assert!(eql(fahrenheit_to_celsius(83.0), 28.333333333333332));
}

#[test]
fn test_c_to_f() {
	assert!(eql(celsius_to_fahrenheit(27.0), 80.6));
	assert!(eql(celsius_to_fahrenheit(0.0), 32.0))
}
