## ref_cell

### Instructions

Create a `Tracker` structure with the following fields:

- `messages`: a **public** vector of all the sent messages.
- `value`: the count of how many times the value was referenced. It should not exceed `max`.
- `max`: the maximum count of references.

Add the following methods to `Tracker`:

- `new`: that initializes the structure with a passed `max` value.
- `set_value`: that sets the structure's `value` to the passed argument `value`'s reference count only if it isn't higher than `max`. Should it exceed `max` it should write to `messages` the message `"Error: You can't go over your quota!"`. Should it not exceed `max` but exceed the percentage of 70%, it should write to `messages` the message `"Warning: You have used up over X% of your quota!"`, where `X` should be replaced with the calculated percentage.

- `peek`: that will take a peek at how much usage the passed argument `value` already has compared to the structure's `max`. It should write into `messages` the string `"Info: This value would use X% of your quota"`. `X` should be replaced with the calculated percentage.

### Usage

Here is a program to test your function,

```rust
use std::rc::Rc;

use ref_cell::*;

fn main() {
    let v = Rc::new(1); // we have one reference to this Rc

    // initialize the tracker, with the max number of
    // called references as 10
    let track = Tracker::new(10);

    let _v = Rc::clone(&v); // |\
    let _v = Rc::clone(&v); // | -> increase the Rc to 4 references
    let _v = Rc::clone(&v); // |/

    // take a peek of how much we already used from our quota
    track.peek(&v);

    let _v = Rc::clone(&v); // |\
    let _v = Rc::clone(&v); // |  -> increase the Rc to 8 references
    let _v = Rc::clone(&v); // | /
    let _v = Rc::clone(&v); // |/

    // this will change the tracker's inner value
    // and make a verification of how much we already used of our quota
    track.set_value(&v);

    let _v = Rc::clone(&v); // increase the Rc to 9 references
    let _v = Rc::clone(&v); // increase the Rc to 10 references, the maximum we allow

    track.set_value(&v);

    let _v = Rc::clone(&v); // surpass the maximum allowed references

    track.peek(&v);
    track.set_value(&v);

    track
        .messages
        .borrow()
        .iter()
        .for_each(|msg| println!("{}", msg));
}
```

And its output:

```console
$ cargo run
Info: This value would use 40% of your quota
Warning: You have used up over 80% of your quota!
Warning: You have used up over 100% of your quota!
Info: This value would use 110% of your quota
Error: You can't go over your quota!
$
```

### Notions

- [std::cell::RefCell](https://doc.rust-lang.org/std/cell/struct.RefCell.html)
- [Struct std::rc::Rc](https://doc.rust-lang.org/std/rc/struct.Rc.html)
