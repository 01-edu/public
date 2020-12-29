/*
## searching

### Instructions

In this exercise you will have to complete the function `search`.
`search` receives an array and a key of `i32`, then it will return the position
of the given key in the array.

```
*/

fn main() {
    let array = [1, 3, 4, 6, 8, 9, 11];
    let key = search(&array, 6);
    println!(
        "the element 6 is in the position {:?} in the array {:?}",
        key, array
    );
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_search_a_value_in_an_array() {
        assert_eq!(search(&[6], 6), Some(0));
        assert_eq!(search(&[1, 2], 1), Some(0));
        assert_eq!(search(&[1, 2], 2), Some(1));
    }
    #[test]
    fn test_middle_of_an_array() {
        assert_eq!(search(&[1, 3, 4, 6, 8, 9, 11], 6), Some(3));
    }

    #[test]
    fn test_beginning_of_an_array() {
        assert_eq!(search(&[1, 3, 4, 6, 8, 9, 11], 1), Some(0));
    }

    #[test]
    fn test_end_of_an_array() {
        assert_eq!(search(&[1, 3, 4, 6, 8, 9, 11], 11), Some(6));
    }

    #[test]
    fn test_long_array() {
        assert_eq!(
            search(&[1, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 634], 144),
            Some(9)
        );
        assert_eq!(
            search(&[1, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377], 21),
            Some(5)
        );
    }

    #[test]
    fn test_value_is_not_included() {
        assert_eq!(search(&[1, 3, 4, 6, 8, 9, 11], 7), None);
        assert_eq!(search(&[1, 3, 4, 6, 8, 9, 11], 0), None);
        assert_eq!(search(&[1, 3, 4, 6, 8, 9, 11], 13), None);
        assert_eq!(search(&[], 1), None);
    }
}
