## box_recursion

### Instructions

Using the given code create the following **functions**:

- `new` which will initialize the `WorkEnvironment` as `None`
- `add_worker`, which receives two strings, one being the type of worker and the other the name of the worker.
- `remove_worker`, which removes the last worker that was placed in the `WorkEnvironment`, this function should
  returns an `Option` with the name of the worker.
- `search_worker`, which returns a tuple with the name and type of worker.

You must also create a type called `Link`. This will be the connection of the structures `WorkEnvironment` and `Worker`.
Do not forget that this will be a recursion type and it must point to `None` if there is no workers.

### Notions

- [enum](https://doc.rust-lang.org/rust-by-example/custom_types/enum.html)
- [boc](https://doc.rust-lang.org/book/ch15-01-box.html)
- [Module std::option](https://doc.rust-lang.org/std/option/)

### Expected Functions and structures

```rust
pub struct WorkEnvironment {
    pub grade: Link,
}

pub type Link =

pub struct Worker {
    pub worker_type: String,
    pub worker_name: String,
    pub next_worker: Link,
}

impl WorkEnvironment {
    pub fn new() -> WorkEnvironment {}
    pub fn add_worker(&mut self, t: String, name: String) {}
    pub fn remove_worker(&mut self) -> Option<String> {}
    pub fn search_worker(&self) -> Option<(String, String)> {}
}


```

### Usage

Here is a program to test your function,

```rust
use box_recursion::*;

fn main() {
    let mut list = WorkEnvironment::new();
    list.add_worker(String::from("CEO"), String::from("Marie"));
    list.add_worker(String::from("Manager"), String::from("Monica"));
    list.add_worker(String::from("Normal Worker"), String::from("Ana"));
    list.add_worker(String::from("Normal Worker"), String::from("Alice"));
    println!("{:?}", list);

    println!("{:?}", list.search_worker());

    list.remove_worker();
    list.remove_worker();
    list.remove_worker();
    list.remove_worker();
    println!("{:?}", list);
}
```

And its output:

```console
student@ubuntu:~/box_recursion/test$ cargo run
WorkEnvironment { grade: Some(Worker { worker_type: "Normal Worker", worker_name: "Alice", next_worker: Some(Worker { worker_type: "Normal Worker", worker_name: "Ana", next_worker: Some(Worker { worker_type: "Manager", worker_name: "Monica", next_worker: Some(Worker { worker_type: "CEO", worker_name: "Marie", next_worker: None }) }) }) }) }
Some(("Alice", "Normal Worker"))
WorkEnvironment { grade: None }
student@ubuntu:~/box_recursion/test$
```
