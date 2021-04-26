## speed_transformation

### Instructions

Create a **function** that receives the speed in km/h (kilometers per hour) and returns the equivalent in m/s (meters per second).

### Expected Function

```rust
pub fn km_per_hour_to_meters_per_second(km_h: f64) -> f64 {
}
```

### Usage

```rust
use speed_transformation::km_per_hour_to_meters_per_second;

fn main() {
	let km_h = 100.0;
	let m_s = km_per_hour_to_meters_per_second(km_h);
	println!("{} km/h is equivalent to {} m/s", km_h, m_s);
}
```

And its output:

```console
student@ubuntu:~/speed_transformation/test$ cargo run
100 km/h is equivalent to 27.77777777777778 m/s
student@ubuntu:~/speed_transformation/test$
```
