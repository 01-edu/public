## doubtful

### Instructions

Write a function called `doubtful` that adds to every string passed to it a question mark (?)

You have to fix the code to make it compile an for that you can only modify the code where is indicated

### Expected

```rust
pub fn doubtful(s: &mut String) {
}
```

### Usage

Here is a program to test your function

```rust
fn main() {
	let mut s = String::from("Hello");

	println!("Before changing the string: {}", s);
	doubtful(&mut s);
	println!("After changing the string: {}", s);
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
Before changing the string: Hello
After changing the string: Hello?
student@ubuntu:~/[[ROOT]]/test$
```
