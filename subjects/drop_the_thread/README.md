## drop_the_thread

### Instructions

> Interior mutability is a design pattern in Rust that allows you to mutate data even when there are immutable references to that data.

In this exercise, you will create a Drop Checker API.

Define the following structures:


- `Workers`: containing:
  - `drops`: that will save the number of dropped threads.
  - `states`: that will save the state of multiple threads. If the thread is not dropped, the state will be `false`, and will be `true` otherwise.
- `Thread`: containing:
  - `pid`: the id of the thread.
  - `cmd`: the name of the thread.
  - `parent`: a link to the structure `Workers`. (Tip: this should be a reference).

You'll need to also add the following associated functions to the structures:

- `Workers` :
  - `new`: that creates a default worker.
  - `new_worker`: that returns a tuple with the `pid` and a new `Thread`. This function must receive a `String` representing the `cmd`.
  - `is_dropped`: that receives a `pid` and returns a `bool` that indicates the state of the thread.
  - `track_worker`: which returns a `usize` representing the length of the `states` vector. (The index of the next new thread).
  - `add_drop`: which is **called by the `Drop` trait**. It will receive a `pid` that will be used to change the state of the thread. If the state of that thread is `true` then it will panic with the message `"X is already dropped"`, where `X` represents the `pid`). Otherwise it should change the state to `true` and increment the `drops` field by 1.

- `Thread`:
  - `new_thread`: that initializes a new thread.
  - `skill`: that drops the thread.

- You must implement the `Drop` trait for the `Thread` structure. In this trait you must call the function `add_drop` so that the state of the thread changes.

### Expected Functions

```rust
use std::cell::{RefCell, Cell};

#[derive(Debug, Default, Clone, Eq, PartialEq)]
pub struct Workers {
    pub drops: Cell<usize>,
    pub states: RefCell<Vec<bool>>
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
$ cargo run
(true, 0, Cell { value: 1 })
(true, 1, Cell { value: 2 })
(false, 2, Cell { value: 2 }, 1)
$
```

### Notions

- [Trait std::ops::Drop](https://doc.bccnsoft.com/docs/rust-1.36.0-docs-html/std/ops/trait.Drop.html)
- [Struct std::cell::RefCell](https://doc.rust-lang.org/std/cell/struct.RefCell.html)
- [Interior Mutability](https://doc.rust-lang.org/book/ch15-05-interior-mutability.html)
