/*
## Highest

### Instructions

 In this exercise you will be given a `Numbers` struct.
 Your task is to write these methods:
 `List` that returns an `array` with every number in the struct
 `Latest` that returns an `Option<u32>` with the last added number
 `Highest` that return an `Option<u32>` with the highest number from the list,
 `Highest_Three` that returns a `Vec<u32>` with the three highest numbers.

 ### Notions

- https://doc.rust-lang.org/book/ch13-02-iterators.html

*/
use highest::*;

fn main() {
    let expected = [30, 500, 20, 70];
    let n = Numbers::new(&expected);
    println!("{:?}", n.List());
    
    println!("{:?}", n.Highest());
    println!("{:?}", n.Latest());
    println!("{:?}", n.Highest_Three());
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_List() {
        let expected = [30, 50, 20, 70];
        let n = Numbers::new(&expected);
        assert_eq!(n.List(), &expected);
    }

    #[test]
    fn test_Latest() {
        let n = Numbers::new(&[100, 0, 90, 30]);
        let f = Numbers::new(&[]);
        assert_eq!(n.Latest(), Some(30));
        assert!(f.Latest().is_none(), "It should have been None, {:?}", f.Latest());
    }

    #[test]
    fn test_Highest() {
        let n = Numbers::new(&[40, 100, 70]);
        let f = Numbers::new(&[]);
        assert_eq!(n.Highest(), Some(100));
        assert!(f.Highest().is_none(), "It should have been None, {:?}", f.Highest());
    }

    #[test]
    fn test_Highest_Three() {
        let e = Numbers::new(&[10, 30, 90, 30, 100, 20, 10, 0, 30, 40, 40, 70, 70]);
        let f = Numbers::new(&[40, 20, 40, 30]);
        let g = Numbers::new(&[30, 70]);
        let h = Numbers::new(&[40]);
        let i = Numbers::new(&[]);
        let j = Numbers::new(&[20, 10, 30]);
        assert_eq!(e.Highest_Three(), vec![100, 90, 70]);
        assert_eq!(f.Highest_Three(), vec![40, 40, 30]);
        assert_eq!(g.Highest_Three(), vec![70, 30]);
        assert_eq!(h.Highest_Three(), vec![40]);
        assert!(i.Highest_Three().is_empty());
        assert_eq!(j.Highest_Three(), vec![30, 20, 10]);
    }
}
