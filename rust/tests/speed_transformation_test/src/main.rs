// Create a function that receives the speed in km/h (kilometers per hour) and returns the
// equivalent in m/s (meters per second)

use speed_transformation::*;

fn main() {
	let km_h = 100.0;
	let m_s = km_per_hour_to_meters_per_second(km_h);
	println!("{} km/h is equivalent to {} m/s", km_h, m_s);
}

#[test]
fn kmh_to_ms() {
	assert_eq!(km_per_hour_to_meters_per_second(90.0), 25.0);
	assert_eq!(km_per_hour_to_meters_per_second(50.0), 13.88888888888889);
	assert_eq!(km_per_hour_to_meters_per_second(10.0), 2.7777777777777777);
	assert_eq!(km_per_hour_to_meters_per_second(100.0), 27.77777777777778);
}
