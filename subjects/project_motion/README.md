## project_motion

### Instructions

For this exercise, you will have to create a [projectile motion](https://cimg2.ck12.org/datastreams/f-d%3Abb024be6673110b31e78b46819e792adaed8dc661e082a61f0a6d64e%2BIMAGE%2BIMAGE.1).

Two structures will be provided. A structure called `ThrowObject` that will contain all the variables that are
essential for the projectile physics (initial position, initial velocity, current position, current velocity and time).

A structure called `Object` which will have the corresponding values of X and Y of the initial position, the initial velocity, the current position and the current velocity. 

You must implement :

- A function `new` that will initialize the ThrowObject with a given initial position and an initial velocity.
- The trait Iterator with the `.next()` in which the position and speed of the object must be calculated after 1 second.
  It will return an `Option` with the ThrowObject, or it will return `None` if the ThrowObject has already reached the floor.

Consider the value of gravity is 9.8m/(s*s) and that the position (p) in the instant s of an object is given by:

![Position Formula](position_formula.png)

and velocity (v) in the instant s of an object is given by:

![Speed Formula](speed_formula.png)


### Notions

- [trait Iterator](https://doc.rust-lang.org/std/iter/trait.Iterator.html)
- [iter](https://doc.rust-lang.org/rust-by-example/trait/iter.html)

### Expected Function

```rust

#[derive(Debug, Clone, PartialEq)]
pub struct Object {
    pub x: f32,
    pub y: f32,
}

#[derive(Debug, Clone, PartialEq)]
pub struct ThrowObject {
    pub init_position: Object,
    pub init_velocity: Object,
    pub actual_position: Object,
    pub actual_velocity: Object,
    pub time: f32,
}

impl ThrowObject {
    pub fn new(init_position: Object, init_velocity: Object) -> ThrowObject {
    }
}

impl Iterator for ThrowObject {
  // next
}

```

### Usage

Here is a program to test your function

```rust
use project_motion::*;

fn main() {
    let mut obj = ThrowObject::new(Object { x: 50.0, y: 50.0 }, Object { x: 0.0, y: 0.0 });
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
Some(ThrowObject { init_position: Object { x: 50.0, y: 50.0 }, init_velocity: Object { x: 0.0, y: 0.0 }, actual_position: Object { x: 50.0, y: 45.1 }, actual_velocity: Object { x: 0.0, y: -9.8 }, time: 1.0 })
Some(ThrowObject { init_position: Object { x: 50.0, y: 50.0 }, init_velocity: Object { x: 0.0, y: 0.0 }, actual_position: Object { x: 50.0, y: 30.4 }, actual_velocity: Object { x: 0.0, y: -19.6 }, time: 2.0 })
Some(ThrowObject { init_position: Object { x: 50.0, y: 50.0 }, init_velocity: Object { x: 0.0, y: 0.0 }, actual_position: Object { x: 50.0, y: 5.9 }, actual_velocity: Object { x: 0.0, y: -29.4 }, time: 3.0 })
None
None
$
```
