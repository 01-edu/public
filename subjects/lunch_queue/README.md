## lunch_queue

### Instructions

You will need to create an _API_, so that a program can organize a queue of people.

The program requires the following functions. Add them as associated functions to the `Queue` structure:

- `new`: which will initialize the `Queue`.
- `add`: which adds a person to the queue.
- `invert_queue`: which reverses the queue.
- `rm`: which removes the person who finished ordering their food. The removal should respect the FIFO method (first in first out). It should return the person's details.
- `search`: which returns the details for a given person's `name`.

You must also create a type named `Link`. This will be the connection of the structures `Queue` and `Person`. This will be a recursion type, and must point to `None` if there is no `Person` to point to.

### Expected Function and Structures

```rust
pub struct Queue {
    pub node: Link,
}

pub type Link =

pub struct Person {
    pub discount: i32,
    pub name: String,
}

impl Queue {
    pub fn new() -> Queue {

    }
    pub fn add(&mut self, name: String, discount: i32) {

    }
    pub fn invert_queue(&mut self) {

    }
    pub fn rm(&mut self) -> Option<(String, i32)> {

    }
    pub fn search(&self, name: &str) -> Option<(String, i32)> {

    }
}
```

### Usage

Here is a program to test your function:

```rust
use lunch_queue::*;

fn main() {
    let mut list = Queue::new();
    list.add(String::from("Marie"), 20);
    list.add(String::from("Monica"), 15);
    list.add(String::from("Ana"), 5);
    list.add(String::from("Alice"), 35);
    println!("{:?}", list);

    println!("{:?}", list.search("Marie"));
    println!("{:?}", list.search("Alice"));
    println!("{:?}", list.search("someone"));

    println!("removed {:?}", list.rm());
    println!("list {:?}", list);
    list.invert_queue();
    println!("invert {:?}", list);
}
```

And its output:

```console
$ cargo run
Queue { node: Some(Person { name: "Alice", discount: 35, next_person: Some(Person { name: "Ana", discount: 5, next_person: Some(Person { name: "Monica", discount: 15, next_person: Some(Person { name: "Marie", discount: 20, next_person: None }) }) }) }) }
Some(("Marie", 20))
Some(("Alice", 35))
None
removed Some(("Marie", 20))
list Queue { node: Some(Person { name: "Alice", discount: 35, next_person: Some(Person { name: "Ana", discount: 5, next_person: Some(Person { name: "Monica", discount: 15, next_person: None }) }) }) }
invert Queue { node: Some(Person { name: "Monica", discount: 15, next_person: Some(Person { name: "Ana", discount: 5, next_person: Some(Person { name: "Alice", discount: 35, next_person: None }) }) }) }
$
```

### Notions

- [enum](https://doc.rust-lang.org/rust-by-example/custom_types/enum.html)
- [Box](https://doc.rust-lang.org/book/ch15-01-box.html)
- [std::option](https://doc.rust-lang.org/std/option/)
