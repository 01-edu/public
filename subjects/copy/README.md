## copy

### Instructions

Your objective is to fix the program so that all functions work.

- `nbr_function`, returns a tuple with the original value, the exponential and logarithm of that value
- `str_function`, returns a tuple with the original value and the exponential of that value as a string
- `vec_function`, returns a tuple with the original value and the logarithm of each value

The objective is to now how ownership works with different types.

### Notions

- https://doc.rust-lang.org/rust-by-example/scope/move.html

### Expected Functions

```rust
pub fn nbr_function(c: u32) -> (u32, f64, f64) {
}

pub fn str_function(a: String) -> (String, String) {
}

pub fn vec_function(b: Vec<u32>) -> (Vec<u32>, Vec<f64>) {
}
```

### Usage

Here is a possible program to test your function :

```rust
fn main() {
	let a = String::from("1 2 4 5 6");
	let b = vec![1, 2, 4, 5];
	let c: u32 = 0;

	println!("{:?}", nbr_function(c));
	// output: (12, 162754.79141900392, 2.4849066497880004)
	println!("{:?}", vec_function(b));
	// output: ([1, 2, 4], [0.0, 0.6931471805599453, 1.3862943611198906])
	println!("{:?}", str_function(a));
	// output: ("1 2 4", "2.718281828459045 7.38905609893065 54.598150033144236")
}
```

And its output:

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
(0, 1.0, inf)
([1, 2, 4, 5], [0.0, 0.6931471805599453, 1.3862943611198906, 1.6094379124341003])
("1 2 4 5 6", "2.718281828459045 7.38905609893065 54.598150033144236 148.4131591025766 403.4287934927351")
student@ubuntu:~/[[ROOT]]/test$
```
