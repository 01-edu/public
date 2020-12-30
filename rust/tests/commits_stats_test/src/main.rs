// # Instructions:

// In this exercise you will be provided with a json file with data
// corresponding to git commits in github (extracted using the github
// rest api) your job is to extract the relevant data and place it in
// a struct called `CommitData` to get the following information:

// 1. Number of commits per author (identified by the github login)
// 2. And the number of commits per author

// Create two functions:
// fn commits_per_author(data: &Vec<CommitData>) -> HashMap<&str, u32>
// fn commits_per_date(data: &Vec<CommitData>) -> HashMap<String, u32>
// A week is represented by the a year followed by the number of the
// week for example January 1, 2020 is in week 1 of 2020 an will be
// represented by a String with the form "2020-W1"

// # Notions:
// https://docs.rs/chrono/0.4.19/chrono/#modules
// https://serde.rs/

#[allow(unused_imports)]
use commits_stats::CommitData;
#[allow(unused_imports)]
use commits_stats::*;
#[allow(unused_imports)]
use std::fs;

fn main() {
	let contents = fs::read_to_string("commits.json").unwrap();
	let serialized: Vec<CommitData> = serde_json::from_str(&contents).unwrap();
	println!("{:?}", commits_per_week(&serialized));
	println!("{:?}", commits_per_author(&serialized));
}

#[cfg(test)]
mod tests {
	use super::*;
	use std::collections::HashMap;

	fn test_setup() -> Vec<CommitData> {
		let contents = fs::read_to_string("commits.json").unwrap();
		let serialized: Vec<CommitData> = serde_json::from_str(&contents).unwrap();
		serialized
	}

	#[test]
	fn test_commits_per_week() {
		let serialized = test_setup();
		let commits_per_week = commits_per_week(&serialized);
		println!("{:#?}", &commits_per_week);
		let date = [
			"2020-W47".to_string(),
			"2020-W43".to_string(),
			"2020-W36".to_string(),
			"2020-W50".to_string(),
			"2020-W40".to_string(),
			"2020-W44".to_string(),
			"2020-W46".to_string(),
			"2020-W31".to_string(),
			"2020-W45".to_string(),
			"2020-W49".to_string(),
		];

		let mut com_per_week = HashMap::new();
		let commits = [3, 1, 1, 2, 2, 5, 4, 1, 4, 7];

		for i in 0..date.len() {
			com_per_week.insert(date[i].clone(), commits[i].clone());
		}

		assert_eq!(com_per_week, commits_per_week);
	}

	#[test]
	fn test_commits_per_author() {
		let serialized = test_setup();
		let logins = [
			"RPigott",
			"RedSoxFan",
			"Xyene",
			"paul-ri",
			"JayceFayne",
			"mwenzkowski",
			"psnszsn",
			"emersion",
			"tamirzb",
			"ifreund",
			"homembaixinho",
		];
		let commits = [1, 1, 7, 2, 1, 3, 1, 10, 1, 1, 2];
		let mut expected = HashMap::new();

		for i in 0..logins.len() {
			expected.insert(logins[i], commits[i].to_owned());
		}

		let commits_per_author = commits_per_author(&serialized);
		println!("{:#?}", &commits_per_author);
		assert_eq!(expected, commits_per_author);
	}
}
