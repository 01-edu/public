use panic::*;
use std::fs::{self, File};

fn main() {
    let filename = "created.txt";
    File::create(filename).unwrap();
    let a = open_file(filename);
    println!("{:?}", a);
    fs::remove_file(filename).unwrap();
}

#[cfg(test)]
mod tests {
	use super::*;

	#[test]
	#[should_panic]
	fn test_opening() {
		open_file("file.txt");
	}

	#[test]
	fn test_opening_existing() {
		let filename = "created.txt";
		File::create(filename).unwrap();
		open_file(filename);
		fs::remove_file(filename).unwrap();
	}
}
