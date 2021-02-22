use temperature_conv::*;

use std::f64::EPSILON;

fn eql(a: f64, b: f64) -> bool {
	(b - a).abs() < EPSILON
}

#[test]
fn test_f_to_c() {
	let temp_f = 20.0;
	println!("{}°F = {}°C", temp_f, fahrenheit_to_celsius(temp_f));
	assert!(eql(fahrenheit_to_celsius(temp_f), -6.666666666666666));
	let temp_f = 83.0;
	println!("{}°F = {}°C", temp_f, fahrenheit_to_celsius(temp_f));
	assert!(eql(fahrenheit_to_celsius(temp_f), 28.333333333333332));
}

#[test]
fn test_c_to_f() {
	let temp_c = 27.0;
	println!("{}°C = {}°F", temp_c, fahrenheit_to_celsius(temp_c));
	assert!(eql(celsius_to_fahrenheit(27.0), 80.6));
	let temp_c = 0.0;
	println!("{}°C = {}°F", temp_c, fahrenheit_to_celsius(temp_c));
	assert!(eql(celsius_to_fahrenheit(temp_c), 32.0))
}
