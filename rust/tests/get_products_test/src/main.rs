/*
## get_products

### Instructions

Create a function `get_products` that takes a vector of integers, and returns a vector of the products
of each index. For this exercise to be correct you will have to return the product of every index
except the current one.

### Example:

Input: arr[]  = {10, 3, 5, 6, 2}
Output: prod[]  = {180, 600, 360, 300, 900}

*/
use get_products::*;
fn main() {
    let arr: Vec<usize> = vec![1, 7, 3, 4];
    let output = get_products(arr);
    println!("{:?}", output);
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_multiple() {
        let arr: Vec<usize> = vec![1, 7, 3, 4];
        let output = get_products(arr);
        let arr2: Vec<usize> = vec![10, 3, 5, 6, 2];
        let output2 = get_products(arr2);
        assert_eq!(output, vec![84, 12, 28, 21]);
        assert_eq!(output2, vec![180, 600, 360, 300, 900]);
    }

    #[test]
    fn test_empty_case() {
        let arr: Vec<usize> = Vec::new();
        let output = get_products(arr);
        assert_eq!(output, vec![]);
    }

    #[test]
    fn test_single_case() {
        let arr: Vec<usize> = vec![2];
        let output = get_products(arr);
        assert_eq!(output, vec![]);
    }
}
