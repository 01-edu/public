/*
 ## rot

### Instructions

By now you will have the knowledge of the so called rotational cipher "ROT13".

A ROT13 on the Latin alphabet would be as follows:

- Plain: abcdefghijklmnopqrstuvwxyz

- Cipher: nopqrstuvwxyzabcdefghijklm

Your purpose in this exercise is to create a similar `rot` function that is a better version of the ROT13 cipher.
Your function will receive a string and a number and it will rot each letter of that string, the number of times, settled by the second argument, to the right or to the left if the number are negative.

Your function should only change letters. If the string includes punctuation and numbers
they will remain the same.

### Notions

- https://doc.rust-lang.org/book/ch18-00-patterns.html

### Expected functions

```rust

```

### Usage

Here is a program to test your function.
*/

fn main() {

    println!("The letter \"a\" becomes: {}", rot("a", 26));
    println!("The letter \"m\" becomes: {}", rot("m", 0));
    println!("The letter \"m\" becomes: {}", rot("m", 13));
    println!("The letter \"a\" becomes: {}", rot("a", 15));
    println!("The word \"MISS\" becomes: {}", rot("MISS", 5));
    println!(
        "The decoded message is: {}",
        rot("Gur svir obkvat jvmneqf whzc dhvpxyl.", 13)
    );
    println!(
        "The decoded message is: {}",
        rot("Mtb vznhpqd ifky ozrunsl ejgwfx ajc", 5)
    );
    println!(
        "Your cypher wil be: {}",
        rot("Testing with numbers 1 2 3", 4)
    );
    println!("Your cypher wil be: {}", rot("Testing", -14));
    println!("The letter \"a\" becomes: {}", rot("a", -1));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_simple() {
        assert_eq!("z", rot("m", 13));
        assert_eq!("n", rot("m", 1));
        assert_eq!("a", rot("a", 26));
        assert_eq!("z", rot("a", 25));
        assert_eq!("b", rot("l", 16));
        assert_eq!("j", rot("j", 0));
        assert_eq!("RNXX", rot("MISS", 5));
        assert_eq!("M J Q Q T", rot("H E L L O", 5));
    }

    #[test]
    fn test_all_letters() {
        assert_eq!(
            "Gur svir obkvat jvmneqf whzc dhvpxyl.",
            rot("The five boxing wizards jump quickly.", 13)
        );
    }

    #[test]
    fn test_numbers_punctuation() {
        assert_eq!(
            "Xiwxmrk amxl ryqfivw 1 2 3",
            rot("Testing with numbers 1 2 3", 4)
        );
        assert_eq!("Gzo\'n zvo, Bmviyhv!", rot("Let\'s eat, Grandma!", 21));
    }

    #[test]
    fn test_negative() {
        assert_eq!("z", rot("a", -1));
        assert_eq!("W", rot("A", -4));
        assert_eq!("Fqefuzs", rot("Testing", -14));
    }
}
