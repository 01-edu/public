## ownership

### Instruction

Create a **function** named `first_subword`, that takes ownership of a string and returns the first sub-word in it. It should work for `camelCase`, `PascalCase`, and `snake_case`.

### Expected functions
```rust
pub fn first_subword(mut s: String) -> String {
}
```

### Usage

Here is a program to test your function:

```rust
use ownership::first_subword;

fn main() {
	let s1 = String::from("helloWorld");
	let s2 = String::from("snake_case");
	let s3 = String::from("CamelCase");
	let s4 = String::from("just");

	println!("first_subword({}) = {}", s1.clone(), first_subword(s1));
	println!("first_subword({}) = {}", s2.clone(), first_subword(s2));
	println!("first_subword({}) = {}", s3.clone(), first_subword(s3));
	println!("first_subword({}) = {}", s4.clone(), first_subword(s4));
}
```

And its output:

```console
$ cargo run
first_subword(helloWorld) = hello
first_subword(snake_case) = snake
first_subword(CamelCase) = Camel
first_subword(just) = just
$
```

### Notions

- [ownership](https://doc.rust-lang.org/book/ch04-00-understanding-ownership.html)
