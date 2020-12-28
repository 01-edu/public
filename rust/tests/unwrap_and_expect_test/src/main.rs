use unwrap_and_expect::*;

fn main() {
    println!("{:?}", unwrap_or(vec![1, 3, 2, 5]));
    println!("{:?}", unwrap_or(vec![1, 3, 5]));
    println!("{:?}", unwrap_err(vec![1, 3, 2, 5]));
    println!("{:?}", unwrap(vec![1, 3, 5]));
    println!("{:?}", unwrap_or_else(vec![1, 3, 5]));
    println!("{:?}", unwrap_or_else(vec![3, 2, 6, 5]));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    #[should_panic(expected = "ERROR : (\"There is a even value in the vector!\", [2])")]
    fn test_expect() {
        expect(vec![1, 3, 2, 5]);
    }
    #[test]
    #[should_panic(expected = "called `Result::unwrap()` on an `Err` value: (\"There is a even value in the vector!\", [2])")]
    fn test_unwrap() {
        unwrap(vec![1, 3, 2, 5]);
    }
    #[test]
    #[should_panic]
    fn test_unwrap_err() {
        unwrap_err(vec![1, 3, 5]);
    }
    #[test]
    fn test_unwrap_or() {
        assert_eq!(unwrap_or(vec![1, 3, 2, 5]), vec![]);
    }
    #[test]
    fn test_unwrap_or_else() {
        assert_eq!(unwrap_or_else(vec![1, 3, 5]), vec![2, 4, 6]);
        assert_eq!(unwrap_or_else(vec![6, 8, 3, 2, 5, 4]), vec![6, 8, 2, 4]);
    }
    #[test]
    fn test_ok() {
        assert_eq!(expect(vec![1, 3, 5]), vec![2, 4, 6]);
        assert_eq!(unwrap_or(vec![1, 3, 5]), vec![2, 4, 6]);
        assert_eq!(unwrap_or_else(vec![1, 3, 5]), vec![2, 4, 6]);
        assert_eq!(unwrap(vec![1, 3, 5]), vec![2, 4, 6]);
        assert_eq!(unwrap_err(vec![1, 2, 3, 4, 5]).0, "There is a even value in the vector!");
        assert_eq!(unwrap_err(vec![1, 2, 3, 4, 5]).1, vec![2, 4]);
    }
}
