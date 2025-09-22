use smallest::smallest;
use std::collections::HashMap;

fn main() {
    let hash = HashMap::from([
        ("Cat", 122),
        ("Dog", 333),
        ("Elephant", 334),
        ("Gorilla", 14),
    ]);

    println!(
        "The smallest of the elements in the HashMap is {}",
        smallest(hash)
    );
}
