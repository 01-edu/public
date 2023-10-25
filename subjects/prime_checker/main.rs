use prime_checker::*;

fn main() {
    println!("Is {} prime? {:?}", 2, prime_checker(2));
    println!("Is {} prime? {:?}", 14, prime_checker(14));
    println!("Is {} prime? {:?}", 2147483647, prime_checker(2147483647));
}
