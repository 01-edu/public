/*
## tic_tac_toe

### Instructions

You must create a tic tac toe checker.

Create the following functions:

- `tic_tac_toe` that receives a table of vectors (Vec<Vec<&str>>) and returns a string : `player O won` or `player X won` or `Tie`
- `diagonals` that will receive a player and a table. It should return a boolean, this must return true if all the diagonals are completed by the player
- `horizontal` that will receive a player and a table. It should return a boolean, this must return true if one of the horizontal lines are completed by the player
- `vertical` that will receive a player and a table. It should return a boolean, this must return true if one of the vertical lines are completed by the player

### Example


```

### Notions

- https://doc.rust-lang.org/book/ch04-02-references-and-borrowing.html

*/

// fn check(player: &str, table: &Vec<Vec<&str>>) -> bool {
//     diagonals(player, &table) || horizontal(player, &table) || vertical(player, &table)
// }

// fn tic_tac_toe(table: Vec<Vec<&str>>) -> String {
//     let player1 = "X";
//     let player2 = "O";
//     if check(player2, &table) {
//         return format!("player {} won", player2);
//     } else if check(player1, &table) {
//         return format!("player {} won", player1);
//     } else {
//         return "Tie".to_string();
//     }
// }

// fn diagonals(player: &str, table: &Vec<Vec<&str>>) -> bool {
//     let mut count = 0;
//     let mut count_inv = 0;
//     for (i, v) in table.iter().enumerate() {
//         if v.get(i).unwrap() == &player {
//             count += 1;
//         }
//         if v.get((table.len() - 1) - i).unwrap() == &player {
//             count_inv += 1;
//         }
//     }
//     return count == table.len() || count_inv == table.len();
// }

// // not good will change the solution
// fn horizontal(player: &str, table: &Vec<Vec<&str>>) -> bool {
//     let mut count = 0;
//     for v in table.iter() {
//         for value in v.iter() {
//             if value == &player {
//                 count += 1
//             }
//         }
//         if count == table.len() {
//             return true;
//         }
//         count = 0;
//     }
//     return false;
// }

// // not good will change the solution
// fn vertical(player: &str, table: &Vec<Vec<&str>>) -> bool {
//     let mut count = 0;
//     for (i, _) in table.iter().enumerate() {
//         for v in table.iter() {
//             if v.get(i).unwrap() == &player {
//                 count += 1;
//             }
//         }
//         if count == table.len() {
//             return true;
//         }
//         count = 0;
//     }
//     return false;
// }
use tic_tac_toe::*;

fn main() {
	println!(
		"{:?}",
		tic_tac_toe(vec![
			vec!["O", "X", "O"],
			vec!["O", "P", "X"],
			vec!["X", "#", "X"]
		])
	);
	// "Tie"
	println!(
		"{:?}",
		tic_tac_toe(vec![
			vec!["X", "O", "O"],
			vec!["X", "O", "O"],
			vec!["#", "O", "X"]
		])
	);
	// "player O won"

	let dig = vec![
		vec!["O", "O", "X"],
		vec!["O", "X", "O"],
		vec!["X", "#", "X"],
	];

	println!("{:?}", tic_tac_toe(dig));
	// "player X won"
}
#[cfg(test)]
mod tests {
	use super::*;

