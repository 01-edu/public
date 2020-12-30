## commits_stats

### Instructions:

In this exercise you will be provided with a json file with data corresponding to git commits in github (extracted using the github rest api) your job is to extract the relevant data and place it in a struct called `CommitData` to get the following information:

- Number of commits per author (identified by the github login)
- And the number of commits per author

Create two functions:

- `commits_per_author` which returns a HashMap where the key is the github login and the value is the number of commits that author made.

- `commits_per_date` which returns the number of commits per week as a HashMap where the key is the `week` and the key is the number of commits in that week.

Note: A week is represented by the a year followed by the number of the week. For example: January 1, 2020 is in week 1 of 2020 an will be
represented by a String with the form "2020-W1"

### Notions:

- https://docs.rs/chrono/0.4.19/chrono/#modules

- https://serde.rs/

### Expected Functions and Structure

```rust
#[derive(Serialize, Deserialize, Debug)]
pub struct CommitData {
...
}

fn commits_per_author(data: &Vec<CommitData>) -> HashMap<&str, u32> {
}

fn commits_per_week(data: &Vec<CommitData>) -> HashMap<String, u32> {
}
```

### Usage

Here is a program to test your function.

```rust
fn main() {
	let contents = fs::read_to_string("commits.json").unwrap();
	let serialized: Vec<CommitData> = serde_json::from_str(&contents).unwrap();
	println!("{:?}", commits_per_week(&serialized));
	println!("{:?}", commits_per_author(&serialized));
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
{"2020-W50": 2, .. , "2020-W43": 1, "2020-W44": 5, "2020-W40": 2, "2020-W36": 1}
{"ifreund": 1, .. , "tamirzb": 1, "paul-ri": 2, "emersion": 10, "psnszsn": 1, "mwenzkowski": 3, "Xyene": 7}
student@ubuntu:~/[[ROOT]]/test$
```
