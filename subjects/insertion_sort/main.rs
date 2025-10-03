use insertion_sort::*;

fn main() {
    let mut target = [5, 3, 7, 2, 1, 6, 8, 4];
    // executes the first iteration of the algorithm
    insertion_sort(&mut target, 1);
    println!("{:?}", target);

    let mut target = [5, 3, 7, 2, 1, 6, 8, 4];
    // executes len - 1 iterations of the algorithm (sorts the slice)
    let len = target.len() - 1;
    insertion_sort(&mut target, len);
    println!("{:?}", target);
}
