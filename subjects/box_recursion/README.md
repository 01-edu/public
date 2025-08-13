## box_recursion

### Instructions

Using the given code, create the following **associated functions**:

- `new`: which will initialize the `WorkEnvironment` with `grade` set to `None`.
- `add_worker`: which receives a name of a worker and a role. It will add the worker at the start of the list.
- `remove_worker`: which removes the last worker that was placed in the `WorkEnvironment`, this function returns an `Option` with the name of the removed worker.
- `last_worker`: which returns an `Option` with a tuple containing the name and role of the last added worker.

You must also create a `Role` enum which can be `CEO`, `Manager`, or `Worker`, and will represent the role of a worker. This enum should implement the `From` trait for `&str`.

Additionally, create a type named `Link`, which will be the connection between the `WorkEnvironment` and `Worker` structures. It will point to `None` if there is no `Worker` to point to.

### Expected Functions and structures

```rust
#[derive(Debug, PartialEq)]
pub enum Role {
    CEO,
    Manager,
    Worker,
}

impl From<&str> for Role {}

#[derive(Debug)]
pub struct WorkEnvironment {
    pub grade: Link,
}

pub type Link = _; // Complete type alias

#[derive(Debug)]
pub struct Worker {
    pub role: Role,
    pub name: String,
    pub next: Link,
}

impl WorkEnvironment {
    pub fn new() -> Self {
        todo!()
    }

    pub fn add_worker(&mut self, name: &str, role: &str) {
        todo!()
    }

    pub fn remove_worker(&mut self) -> Option<String> {
        todo!()
    }

    pub fn last_worker(&self) -> Option<(String, Role)> {
        todo!()
    }
}
```

### Usage

Here is a program to test your function,

```rust
use box_recursion::*;

fn main() {
    let mut list = WorkEnvironment::new();

    list.add_worker("Marie", "CEO");
    list.add_worker("Monica", "Manager");
    list.add_worker("Ana", "Normal Worker");
    list.add_worker("Alice", "Normal Worker");

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
            role: Worker,
            name: "Alice",
            next: Some(
                Worker {
                    role: Worker,
                    name: "Ana",
                    next: Some(
                        Worker {
                            role: Manager,
                            name: "Monica",
                            next: Some(
                                Worker {
                                    role: CEO,
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
Some(("Alice", Worker))
WorkEnvironment { grade: Some(Worker { role: CEO, name: "Marie", next: None }) }
WorkEnvironment { grade: None }
$
```

### Notions

- [Box\<T\>](https://doc.rust-lang.org/book/ch15-01-box.html)
- [Module std::option](https://doc.rust-lang.org/std/option/)
