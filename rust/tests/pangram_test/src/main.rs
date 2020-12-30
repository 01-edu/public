/*
## pangram

### Instructions

Determine if the string is a pangram.
A pangram is a sentence using every letter of the alphabet at least once.

Example:

"The quick brown fox jumps over the lazy dog."

*/
use pangram::*;
use std::collections::HashSet;

fn main() {
    println!(
        "{}",
        is_pangram("the quick brown fox jumps over the lazy dog!")
    );
    println!("{}", is_pangram("this is not a pangram!"));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_empty_strings() {
        assert!(!is_pangram(""));
        assert!(!is_pangram(" "));
    }

    #[test]
    fn test_is_pangram() {
        assert!(is_pangram("the quick brown fox jumps over the lazy dog"));
        assert!(is_pangram("the_quick_brown_fox_jumps_over_the_lazy_dog"));
        assert!(is_pangram(
            "the 1 quick brown fox jumps over the 2 lazy dogs"
        ));
    }
    #[test]
    fn test_not_pangram() {
        assert!(!is_pangram(
            "a quick movement of the enemy will jeopardize five gunboats"
        ));
        assert!(!is_pangram("the quick brown fish jumps over the lazy dog"));
        assert!(!is_pangram("the quick brown fox jumps over the lay dog"));
        assert!(!is_pangram("7h3 qu1ck brown fox jumps ov3r 7h3 lazy dog"));
        assert!(is_pangram("\"Five quacking Zephyrs jolt my wax bed.\""));
        assert!(is_pangram(
            "Victor jagt zwölf Boxkämpfer quer über den großen Sylter Deich."
        ));
    }
}
