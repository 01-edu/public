## box_recursion

### Instructions

Using the given code, create the following **associated functions**:

- `new`: which will initialize the `WorkEnvironment` with `grade` set to `None`.
- `add_worker`: which receives two strings, one being the role and the other the name of the worker. It will add the worker at the start of the list.
- `remove_worker`: which removes the last worker that was placed in the `WorkEnvironment`, this function returns an `Option` with the name of the worker.
- `last_worker`: which returns an `Option` with a tuple containing the name and role of the last added worker.

You must also create a type named `Link`. This will be the connection between the `WorkEnvironment` and `Worker` structures. This will be a recursion type, and it must point to `None` if there is no `Worker` to point to.

### Expected Functions and structures

```rust
#[derive(Debug)]
pub struct WorkEnvironment {
    pub grade: Link,
}

pub type Link =

#[derive(Debug)]
pub struct Worker {
    pub role: String,
    pub name: String,
    pub next: Link,
}

impl WorkEnvironment {
    pub fn new() -> WorkEnvironment {}
    pub fn add_worker(&mut self, role: String, name: String) {}
    pub fn remove_worker(&mut self) -> Option<String> {}
    pub fn last_worker(&self) -> Option<(String, String)> {}
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
    println!("{:#?}", list);

    println!("{:?}", list.last_worker());

    list.remove_worker();
    list.remove_worker();
    list.remove_worker();
    println!("{:?}", list);
    list.remove_worker();
    println!("{:?}", list);
}
```

And its output:

```console
$ cargo run
WorkEnvironment {
    grade: Some(
        Worker {
            role: "Normal Worker",
            name: "Alice",
            next: Some(
                Worker {
                    role: "Normal Worker",
                    name: "Ana",
                    next: Some(
                        Worker {
                            role: "Manager",
                            name: "Monica",
                            next: Some(
                                Worker {
                                    role: "CEO",
                                    name: "Marie",
                                    next: None,
                                },
                            ),
                        },
                    ),
                },
            ),
        },
    ),
}
Some(("Alice", "Normal Worker"))
WorkEnvironment { grade: Some(Worker { role: "CEO", name: "Marie", next: None }) }
WorkEnvironment { grade: None }
$
```

### Notions

- [Box\<T\>](https://doc.rust-lang.org/book/ch15-01-box.html)
- [Module std::option](https://doc.rust-lang.org/std/option/)
