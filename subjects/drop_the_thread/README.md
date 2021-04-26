## drop_the_thread

### Instructions

"Interior mutability is a design pattern in Rust that allows you to mutate data even when there are immutable references to that data"

in this exercise a Drop checker API has to be created. For this you must define:

- Two structures:

  - `Workers` that will have two fields:
    - `drops` that will save the number of dropped threads.
    - `states` that will save a state of multiple threads.
      If the thread is not dropped, the state will be false otherwise true.
  - `Thread` that will have the following fields:
    - `pid`, the id of the thread.
    - `cmd`, the name of the thread.
    - `parent`, that will be the link to the structure `Workers` (Tip: this must be a reference to the structure Workers)

- Implementation of each structure:

  - `Workers` :

    - `new`, that creates a default worker
    - `new_worker`, that returns a tuple with the `pid` and a new `Thread`,
      this function must receive a `String` being the `cmd`
      - `is_dropped`, that receives a `pid` and returns a `bool` that indicates the state of the thread by using the `pid`
      - `track_worker`, it should return a `usize`, that will be the last available index of the `states` vector, being the new next thread
    - `add_drop`, this function must be **called by the `Drop` trait**. It will receive a `pid` that will be used to change the
      state of the thread. If the state of that thread is `true` then it will panic with the message ("Cannot drop {}, because its already dropped", pid).
      Otherwise it should change the state to true and increment the `drops` field by one.

  - `Thread`:
    - `new_thread`, that initializes a new thread
    - `skill`, that drops the thread

- You must implement for the structure `Thread` the `Drop` trait. In this trait you must call the function `add_drop` so that the state of the thread changes

### Notions

- [Trait std::ops::Drop](https://doc.bccnsoft.com/docs/rust-1.36.0-docs-html/std/ops/trait.Drop.html)
- [Struct std::cell::RefCell](https://doc.rust-lang.org/std/cell/struct.RefCell.html)
- [Interior Mutability](https://doc.rust-lang.org/book/ch15-05-interior-mutability.html)

### Expected Functions

```rust
use std::cell::{RefCell, Cell};

#[derive(Debug, Default, Clone, Eq, PartialEq)]
pub struct Workers {
    pub drops: Cell<usize>,
    pub state: RefCell<Vec<bool>>
}

impl Workers {
    pub fn new() -> Workers {}
    pub fn new_worker(&self, c: String) -> (usize, Thread) {}
    pub fn track_worker(&self) -> usize {}
    pub fn is_dropped(&self, id: usize) -> bool {}
    pub fn add_drop(&self, id: usize) {}
}

#[derive(Debug, Clone, Eq, PartialEq)]
pub struct Thread<'a> {
    // expected public fields
}

impl<'a> Thread<'a> {
    pub fn new_thread(p: usize, c: String, t: &'a Workers) -> Thread {}
    pub fn skill(self) {}
}
```

### Usage

Here is a program to test your function,

```rust
use std::rc::Rc;
use drop_the_thread::*;

fn main() {
    let worker = Workers::new();
    let (id, thread) = worker.new_worker(String::from("command"));
    let (id1, thread1) = worker.new_worker(String::from("command1"));

    thread.skill();

    println!("{:?}", (worker.is_dropped(id), id, &worker.drops));

    thread1.skill();
    println!("{:?}", (worker.is_dropped(id1), id1, &worker.drops));

    let (id2, thread2) = worker.new_worker(String::from("command2"));
    let thread2 = Rc::new(thread2);
    let thread2_clone = thread2.clone();

    drop(thread2_clone);

    println!("{:?}", (worker.is_dropped(id2), id2, &worker.drops, Rc::strong_count(&thread2)));
}
```

And its output:

```console
student@ubuntu:~/drop_the_thread/test$ cargo run
(true, 0, Cell { value: 1 })
(true, 1, Cell { value: 2 })
(false, 2, Cell { value: 2 }, 1)
student@ubuntu:~/drop_the_thread/test$
```
