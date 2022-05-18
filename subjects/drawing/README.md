## drawing

### Instructions

The purpose of this exercise is to create an image  in the example bellow:

![example](image.png)

For this exercise you will need to do the following:

- Copy the code in [usage](#usage) to your `main.rs`.

- Create a module called `geometrical_shapes` in another file.

This module will keep all the logic for creating and operating with the different geometrical shapes. You need to define two traits, `Displayable` and `Drawable`.

- `Drawable` which contains the methods `draw` and `color`.

- `Displayable` which contains the method `display`.

Define them according to the way they are called in the `main.rs` function.

In order to make the code in the `main.rs` compile and run, you have to define the following structures:

- `Point`
- `Circle`
- `Line`
- `Rectangle`
- `Triangle`

You are free to implement all the shapes with whatever internal structure you see fit, but you must provide an associated function `new` for all the shapes, which will be described below:

- `Point`: a new point should be created from two `i32` values.
- `Line`: a new line should be created from references to two different points. Also it will define an associated function called `random`:
    - `random`: receives two arguments, the first is the `maximum x` value a point can have and the second is the `maximum y` value that a point can have.
- `Triangle`: a new triangle should be created from references to three different points.
- `Rectangle`: a new rectangle should be created from two references to different points. ??????????????????????????????
- `Circle`: a new circle should be created from a point representing the center and an `i32` value representing the radius for the circle.

- The main function also requires a definition of an associated function called `random` for the types `Line`, `Point` and `Circle`. You should derive their signature from the usage.

**Note**: Don't forget to add the dependencies in your Cargo.toml.

#### Bonus

- Implement the possibility of drawing a pentagon (implement the structure Pentagon, and the trait needed to draw in the image)

- Implement the possibility of drawing a cube (implement the structure Cube, and the trait needed to draw in the image)

### Usage

```rust
mod geometrical_shapes;

use geometrical_shapes as gs;
use gs::{Displayable, Drawable};
use raster::{Color, Image};

fn main() {
    let mut image = Image::blank(1000, 1000);

    gs::Line::random(image.width, image.height).draw(&mut image);

    gs::Point::random(image.width, image.height).draw(&mut image);

    let rectangle = gs::Rectangle::new(&gs::Point::new(150, 150), &gs::Point::new(50, 50));
    rectangle.draw(&mut image);

    let triangle = gs::Triangle::new (
            &gs::Point::new(500, 500),
            &gs::Point::new(250, 700),
            &gs::Point::new(700, 800),
    );
    triangle.draw(&mut image);

    for _ in 1..50 {
        gs::Circle::random(image.width, image.height).draw(&mut image);
    }

    raster::save(&image, "image.png").unwrap();
}

impl Displayable for Image {
    fn display(&mut self, x: i32, y: i32, color: Color) {
        if x >= 0 && x < self.width && y >= 0 && y < self.height {
            self.set_pixel(x, y, color).unwrap();
        }
    }
}
```

**Note**: The expected output is a png file: `image.png`

### Notions

- [Image processing library](https://docs.rs/raster/0.2.0/raster/)

- [Traits](https://doc.rust-lang.org/stable/book/ch10-02-traits.html)