// Create a function call `edit_distance` that calculates the minimum
// number of changes (insertion, deletions and substitutions) that
// need to be made to a string `source` to arrive to another `target`
// string

// For more information and examples https://en.wikipedia.org/wiki/Edit_distance

// pub fn edit_distance(source: &str, target: &str) -> usize {
// 	let src = source.chars().collect::<Vec<_>>();
// 	let tar = target.chars().collect::<Vec<_>>();
// 	let source_len = src.len() + 1;
// 	let target_len = tar.len() + 1;

// 	if source_len == 0 {
// 		return target_len;
// 	}
// 	if target_len == 0 {
// 		return source_len;
// 	}

// 	let mut matrix = vec![vec![0; source_len]; target_len];

// 	for i in 1..target_len {
// 		matrix[i][0] = i
// 	}
// 	for j in 1..source_len {
// 		matrix[0][j] = j
// 	}

// 	for i in 1..target_len {
// 		for j in 1..source_len {
// 			let x = if src[j - 1] == tar[i - 1] {
// 				matrix[i - 1][j - 1]
// 			} else {
// 				1 + std::cmp::min(
// 					std::cmp::min(matrix[i][j - 1], matrix[i - 1][j]),
// 					matrix[i - 1][j - 1],
// 				)
// 			};
// 			matrix[i][j] = x;
// 		}
// 	}
// 	matrix[target_len - 1][source_len - 1]
// }

use edit_distance::edit_distance;

#[allow(dead_code)]
fn main() {
	let source = "alignment";
	let target = "assignment";
	println!(
		"It's necessary to make {} change(s) to {}, to get {}",
		edit_distance(source, target),
		source,
		target
	);
}

#[cfg(test)]
mod test {
	use super::*;

	#[test]
	fn test_distance() {
		assert_eq!(edit_distance("gumbo", "gambol"), 2);
		assert_eq!(edit_distance("kitten", "sitting"), 3);
		assert_eq!(edit_distance("rosettacode", "raisethysword"), 8);
	}
}
