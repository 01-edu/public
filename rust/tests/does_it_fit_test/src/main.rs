use does_it_fit::*;

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
