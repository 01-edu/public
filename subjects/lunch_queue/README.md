## lunch_queue

### Instructions

You will have to create an API to organize a queue of people.

Using the given code create the following functions:

- `new` that will initialize the `Queue`.
- `add` that receives the person's information, so it can be added to the `Queue`
- `invert_queue` that invert the queue of persons
- `rm`, that will remove the person that finished ordering their food.
  The method for organizing the manipulation of a data structure should be a
  FIFO (first in first out) process. This function should return a tuple wrapped in an `Option`
  with the information of that person
- `search`, that return a tuple with the information of a given person `id`.

You must also create a type called `Link` this will be the connection of the structures `Queue` and `Person`.
Do not forget that this will be a recursion type and it must point to `None` if there is no persons.

### Expected Function

```rust
pub struct Queue {
    pub node: Link,
}

pub type Link =

pub struct Person {
    pub id: i32,
    pub name: String,
}

impl Queue {
    pub fn new() -> Queue {}
    pub fn add(&mut self, t: String, name: String) {}
    pub fn invert_queue(&mut self) {}
    pub fn rm(&mut self) -> Option<String> {}
    pub fn search(&self) -> Option<(String, String)> {}

}
```

### Usage

Here is a program to test your function

```rust
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
    println!("invert {:?}", list.invert_queue());
}
```

And its output:

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
Queue { node: Some(Person { name: "Alice", discount: 35, next_person: Some(Person { name: "Ana", discount: 5, next_person: Some(Person { name: "Monica", discount: 15, next_person: Some(Person { name: "Marie", discount: 20, next_person: None }) }) }) }) }
Some(("Marie", 20))
Some(("Alice", 35))
None
removed Some(("Marie", 20))
removed Some(("Monica", 15))
list Queue { node: Some(Person { name: "Alice", discount: 35, next_person: Some(Person { name: "Ana", discount: 5, next_person: None }) }) }
inverted list Queue { node: Some(Person { name: "Ana", discount: 5, next_person: Some(Person { name: "Alice", discount: 35, next_person: None }) }) }
student@ubuntu:~/[[ROOT]]/test$
```

### Notions

- https://doc.rust-lang.org/rust-by-example/custom_types/enum.html
- https://doc.rust-lang.org/book/ch15-01-box.html
- https://doc.rust-lang.org/std/option/
- https://doc.rust-lang.org/book/ch15-01-box.html
