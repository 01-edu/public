## borrow

### Instructions

Complete the signature and the body of the `str_len` that receives a string or a string literal and returns its length without taking ownership of the value (i.e, borrowing the value)

### Expected Function

```rust
pub fn str_len(s: ) -> usize {
}
```

### Usage

Here is a possible program to test your function :

```rust
fn main() {
	let s = "hello";
	let s1 = "camelCase".to_string();

	println!("\tstr_len(\"{}\") = {}", s, str_len(s));
	println!("\tstr_len(\"{}\") = {}", s1, str_len(&s1));
}
```

And its output:

```rust
student@ubuntu:~/[[ROOT]]/test$ cargo run
str_len("hello") = 5
str_len("camelCase") = 9
student@ubuntu:~/[[ROOT]]/test$
```
