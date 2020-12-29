/*
## logic_number

### Instructions

In this exercise it will be given an example of a sequence of numbers, your purpose is to
determinate if the sequence returns true or false.
For this you have to create a function `number_logic` that will take a number `u32` and return true
if the number is the sum of its own digits, each raised to the power of the number of digits,
and false otherwise.

### Example:

    9 returns true, because 9 = 9^1 = 9
    10 returns false, because 10 != 1^2 + 0^2 = 1
    153 returns true, because: 153 = 1^3 + 5^3 + 3^3 = 1 + 125 + 27 = 153
    154 returns false, because: 154 != 1^3 + 5^3 + 4^3 = 1 + 125 + 64 = 190
*/
use logic_number::*;
fn main() {
    let array = [9, 10, 153, 154];
    for pat in &array {
        if number_logic(*pat) == true {
            println!(
                "this number returns {} because the number {} obey the rules of the sequence",
                number_logic(*pat),
                pat
            )
        }
        if number_logic(*pat) == false {
            println!("this number returns {} because the number {} does not obey the rules of the sequence", number_logic(*pat),pat )
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_zero() {
        assert!(number_logic(0))
    }

    #[test]
    fn test_single_digit_numbers() {
        assert!(number_logic(5));
        assert!(number_logic(9))
    }

    #[test]
    fn test_two_digit_numbers() {
        assert!(!number_logic(10))
    }

    #[test]
    fn test_three_or_more_digit_number() {
        assert!(number_logic(153));
        assert!(!number_logic(100));
        assert!(number_logic(9474));
        assert!(!number_logic(9475));
        assert!(number_logic(9_926_315));
        assert!(!number_logic(9_926_316))
    }
}
