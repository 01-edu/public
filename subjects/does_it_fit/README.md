## does_it_fit

### Instructions

Using the `areas_volumes` module provided, create two **functions**:

- `area_fit` which receives 6 arguments and returns a boolean:

  - `x` and `y`, length and width of the Rectangle in which it is going to be tried to fit the geometrical shapes (both usize)
  - `objects`, the type of geometrical shape(s) which are going to be tried to be fitted in the square (areas_volumes::GeometricalShapes)
  - `times`, the number of geometrical shapes which are going to be tried to be fitted in the square (usize)
  - `a` and `b`, the dimensions which the plane(s) shape(s) passed will have (both usize)
  - `a` will refer to the side of the Square, the radius of the Circle, the side_a of the Rectangle or the base of the Triangle
  - `b` will refer to the side_b of the Rectangle or the height of the Triangle
  - `area_fit` should return `true` if the geometrical shape(s) fit inside of the square.

- `volume_fit` which receives 8 arguments and returns a boolean:

  - `x`, `y` and `z`, length, width and depth of the box in which it is going to be tried to fit the geometrical volumes (both usize)
  - `objects`, the type of geometrical volume(s) which are going to be tried to be fitted in the box (areas_volumes::GeometricalVolumes)
  - `times`, the number of geometrical volumes which are going to be tried to be fitted in the box (usize)
  - `a`, `b` and `c`, the dimensions which the geometrical volume(s) passed will have (all of them usize)
  - `a` will refer to the side of the Cube, the radius of the Sphere, the side_a of the Parallelepiped, the area of the base of the Triangular Pyramid or the base radius of the Cone
  - `b` will refer to the side_b of the Parallelepiped, the height of the Triangular Pyramid or the height of the Cone
  - `c` will refer to the side_c of the Parallelepiped
  - `volume_fit` should return `true` if the geometrical volume(s) fit inside of the box.

### Expected Functions

```rs
pub fn area_fit(
	x: usize,
	y: usize,
	objects: areas_volumes::GeometricalShapes,
	times: usize,
	a: usize,
	b: usize,
) -> bool {

}
pub fn volume_fit(
	x: usize,
	y: usize,
	z: usize,
	objects: areas_volumes::GeometricalVolumes,
	times: usize,
	a: usize,
	b: usize,
	c: usize,
) -> bool {

}
```

### areas_volumes.rs

```rust
pub enum GeometricalShapes {
	Square,
	Circle,
	Rectangle,
	Triangle,
}

pub enum GeometricalVolumes {
	Cube,
	Sphere,
	Cone,
	Pyramid,
	Parallelepiped,
}

pub fn square_area(side: usize) -> usize {
	side.pow(2)
}

pub fn triangle_area(base: usize, height: usize) -> f64 {
	(base as f64 * height as f64) / 2.0
}

pub fn circle_area(radius: usize) -> f64 {
	std::f64::consts::PI * (radius.pow(2) as f64)
}

pub fn rectangle_area(side_a: usize, side_b: usize) -> usize {
	side_a * side_b
}

pub fn cube_volume(side: usize) -> usize {
	side.pow(3)
}

pub fn sphere_volume(radius: usize) -> f64 {
	(4.0 / 3.0) * std::f64::consts::PI * (radius.pow(3) as f64)
}

pub fn triangular_pyramid_volume(base_area: f64, height: usize) -> f64 {
	(base_area * height as f64) / 3.0
}

pub fn parallelepiped_volume(side_a: usize, side_b: usize, side_c: usize) -> usize {
	side_a * side_b * side_c
}

pub fn cone_volume(base_radius: usize, height: usize) -> f64 {
	(1.0 / 3.0) * std::f64::consts::PI * base_radius.pow(2) as f64 * height as f64
}
```

### Usage

Here is a program to test your function:

```rust
use does_it_fit::*;

fn main() {
	println!(
		"Do 100 rectangles (2x1) fit in a 2 by 4 square? {}",
		area_fit(2, 4, GeometricalShapes::Rectangle, 100, 2, 1)
	);
	println!(
		"Do 3 triangles (5 base and 3 height) fit in a 5 by 5 square? {}",
		area_fit(5, 5, GeometricalShapes::Triangle, 3, 5, 3)
	);
	println!(
		"Do 3 spheres (2 radius) fit in a 5 by 5 by 5 box? {}",
		volume_fit(5, 5, 5, GeometricalVolumes::Sphere, 3, 2, 0, 0)
	);
	println!(
		"Does 1 parallelepiped (6 base, 7 height and depth 4) fit in a 5 by 7 by 5 parallelepiped? {}",
		volume_fit(5, 7, 5, GeometricalVolumes::Parallelepiped, 1, 6, 7, 4)
	);
}
```

And its output:

```sh
$ cargo run
Do 100 rectangles (2x1) fit in a 2 by 4 square? false
Do 3 triangles (5 base and 3 height) fit in a 5 by 5 square? true
Do 3 spheres (2 radius) fit in a 5 by 5 by 5 box? true
Does 1 parallelepiped (6 base, 7 height and depth 4) fit in a 5 by 7 by 5 parallelepiped? true
$
```
