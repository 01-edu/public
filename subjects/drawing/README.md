## drawing

### Instructions

- Copy the code in [usage](#usage) to your main.rs

- Create a module called geometrical_shapes in another file.

- This module will keeps all the logic for creating and operating with the different geometrical shapes and define two traits `Displayable` and `Drawable`.

- `Drawable` contains the methods `draw` and `color`

- `Displayable` contains the methods `display`.

- Define them in correspondence with the way they're called in the main function

- You have to define the structures for Point, Circle, Line, Rectangle and Triangle and make the code in `main.rs` compile and run.

- You are free to implement all the shapes with the internal structure that you find more adequate, but you have to provide for all the shapes an associated function `new` which will be described next:

- Point: a new point should be created from two i32 values
- Line: a new line should be created from references to two points also define an associated function called `random` that receives two argument the first is the maximum x value a point can have and the
second the maximum y value that a point can have
- Triangle: a new triangle should be created from references to three points
- Rectangle: a new rectangle should be created from two reference to points
- Circle: a new circle should be created from a point represented the center and a i32 value representing the radius
   - also define an associated function called `random` that receives two arguments the first is the maximum x value the center point can have and the second the maximum y value that the center point can have

Don't forget to add the dependencies in your Cargo.toml.

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

    let rectangle = gs::Rectangle::new(gs::Point::new(150, 150), gs::Point::new(50, 50));
    rectangle.draw(&mut image);

    let triangle = gs::Triangle {
        vertices: (
            gs::Point::new(500, 500),
            gs::Point::new(250, 700),
            gs::Point::new(700, 800),
        ),
    };
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

### And the expected output is a png file: `image.png`
