## does_it_fit

### Instructions

Use the `areas_volumes` module provided to create two **functions**.

Create `area_fit`. It should return `true` if the geometric shape can fit inside the rectangular area as many times as is indicated by `times`.

The arguments of the function will be:
- `x` and `y`: length and width of the rectangular area.
- `objects`: the type of geometric shape.
- `times`: the number of geometric shapes to fit inside the rectangular area.
- `a`: size of dimension for:
	- `side` of a `Square`
	- `base` of a `Triangle`
	- `radius` of a `Circle`
	- `side_a` of a `Rectangle`
- `b`: size of dimension for:
	- `height` of a `Triangle`
	- `side_b` of a `Rectangle`

Create `volume_fit`. It should return `true` if the geometric volume can fit inside the box as many times as is indicated by `times`.

The arguments of the function will be:
  - `x`, `y` and `z`: length, width and depth of the box.
  - `objects`: the type of geometric volume.
  - `times`: the number of geometric volumes to fit inside the box.
  - `a`: size of dimension for:
    -  `side` of a `Cube`
	- `radius` of a `Sphere`
	- `base_area` of a triangular `Pyramid`
	- `side_a` of a `Parallelepiped`
	- `base_radius` of a `Cone`
- `b`: size dimension for:
	- `height` of the triangular `Pyramid`
	- `side_b` of a `Parallelepiped`
	- `height` of the `Cone`
- `c`: size dimension for:
	- `side_c` of a `Parallelepiped`

### Expected Functions

```rust
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