	struct Test<'a> {
		player: &'a str,
		table: Vec<Vec<&'a str>>,
		result: &'a str,
	}

	impl Test<'_> {
		fn init_horizontal() -> Vec<Test<'static>> {
			vec![
				Test {
					player: "O",
					table: vec![
						vec!["O", "O", "O"],
						vec!["X", "X", "O"],
						vec!["O", "#", "X"],
					],
					result: "player O won",
				},
				Test {
					player: "O",
					table: vec![
						vec!["X", "X", "O"],
						vec!["O", "O", "O"],
						vec!["O", "#", "X"],
					],
					result: "player O won",
				},
				Test {
					player: "X",
					table: vec![
						vec!["O", "X", "O"],
						vec!["O", "#", "O"],
						vec!["X", "X", "X"],
					],
					result: "player X won",
				},
				Test {
					player: "X",
					table: vec![
						vec!["O", "X", "O", "O"],
						vec!["X", "X", "X", "X"],
						vec!["X", "#", "O", "X"],
						vec!["X", "X", "O", "O"],
					],
					result: "player X won",
				},
			]
		}
		fn init_tie() -> Vec<Test<'static>> {
			vec![
				Test {
					player: "none",
					table: vec![
						vec!["O", "X", "O"],
						vec!["O", "X", "O"],
						vec!["X", "#", "X"],
					],
					result: "Tie",
				},
				Test {
					player: "none",
					table: vec![
						vec!["O", "X", "O"],
						vec!["X", "X", "O"],
						vec!["X", "#", "X"],
					],
					result: "Tie",
				},
				Test {
					player: "none",
					table: vec![
						vec!["O", "X", "O", "O"],
						vec!["X", "O", "X", "X"],
						vec!["X", "#", "X", "X"],
						vec!["X", "X", "O", "O"],
					],
					result: "Tie",
				},
			]
		}
		fn init_vertical() -> Vec<Test<'static>> {
			vec![
				Test {
					player: "O",
					table: vec![
						vec!["O", "X", "O"],
						vec!["O", "X", "O"],
						vec!["O", "#", "X"],
					],
					result: "player O won",
				},
				Test {
					player: "O",
					table: vec![
						vec!["X", "O", "O"],
						vec!["X", "O", "O"],
						vec!["#", "O", "X"],
					],
					result: "player O won",
				},
				Test {
					player: "X",
					table: vec![
						vec!["O", "X", "X"],
						vec!["O", "X", "X"],
						vec!["X", "#", "X"],
					],
					result: "player X won",
				},
			]
		}
		fn init_diagonals() -> Vec<Test<'static>> {
			vec![
				Test {
					player: "X",
					table: vec![
						vec!["O", "O", "X"],
						vec!["O", "X", "O"],
						vec!["X", "#", "X"],
					],
					result: "player X won",
				},
				Test {
					player: "O",
					table: vec![
						vec!["O", "X", "O"],
						vec!["X", "O", "O"],
						vec!["X", "#", "O"],
					],
					result: "player O won",
				},
				Test {
					player: "O",
					table: vec![
						vec!["O", "X", "O", "O"],
						vec!["X", "O", "X", "O"],
						vec!["X", "#", "O", "X"],
						vec!["X", "X", "O", "O"],
					],
					result: "player O won",
				},
			]
		}
	}

	#[test]
	fn test_diagonals() {
		let new_tests = Test::init_diagonals();
		for v in new_tests {
			assert_eq!(diagonals(v.player, &v.table), true)
		}

		assert_eq!(
			diagonals(
				"O",
				&vec![
					vec!["O", "X", "X"],
					vec!["O", "X", "X"],
					vec!["X", "#", "X"]
				]
			),
			false
		);
	}

	#[test]
	fn test_horizontal() {
		let new_tests = Test::init_horizontal();
		for v in new_tests {
			assert_eq!(horizontal(v.player, &v.table), true)
		}

		assert_eq!(
			horizontal(
				"O",
				&vec![
					vec!["O", "X", "X"],
					vec!["O", "O", "X"],
					vec!["X", "#", "O"]
				]
			),
			false
		);
	}

	#[test]
	fn test_vertical() {
		let new_tests = Test::init_vertical();
		for v in new_tests {
			assert_eq!(vertical(v.player, &v.table), true)
		}

		assert_eq!(
			vertical(
				"O",
				&vec![
					vec!["O", "X", "X"],
					vec!["O", "O", "X"],
					vec!["X", "#", "O"]
				]
			),
			false
		);
	}

	#[test]
	fn test_tic_tac_toe() {
		let mut new_tests = Test::init_diagonals();
		new_tests.append(&mut Test::init_horizontal());
		new_tests.append(&mut Test::init_vertical());
		new_tests.append(&mut Test::init_tie());

		for v in new_tests {
			assert_eq!(tic_tac_toe(v.table), v.result.to_string());
		}
	}
}
