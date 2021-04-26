## temperature_conv

### Instructions

Write two functions which convert values of temperatures from `fahrenheit` to `celcius` and the other way around.

To pass this exercise you must use (9/5) in **both** functions.

### Expected functions

```rust
pub fn fahrenheit_to_celsius(f: f64) -> f64 {
}

pub fn celsius_to_fahrenheit(c: f64) -> f64 {
}
```

### Usage

```rust
use temperature_conv::*;

fn main() {
	println!("{} F = {} C", -459.67, fahrenheit_to_celsius(-459.67));
	println!("{} C = {} F", 0.0, celsius_to_fahrenheit(0.0));
}
```

And its output:

```console
student@ubuntu:~/temperature_conv/test$ cargo run
-459.67 F = -273.15 C
0 C = 32 F
student@ubuntu:~/temperature_conv/test$
```

