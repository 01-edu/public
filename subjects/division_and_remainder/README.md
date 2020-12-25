## division_and_remainder

### Instructions

Create a function divide that receives two i32 and returns a tuple in which the first element is the result of the integer division between the two numbers and the second is the remainder of the division.

### Expected Function

```rust
fn divide(x: i32, y: i32) -> (i32, i32) {

}
```

### Usage

Here is a program to test you're function

```rust
fn main() {
	let x = 9;
	let y = 4;
	let (division, remainder) = divide(x, y);
	println!(
		"\t{}/{}: division = {}, remainder = {}",
		x, y, division, remainder
	);
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
9/4: division = 2, remainder = 1
student@ubuntu:~/[[ROOT]]/test$
```
