## project_motion

### Instructions

For this exercise you will have to create a [projectile motion](https://cimg2.ck12.org/datastreams/f-d%3Abb024be6673110b31e78b46819e792adaed8dc661e082a61f0a6d64e%2BIMAGE%2BIMAGE.1).

A structure called `Object` will be provided which will have all variables that are
essential for the projectile physics. (distance, velocity, height, time)

You must implement :

- A function `throw_object` that will initialize the Object with a given velocity and height.
- The trait Iterator with the `.next()` in which,the next position of the object after 1 second, must be calculated.
  It will return an `Option` with the Object or it will return `None` if the object already reached the floor.

### Notions

- [trait Iterator](https://doc.rust-lang.org/std/iter/trait.Iterator.html)
- [iter](https://doc.rust-lang.org/rust-by-example/trait/iter.html)

### Expected Function

```rust

#[derive(Debug, Clone, PartialEq)]
pub struct Object {
    pub distance: f32,
    pub velocity: f32,
    pub height: f32,
    pub time: f32,
}

impl Object {
    pub fn throw_object(velocity: f32, height: f32) -> Object {}
}

impl Iterator for Object {
  // next
}

```

### Usage

Here is a program to test your function

```rust
use project_motion::*;

fn main() {
    let mut obj = Object::throw_object(50.0, 150.0);
    println!("{:?}", obj.next());
    println!("{:?}", obj.next());
    println!("{:?}", obj.next());
    println!("{:?}", obj.next());
    println!("{:?}", obj.next());
    println!("{:?}", obj.next());
}
```

And its output:

```console
$ cargo run
Some(Object { distance: 50.0, velocity: 50.0, height: 145.1, time: 1.0 })
Some(Object { distance: 100.0, velocity: 50.0, height: 125.5, time: 2.0 })
Some(Object { distance: 150.0, velocity: 50.0, height: 81.4, time: 3.0 })
Some(Object { distance: 200.0, velocity: 50.0, height: 3.0, time: 4.0 })
None
None
$
```
