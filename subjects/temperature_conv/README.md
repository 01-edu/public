## temperature_conv

### Instructions

Write two functions to transform values of temperatures from celcius to fahrenheit and the other way arraound:

### Expected funcions

```rust
fn fahrenheit_to_celsius(f: f64) -> f64 {
}

fn celsius_to_fahrenheit(c: f64) -> f64 {
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
