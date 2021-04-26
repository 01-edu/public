## lifetimes

### Instructions

Declare the struct called `Person` that has two fields:

- name of type string slice (&str)
- age of type u8

Additionaly, create the associated **function** `new` which creates a new person with age 0 and with the name given.

The expected Fucntions and Structures need to be completed.

### Notions

- [lifetimes](https://doc.rust-lang.org/book/ch10-03-lifetime-syntax.html)

### Expected Functions and Data Structures (Both need to be completed)

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
use lifetimes::*;

fn main() {
	let person = Person::new("Leo");

	println!("Person = {:?}", person);
}
```

And its output:

```console
student@ubuntu:~/lifetimes/test$ cargo run
Person = Person { name: "Leo", age: 0 }
student@ubuntu:~/lifetimes/test$
```
