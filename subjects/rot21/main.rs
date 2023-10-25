use rot21::rot21;

fn main() {
    println!("The letter \"a\" becomes: {}", rot21("a"));
    println!("The letter \"m\" becomes: {}", rot21("m"));
    println!("The word \"MISS\" becomes: {}", rot21("MISS"));
    println!("Your cypher wil be: {}", rot21("Testing numbers 1 2 3"));
    println!("Your cypher wil be: {}", rot21("rot21 works!"));
}
