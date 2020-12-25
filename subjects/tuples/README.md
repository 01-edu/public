## tuples

### Instructions

- Define a tuple structure to represent a student.

- Each student is identified by an id number of type i32, his/her name and his/her last name as a string Print the content of the tuple to stdout

- Then define three functions to return the id, first name and last name.

### Expected Functions

```rust
pub fn id(student: &Student) -> i32 {
}

pub fn first_name(student: &Student) -> String {
}

pub fn last_name(student: &Student) -> String {
}
```

### Usage

Here is a program to test you're functions

```rust
use tuples::*;

fn main() {
	let student = Student(20, "Pedro".to_string(), "Domingos".to_string());
	println!("Student: {:?}", student);
	println!("Student first name: {}", first_name(&student));
	println!("Student last name: {}", last_name(&student));
	println!("Student Id: {}", id(&student));
}
```

And its output:

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
Student: Student(20, "Pedro", "Domingos")
Student first name: Pedro
Student last name: Domingos
Student Id: 20
student@ubuntu:~/[[ROOT]]/test$
```
