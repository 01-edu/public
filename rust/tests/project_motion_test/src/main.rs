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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_short_distance() {
        let mut object = Object::throw_object(50.0, 20.0);

        assert_eq!(object.next(), Some(Object { distance: 50.0, velocity: 50.0, height: 15.1, time: 1.0 }));
        assert_eq!(object, Object { distance: 50.0, velocity: 50.0, height: 15.1, time: 1.0 });

        assert!(object.next().is_none(), "{:?} instead of None", object);
        assert!(object.next().is_none(), "{:?} instead of None", object);
    }

    #[test]
    fn test_media_distance() {
        let mut object = Object::throw_object(100.0, 30.0);

        assert_eq!(object.next(), Some(Object { distance: 100.0, velocity: 100.0, height: 25.1, time: 1.0 }));
        assert_eq!(object, Object { distance: 100.0, velocity: 100.0, height: 25.1, time: 1.0 });

        assert_eq!(object.next(), Some(Object { distance: 200.0, velocity: 100.0, height: 5.5, time: 2.0 }));
        assert_eq!(object, Object { distance: 200.0, velocity: 100.0, height: 5.5, time: 2.0 });

        assert!(object.next().is_none(), "{:?} instead of None", object);
    }

    #[test]
    fn test_long_distance() {
        let mut object = Object::throw_object(120.0, 100.0);

        assert_eq!(object.next(), Some(Object { distance: 120.0, velocity: 120.0, height: 95.1, time: 1.0 }));
        assert_eq!(object, Object { distance: 120.0, velocity: 120.0, height: 95.1, time: 1.0 });

        assert_eq!(object.next(), Some(Object { distance: 240.0, velocity: 120.0, height: 75.5, time: 2.0 }));
        assert_eq!(object, Object { distance: 240.0, velocity: 120.0, height: 75.5, time: 2.0 });

        assert_eq!(object.next(), Some(Object { distance: 360.0, velocity: 120.0, height: 31.4, time: 3.0 }));
        assert_eq!(object, Object { distance: 360.0, velocity: 120.0, height: 31.4, time: 3.0 });
        
        assert!(object.next().is_none(), "{:?} instead of None", object);
    }
}
