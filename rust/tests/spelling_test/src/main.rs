/*
## spelling

 ### Instructions

In this exercise a number between 0 and 1000000 will be generated.
Your purpose is to create the function `spell` that will spell the numbers generated.

So, if the program generates the number:

- 1 your function will return the string "one"
- 14 your function will return the string "fourteen".
- 96 your function will return the string "ninety-six"
- 100 your function will return the string "one hundred".
- 101 your function will return the string "one hundred one"
- 348 your function will return the string "one hundred twenty-three"
- 1002 your function will return the string "one thousand two".
- 1000000 your function will return the string "one million"
*/
use spelling::*;
use rand::Rng;

fn main() {
    let mut rng = rand::thread_rng();
    println!("{}", spell(rng.gen_range(0, 1000000)));
    println!("{}", spell(rng.gen_range(0, 1000000)));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_one() {
        assert_eq!(spell(0), String::from("zero"));
        assert_eq!(spell(1), String::from("one"));
        assert_eq!(spell(14), String::from("fourteen"));
        assert_eq!(spell(20), String::from("twenty"));
        assert_eq!(spell(22), String::from("twenty-two"));
        assert_eq!(spell(101), String::from("one hundred one"));
        assert_eq!(spell(120), String::from("one hundred twenty"));
        assert_eq!(spell(123), String::from("one hundred twenty-three"));
        assert_eq!(spell(1000), String::from("one thousand"));
        assert_eq!(spell(1055), String::from("one thousand fifty-five"));
        assert_eq!(
            spell(1234),
            String::from("one thousand two hundred thirty-four")
        );
        assert_eq!(
            spell(10123),
            String::from("ten thousand one hundred twenty-three")
        );
        assert_eq!(
            spell(910112),
            String::from("nine hundred ten thousand one hundred twelve")
        );
        assert_eq!(
            spell(651123),
            String::from("six hundred fifty-one thousand one hundred twenty-three")
        );

        assert_eq!(spell(810000), String::from("eight hundred ten thousand"));
        assert_eq!(spell(1000000), String::from("one million"));
    }
}
