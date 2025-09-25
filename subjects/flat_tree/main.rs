use flat_tree::*;
use std::collections::BTreeSet;

fn main() {
    let tree = BTreeSet::from([34, 0, 9, 30]);
    println!("{:?}", flatten_tree(&tree));

    let tree = BTreeSet::from(["Slow", "kill", "will", "Horses"]);
    println!("{:?}", flatten_tree(&tree));
}
