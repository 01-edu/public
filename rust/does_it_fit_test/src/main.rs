/*
## doe_it_fit

### Instructions

Using the `areas_volumes` module provided, create two functions:

- `area_fit` that receives 6 arguments:
  - `x` and `y`, size of the square in which it is going to be tried to fit the geometrical shapes (both usize)
  - `objects`, the type of geometrical shape(s) that it is  going to be tried to be fitted in the square (areas_volumes::Geometrical_Shapes)
  - `times`, the number of geometrical shapes that are going to be tried to be fitted in the square (usize)
  - `a` and `b`, the dimensions that the plane(s) shape(s) passed will have (both usize)
    - `a` will refer to the side of the Square, the radius of the Circle, the side_a of the Rectangle or the base of the Triangle
    - `b` will refer to the side_b of the Rectangle or the height of the Triangle
- `area_fit` should return if the geometrical shape(s) fit inside of the square.

    - `volume_fit` that receives 8 arguments:
  - `x`, `y` and `z`, size of the box in which it is going to be tried to fit the geometrical volumes (both usize)
  - `objects`, the type of geometrical volume(s) that it is  going to be tried to be fitted in the box (areas_volumes::Geometrical_Volumes)
  - `times`, the number of geometrical volumes that are going to be tried to be fitted in the box (usize)
  - `a`, `b` and `c`, the dimensions that the geometrical volume(s) passed will have (all of them usize)
    - `a` will refer to the side of the Cube, the radius of the Sphere, the side_a of the Parallelepiped, the area of the base of the Triangular Pyramid or the base radius of the Cone
    - `b` will refer to the side_b of the Parallelepiped, the height of the Triangular Pyramid or the height of the Cone
    - `c` will refer to the side_c of the Parallelepiped
- `volume_fit` should return if the geometrical volume(s) fit inside of the box.


### Expected Functions (and Structures)

```rs
pub fn area_fit(
    x: usize,
    y: usize,
    objects: areas_volumes::Geometrical_Shapes,
    times: usize,
    a: usize,
    b: usize,
) {}

pub fn volume_fit(
    x: usize,
    y: usize,
    z: usize,
    objects: areas_volumes::Geometrical_Volumes,
    times: usize,
    a: usize,
    b: usize,
    c: usize,
) {}
```

### Usage

Here is a program to test your function:

```rust
fn main() {
    println!(
        "Does 100 rectangles (2x1) fit in a 2 by 4 square? {}",
        area_fit(2, 4, Geometrical_Shapes::Rectangle, 100, 2, 1)
    );
    println!(
        "Does 3 triangles (5 base and 3 height) fit in a 5 by 5 square? {}",
        area_fit(5, 5, Geometrical_Shapes::Triangle, 3, 5, 3)
    );
    println!(
        "Does 3 spheres (2 radius) fit in a 5 by 5 by 5 box? {}",
        volume_fit(5, 5, 5, Geometrical_Volumes::Sphere, 3, 2, 0, 0)
    );
    println!(
        "Does 3 triangles (5 base and 3 height) fit in a 5 by 7 by 5 box? {}",
        volume_fit(5, 7, 5, Geometrical_Volumes::Parallelepiped, 1, 6, 7, 4)
    );
}
```

And its output:

```sh
$ cargo run
false
true
false
true
$
```
*/
mod areas_volumes;

use areas_volumes::*;

fn main() {
    println!(
        "Does 100 rectangles (2x1) fit in a 2 by 4 square? {}",
        area_fit(2, 4, Geometrical_Shapes::Rectangle, 100, 2, 1)
    );
    println!(
        "Does 3 triangles (5 base and 3 height) fit in a 5 by 5 square? {}",
        area_fit(5, 5, Geometrical_Shapes::Triangle, 3, 5, 3)
    );
    println!(
        "Does 3 spheres (2 radius) fit in a 5 by 5 by 5 box? {}",
        volume_fit(5, 5, 5, Geometrical_Volumes::Sphere, 3, 2, 0, 0)
    );
    println!(
        "Does 3 triangles (5 base and 3 height) fit in a 5 by 7 by 5 box? {}",
        volume_fit(5, 7, 5, Geometrical_Volumes::Parallelepiped, 1, 6, 7, 4)
    );
}

#[cfg(test)]
mod tests {

    use super::*;

    #[test]
    fn no_volumes_shapes() {
        assert_eq!(true, area_fit(2, 5, Geometrical_Shapes::Circle, 0, 2, 1));
        assert_eq!(
            true,
            area_fit(2, 2, Geometrical_Shapes::Rectangle, 0, 6, 10)
        );
        assert_eq!(
            true,
            volume_fit(2, 5, 3, Geometrical_Volumes::Cone, 0, 1, 1, 1)
        );
        assert_eq!(
            true,
            volume_fit(3, 5, 7, Geometrical_Volumes::Parallelepiped, 0, 2, 6, 3)
        );
    }

    #[test]
    fn equal_size() {
        assert_eq!(true, area_fit(2, 5, Geometrical_Shapes::Square, 1, 2, 5));
        assert_eq!(
            true,
            volume_fit(3, 1, 4, Geometrical_Volumes::Cube, 1, 1, 3, 4)
        );
    }

    #[test]
    fn all_shapes() {
        assert_eq!(false, area_fit(3, 5, Geometrical_Shapes::Circle, 2, 2, 0));
        assert_eq!(true, area_fit(8, 6, Geometrical_Shapes::Triangle, 5, 5, 2));
        assert_eq!(true, area_fit(7, 3, Geometrical_Shapes::Rectangle, 2, 2, 4));
        assert_eq!(true, area_fit(5, 5, Geometrical_Shapes::Square, 1, 2, 4));
    }

    #[test]
    fn all_volumes() {
        assert_eq!(
            true,
            volume_fit(5, 6, 3, Geometrical_Volumes::Cube, 2, 3, 3, 4)
        );
        assert_eq!(
            false,
            volume_fit(7, 4, 4, Geometrical_Volumes::Cone, 1, 8, 2, 0)
        );
        assert_eq!(
            true,
            volume_fit(2, 5, 3, Geometrical_Volumes::Sphere, 1, 1, 1, 1)
        );
        assert_eq!(
            false,
            volume_fit(2, 5, 3, Geometrical_Volumes::Parallelepiped, 31, 1, 1, 1)
        );
        assert_eq!(
            true,
            volume_fit(7, 5, 3, Geometrical_Volumes::Pyramid, 3, 3, 2, 1)
        );
    }
}
