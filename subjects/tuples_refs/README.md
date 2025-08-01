## tuples_refs

### Instructions

- Define a tuple `struct` to represent a `Student`. Each is identified by an id of type `u32`, their first name and last name.

- Then define three **functions** to return the id, first name and last name.

```rust
pub fn id(student: &Student) -> u32 {
}

pub fn first_name(student: &Student) -> &str {
}

pub fn last_name(student: &Student) -> &str {
}
```

### Usage

Here is a program to test your functions

```rust
use tuples_refs::*;

fn main() {
    let student = Student(20, "Pedro".to_string(), "Domingos".to_string());
    println!("Student's first name: {}", first_name(&student));
    println!("Student's last name: {}", last_name(&student));
    println!("Student's id: {}", id(&student));
}
```

And its output:

```console
$ cargo run
Student's first name: Pedro
Student's last name: Domingos
Student's id: 20
$
```

### Notions

- [Defining a struct](https://doc.rust-lang.org/stable/book/ch05-01-defining-structs.html)

- [The Tuple Type](https://doc.rust-lang.org/stable/book/ch03-02-data-types.html#the-tuple-type)

- [Tuples](https://doc.rust-lang.org/rust-by-example/primitives/tuples.html)

- [Tuple Structs without Named Fields](https://doc.rust-lang.org/stable/book/ch05-01-defining-structs.html#using-tuple-structs-without-named-fields-to-create-different-types)

- [Adding Useful Functionality with Derived Traits](https://doc.rust-lang.org/stable/book/ch05-02-example-structs.html?#adding-useful-functionality-with-derived-traits)

- [Chapter 7](https://doc.rust-lang.org/stable/book/ch07-03-paths-for-referring-to-an-item-in-the-module-tree.html)
