// Implement the function bubble_sort which receives a vector Vec<i32>
// and return the same vector but in increasing order using the bubble
// sort algorithm

use collect::*;

fn main() {
	let ref mut v = vec![3, 2, 4, 5, 1, 7];
	let mut b = v.clone();
	bubble_sort(v);
	println!("{:?}", v);

	b.sort();
	println!("{:?}", b);
}

#[test]
fn test_ordering() {
	let ref mut v = vec![3, 2, 4, 5, 1, 7];
	let mut b = v.clone();

	b.sort();
	bubble_sort(v);
	assert_eq!(*v, b);
}
