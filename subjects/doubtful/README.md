## doubtful

### Instructions

Create a function named `doubtful` which appends a question mark to every string passed to it. It must not return a value.

### Expected functions
```rust
pub fn doubtful(s: /*give the correct type*/ ) {
}
```

You'll need to complete the function signature, so that it works properly with the usage example. You'll also need to complete the usage if you plan to use it.

### Usage

Here is a program to test your function

```rust
fn main() {
	let mut s = String::from("Hello");

	println!("Before changing the string: {}", s);

	doubtful(/*add your code here*/);

	println!("After changing the string: {}", s);
}
```

And its output

```console
$ cargo run
Before changing the string: Hello
After changing the string: Hello?
$
```
