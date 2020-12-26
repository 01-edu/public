## changes

### Instructions

Define the function `add_excitement` that add the exclamation point at the end of every sentenced that is passed to the function without returning anything but just modifying the input.

### Expected Functions

```rust
fn add_excitement(s: &) {
}
```

### Usage

Here is an incomplete program to test your function

```rust
fn main() {
	let mut a = String::from("Hello");

	println!("the value of `a` before changing {}", a);

	add_excitement(a);

	println!("The value of `a` after changing {}", a);
}
```

And its expected output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
the value of `a` before changing Hello
The value of `a` after changing Hello!
student@ubuntu:~/[[ROOT]]/test$
```
