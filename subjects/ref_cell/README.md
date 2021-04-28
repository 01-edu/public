## ref_cell

### Instructions

#### First part (messenger.rs)

Create a module called `messenger`. This module will be able to inform a user of how many references of a given value he is using.
The main objective of this module is to limit how many times a value is referenced.

For this module the following must be created:

- A trait `Logger` which implements three functions: `warning`, `info`, `error`. All **functions** should receive a reference to themselves and a string literal.

```rust
    fn warning(&self, msg: &str);
    fn info(&self, msg: &str);
    fn error(&self, msg: &str);
```

- A structure called `Tracker`, which must have the fields: `logger` being a reference to the `Logger`, `value` being the count of how many times the value was referenced,
  `max` being the max count of references the actual value can achieve.

- An implementation of three functions that are associated to the `Tracker` structure:
  - `new` that will initialize the structure
  - `set_value` that sets the value to the `Tracker` structure and writes to the trait functions. This should be done comparing the **max** and the number of referenced of the actual value.
    If the percentage is equal or greater to 100% of the limit usage, it should write **"Error: you are over your quota!"** to the `error` function
    If the percentage is equal or greater to 70% of the limit usage, it should write **("Warning: you have used up over {}% of your quota! Proceeds with precaution", <calculated_percentage>)** to the `warning` function
  - `peek` that will take a peek of how much usage the variable already has. It should write **("Info: you are using up too {}% of your quote", <calculated_percentage>)** to the `info` function

### Second part (lib.rs)

Afterwards you must use the module `messenger` and create the following:

- A structure `Worker` that has the fields:
  - `track_value` this will be the value that will be tracked by the tracker.
  - `mapped_messages` that will have the latest messages. This must be a HashMap with the key being the type of message
    sent by the logger (info, error or warning) and the value being the message
  - `all_messages` that will be a vector of all messages sent.
- A `new` function that initializes the structure `Worker`
- To use the trait `Logger` you must implement it for the Worker structure. Each function (warning, error and info) must insert the message to the
  respective fields of the structure Worker.

You must use **interior mutability**, this means it must be able to mutate data even when there are immutable references to that data. Consequently, the user will not need to use the keyword `mut` (tip: RefCell)

### Notions

- [std::cell::RefCell](https://doc.rust-lang.org/std/cell/struct.RefCell.html)
- [Struct std::rc::Rc](https://doc.rust-lang.org/std/rc/struct.Rc.html)

### Usage

Here is a program to test your function,

```rust
use ref_cell::*;

fn main() {
    // initialize the worker
    let logger = Worker::new(1);

    // initialize the tracker, with the max number of
    // called references as 10
    let track = Tracker::new(&logger, 10);

    let _a = logger.track_value.clone();    // |\
    let _a1 = logger.track_value.clone();   // | -> increase the Rc to 4 references
    let _a2 = logger.track_value.clone();   // |/

    // take a peek of how much we already used from our quota
    track.peek(&logger.track_value);

    let _b = logger.track_value.clone();  // |\
    let _b1 = logger.track_value.clone(); // |  -> increase the Rc to 8 references
    let _b2 = logger.track_value.clone(); // | /
    let _b3 = logger.track_value.clone(); // |/

    // this will set the value and making a verification of
    // how much we already used of our quota
    track.set_value(&logger.track_value);

    let _c = logger.track_value.clone(); // | -> increase the Rc to 9 references

    // this will set the value and making a verification of
    // how much we already used of our quota
    track.set_value(&logger.track_value);

    let _c1 = logger.track_value.clone(); // | -> increase the Rc to 10 references, this will be the limit

    track.set_value(&logger.track_value);

    for (k ,v) in logger.mapped_messages.into_inner() {
        println!("{:?}", (k ,v));
    }
    println!("{:?}", logger.all_messages.into_inner());
}
```

And its output:

```console
$ cargo run
("Info", "you are using up too 40% of your quote")
("Warning", "you have used up over 90% of your quota! Proceeds with precaution")
("Error", "you are over your quota!")
[
  "Info: you are using up too 40% of your quote",
  "Warning: you have used up over 80% of your quota! Proceeds with precaution",
  "Warning: you have used up over 90% of your quota! Proceeds with precaution",
  "Error: you are over your quota!"
]
$
```
