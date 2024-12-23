## circle

### Instructions

Create the structures `Circle` and `Point`. You'll need to create the necessary methods for the code in the [usage](#usage) to compile, and give the expected output.

#### Methods:

- `Point`:
  - `distance()` -> returns the distance between two coordinates.
- `Circle`:
  - `diameter()` -> returns the diameter of the circle.
  - `area()` -> returns the area of the circle.
  - `intersect()` -> returns if two circles intersect.

#### Associated Functions

- `Circle`:
  - `new()` -> receives three 64-bit floating point numbers in the following order: x, y and radius (x and y are the coordinates of the center of the new circle). The function returns a new circle.

### Expected Functions and Structures

This snippets are incomplete, you'll need to complete them. You'll find some useful information in the [usage](#usage).

```rust
#[derive(Debug, Clone, Copy)]
pub struct Circle {
	pub center //..
	pub radius //..
}

impl Circle {
    // ...
}

#[derive(Debug, Clone, Copy)]
pub struct Point(/* */);

impl Point {
    // ...
}
```

### Usage

Here is a program to test your function

```rust
use circle::*;

fn main() {
    let circle = Circle::new(500.0, 500.0, 150.0);
    let circle1 = Circle {
        center: Point(80.0, 115.0),
        radius: 30.0,
    };
    let point_a = Point(1.0, 1.0);
    let point_b = Point(0.0, 0.0);
    println!("circle = {:?} area = {}", circle, circle.area());
    println!("circle = {:?} diameter = {}", circle, circle.diameter());
    println!("circle1 = {:?} diameter = {}", circle1, circle1.diameter());
    println!(
        "circle and circle1 intersect = {}",
        circle.intersect(circle1)
    );

    println!(
        "distance between {:?} and {:?} is {}",
        point_a,
        point_b,
        point_a.distance(point_b)
    );
}
```

And its output

```console
$ cargo run
circle = Circle { center: Point(500.0, 500.0), radius: 150.0 } area = 70685.83470577035
circle = Circle { center: Point(500.0, 500.0), radius: 150.0 } diameter = 300
circle1 = Circle { center: Point(80.0, 115.0), radius: 30.0 } diameter = 60
circle and circle1 intersect = false
distance between Point(1.0, 1.0) and Point(0.0, 0.0) is 1.4142135623730951
$
```

### Notions

- [Using Structs](https://doc.rust-lang.org/book/ch05-00-structs.html)
- [f64 constants](https://doc.rust-lang.org/std/f64/consts/index.html)
