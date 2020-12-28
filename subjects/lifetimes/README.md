## lifetimes

### Instructions

Create a struct called Person that has two fields: name of type string slice (&str) and age of type u8 and create the associated function new which creates a new person with age 0 and with the name given

### Expected Functions and Data Structures

```rust
#[derive(Debug)]
pub struct Person{
	pub name: &str,
	pub age: u8,
}

impl Person {
	pub fn new(name: &str) -> Person {
	}
}
```

### Usage

Here is a program to test your function.

```rust
fn main() {
	let person = Person::new("Leo");

	println!("Person = {:?}", person);
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
Person = Person { name: "Leo", age: 0 }
student@ubuntu:~/[[ROOT]]/test$
```
