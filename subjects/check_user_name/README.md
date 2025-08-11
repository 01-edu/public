## check_user_name

### Instructions

Sometimes it is more desirable to catch the failure of some parts of a program instead of just calling panic.

For this exercise you will have to create a tool that manages users' access levels.

You will have to create an `AccessLevel` enum which could be `Guest`, `Normal`, `Admin`.

You will also have to create a `User` struct which has:

- Fields:
  - `name`: `String`
  - `access_level`: `AccessLevel`
- Associated functions:
  - `new`: which initializes the struct.
  - `send_name`: which takes only `self` as argument and returns an `Option<&str>` with `None` if the user is a `Guest` or the `name` if the `AccessLevel` has any of the other options.
- Other functions:
  - `check_user_name`: which takes a reference to `User`, calls `send_name` and returns a `tuple` with `true` and the user `name` if `send_name` returns the name or `false` and `"ERROR: User is guest"` if not.

### Expected Functions and Data Structures

```rust
pub enum AccessLevel {}

pub struct User {}

impl User {
  pub fn new(name: String, level: AccessLevel) -> User {}
  pub fn send_name(&self) -> Option<&str> {}
}

pub fn check_user_name(user: &User) -> (bool, &str) {}
```

### Usage

Here is a program to test your function:

```rust
use check_user_name::*;

fn main() {
  let user0 = User::new("Didier".to_string(), AccessLevel::Admin);
  println!("{:?}", check_user_name(&user0));

  let user1 = User::new("Mary".to_string(), AccessLevel::Normal);
  println!("{:?}", check_user_name(&user1));

  let user2 = User::new("John".to_string(), AccessLevel::Guest);
  println!("{:?}", check_user_name(&user2));
}
```

And its output:

```console
$ cargo run
(true, "Didier")
(true, "Mary")
(false, "ERROR: User is guest")
$
```
