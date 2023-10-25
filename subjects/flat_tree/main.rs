use flat_tree::*;
use std::collections::BTreeSet;

fn main() {
    let mut tree = BTreeSet::new();
    tree.insert(34);
    tree.insert(0);
    tree.insert(9);
    tree.insert(30);
    println!("{:?}", flatten_tree(&tree));

    let mut tree = BTreeSet::new();
    tree.insert("Slow");
    tree.insert("kill");
    tree.insert("will");
    tree.insert("Horses");
    println!("{:?}", flatten_tree(&tree));
}
