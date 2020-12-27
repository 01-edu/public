## capitalizing

### Instructions

Complete the `capitalize_first` function that turns the first letter of a string uppercase.

Complete the `title_case` function that turns the first letter of each word in a string uppercase.

Complete the `change_case` function that turns the uppercase letters of a string into lowercase and the lowercase letters into uppercase.

### Expected Functions

```rust
pub fn capitalize_first(input: &str) -> String {
}

pub fn title_case(input: &str) -> String {
}

pub fn change_case(input: &str) -> String {
}
```

### Usage

Here is a program to test your function.

```rust
fn main() {
    println!("{}", capitalize_first("joe is missing"));
    println!("{}", title_case("jill is leaving A"));
    println!("{}",change_case("heLLo THere"));
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
Joe is missing
Jill Is Leaving A
HEllO thERE
student@ubuntu:~/[[ROOT]]/test$
```
