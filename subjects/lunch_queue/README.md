## lunch_queue

### Instructions

An API will have to be created to organize a queue of people.

Using the given code declare the following functions:

- `new` which will initialize the `Queue`
- `add` which receives the person's information, so it can be added to the `Queue`
- `invert_queue` which inverts the queue of persons
- `rm`, which will remove the person who finished ordering their food.
  The removal should respect a FIFO system (first in first out). This function should return a tuple wrapped in an `Option` with the information of the removed person (check the usage)
- `search`, which returns a tuple with the information of a given person `id`

You must also create a type called `Link`. This will be the connection of the structures `Queue` and `Person`.
Do not forget that this will be a recursion type and it must point to `None` if there is no persons.

### Notions

- [enum](https://doc.rust-lang.org/rust-by-example/custom_types/enum.html)
- [Box](https://doc.rust-lang.org/book/ch15-01-box.html)
- [std::option](https://doc.rust-lang.org/std/option/)

### Expected Function adn Structures

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

Here is a program to test your function:

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
    list.invert_queue();
    println!("invert {:?}", list);
}
```

And its output:

```console
student@ubuntu:~/lunch_queue/test$ cargo run
Queue { node: Some(Person { name: "Alice", discount: 35, next_person: Some(Person { name: "Ana", discount: 5, next_person: Some(Person { name: "Monica", discount: 15, next_person: Some(Person { name: "Marie", discount: 20, next_person: None }) }) }) }) }
Some(("Marie", 20))
Some(("Alice", 35))
None
removed Some(("Marie", 20))
list Queue { node: Some(Person { name: "Alice", discount: 35, next_person: Some(Person { name: "Ana", discount: 5, next_person: Some(Person { name: "Monica", discount: 15, next_person: None }) }) }) }
invert Queue { node: Some(Person { name: "Monica", discount: 15, next_person: Some(Person { name: "Ana", discount: 5, next_person: Some(Person { name: "Alice", discount: 35, next_person: None }) }) }) }
student@ubuntu:~/lunch_queue/test$
```
