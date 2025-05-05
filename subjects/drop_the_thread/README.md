## drop_the_thread

### Instructions

> Interior mutability is a design pattern in Rust that allows you to mutate data even with immutable references to that data.

In this exercise, you will create a Drop Checker API.

Define the following structures:

- `ThreadPool` containing:
  - `drops`: will save the number of dropped threads.
  - `states`: will save the state of multiple threads. If the thread is not dropped, the state will be `false`, and will be `true` otherwise.

- `Thread` containing:
  - `pid`: the id of the thread.
  - `cmd`: the name of the thread.
  - `parent`: a reference to the structure `ThreadPool`.

You'll also need to add the following associated functions to the structures:

- `ThreadPool` :
  - `new`: initializes the structure with default values.
  - `new_thread`: creates a new thread and mutates own state accordingly. Returns a tuple with the `pid` and a new `Thread`.
  - `is_dropped`: that receives a `pid` and returns whether the thread was already dropped.
  - `thread_len`: returns the length of the `states` vector. (The index of the next new thread).
  - `drop_thread`: this **should be called when dropping a `Thread`**. It will receive a `pid` that will be used to change the state of the thread. If the thread was already dropped then it will panic with the message `"X is already dropped"`, where `X` represents the `pid`. Otherwise it should change the state to `true` and increment the `drops` field by 1.

- `Thread`:
  - `new`: initializes a new thread.
  - `skill`: function that drops the thread.

- You must implement the `Drop` trait for the `Thread` structure. In this trait you must call the function `add_drop` so that the state of the thread changes.

### Expected Functions

```rust
use std::cell::{Cell, RefCell};

#[derive(Debug)]
pub struct ThreadPool {
    pub drops: Cell<usize>,
    pub states: RefCell<Vec<bool>>
}

impl ThreadPool {
    pub fn new() -> Self {
        todo!()
    }

    pub fn new_thread(&self, c: String) -> (usize, Thread) {
        todo!()
    }

    pub fn thread_len(&self) -> usize {
        todo!()
    }

    pub fn is_dropped(&self, id: usize) -> bool {
        todo!()
    }

    pub fn drop_thread(&self, id: usize) {
        todo!()
    }
}

#[derive(Debug)]
pub struct Thread {
    // expected public fields
}

impl<'a> Thread<'a> {
    pub fn new(p: usize, c: String, t: &'a ThreadPool) -> Self {
        todo!()
    }

    pub fn skill(self) {
        todo!()
    }
}

impl Drop for Thread<'_> {}
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

- [Trait std::ops::Drop](https://doc.rust-lang.org/std/ops/trait.Drop.html)
- [Struct std::cell::RefCell](https://doc.rust-lang.org/std/cell/struct.RefCell.html)
- [Interior Mutability](https://doc.rust-lang.org/book/ch15-05-interior-mutability.html)
