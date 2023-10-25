use smallest::smallest;
use std::collections::HashMap;

fn main() {
    let mut hash = HashMap::new();
    hash.insert("Cat", 122);
    hash.insert("Dog", 333);
    hash.insert("Elephant", 334);
    hash.insert("Gorilla", 14);

    println!(
        "The smallest of the elements in the HashMap is {}",
        smallest(hash)
    );
}
