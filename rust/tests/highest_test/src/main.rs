/*
## highest

### Instructions

 In this exercise you will be given a `Numbers` struct.
 Your task is to write these methods:
 `list` that returns an `array` with every number in the struct
 `latest` that returns an `Option<u32>` with the last added number
 `highest` that return an `Option<u32>` with the highest number from the list,
 `highest_three` that returns a `Vec<u32>` with the three highest numbers.

 ### Notions

- https://doc.rust-lang.org/book/ch13-02-iterators.html

*/
use highest::*;

fn main() {
	let expected = [30, 500, 20, 70];
	let n = Numbers::new(&expected);
	println!("{:?}", n.list());

	println!("{:?}", n.highest());
	println!("{:?}", n.latest());
	println!("{:?}", n.highest_three());
}

#[cfg(test)]
mod tests {
	use super::*;

	#[test]
	fn test_list() {
		let expected = [30, 50, 20, 70];
		let n = Numbers::new(&expected);
		assert_eq!(n.list(), &expected);
	}

	#[test]
	fn test_latest() {
		let n = Numbers::new(&[100, 0, 90, 30]);
		let f = Numbers::new(&[]);
		assert_eq!(n.latest(), Some(30));
		assert!(
			f.latest().is_none(),
			"It should have been None, {:?}",
			f.latest()
		);
	}

	#[test]
	fn test_highest() {
		let n = Numbers::new(&[40, 100, 70]);
		let f = Numbers::new(&[]);
		assert_eq!(n.highest(), Some(100));
		assert!(
			f.highest().is_none(),
			"It should have been None, {:?}",
			f.highest()
		);
	}

	#[test]
	fn test_highest_three() {
		let e = Numbers::new(&[10, 30, 90, 30, 100, 20, 10, 0, 30, 40, 40, 70, 70]);
		let f = Numbers::new(&[40, 20, 40, 30]);
		let g = Numbers::new(&[30, 70]);
		let h = Numbers::new(&[40]);
		let i = Numbers::new(&[]);
		let j = Numbers::new(&[20, 10, 30]);
		assert_eq!(e.highest_three(), vec![100, 90, 70]);
		assert_eq!(f.highest_three(), vec![40, 40, 30]);
		assert_eq!(g.highest_three(), vec![70, 30]);
		assert_eq!(h.highest_three(), vec![40]);
		assert!(i.highest_three().is_empty());
		assert_eq!(j.highest_three(), vec![30, 20, 10]);
	}
}
