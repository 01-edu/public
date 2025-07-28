## commits_stats

### Instructions:

In this exercise, you will be provided with a json file `commits.json`([download](./attachments/commits.json)) with data corresponding to git commits in GitHub (extracted using the GitHub rest API). Your objective is to extract the relevant data and place it in a struct called `CommitData`.

Create two functions:
- `commits_per_author`: which returns a hash map with the number of commits per author. The auditors will be identified by their GitHub login.
- `commits_per_week`: which returns a hash map with the number of commits per week.

> A week is represented by the year followed by the number of the week. For example, January 1, 2020 is in week 1 of 2020 and will be represented by a `String` with the form `"2020-W1"`.

Note: You must use the `json` crate `v0.12.4` for the functions' respective arguments.

### Expected functions

```rust
pub fn commits_per_week(data: &json::JsonValue) -> HashMap<String, u32> {
	todo!()
}

pub fn commits_per_author(data: &json::JsonValue) -> HashMap<String, u32> {
	todo!()
}
```

### Usage

Here is a possible test for your function:

```rust
use commits_stats::{commits_per_week, commits_per_author};

fn main() {
	let contents = include_str!("commits.json").unwrap();
	let serialized = json::parse(&contents).unwrap();
	println!("{:?}", commits_per_week(&serialized));
	println!("{:?}", commits_per_author(&serialized));
}
```

And its output:

```console
$ cargo run
{"2020-W44": 5, "2020-W36": 1, "2020-W31": 1, ... ,"2020-W45": 4, "2020-W46": 4}
{"homembaixinho": 2, "mwenzkowski": 3, ... ,"tamirzb": 1, "paul-ri": 2, "RPigott": 1}
$
```

### Notions:

- [chrono](https://docs.rs/chrono/0.4.41/chrono)
- [json](https://docs.rs/json/0.12.4/json/index.html)
