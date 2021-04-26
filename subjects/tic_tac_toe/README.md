## tic_tac_toe

### Instructions

You must create a `tic tac toe` checker.

Create the following functions:

- `tic_tac_toe` which receives:
  - a table of vectors (Vec<Vec<&str>>).
  - It should return a String `player O won` or `player X won` or `Tie`.
- `diagonals` which will receive:
  - a player and a table.
  - It should return a boolean, this must return `true` if one of the diagonals are completed by the player.
- `horizontal` which will receive:
  - a player and a table.
  - It should return a boolean, this must return `true` if one of the horizontal lines are completed by the player.
- `vertical` which will receive:
  - a player and a table.
  - It should return a boolean, this must return `true` if one of the vertical lines are completed by the player.

### Notions

- [references and borrowing](https://doc.rust-lang.org/book/ch04-02-references-and-borrowing.html)

### Expected Functions

```rust
pub fn tic_tac_toe(table: Vec<Vec<&str>>) -> String {
}

pub fn diagonals(player: &str, table: &Vec<Vec<&str>>) -> bool {
}

pub fn horizontal(player: &str, table: &Vec<Vec<&str>>) -> bool {
}

pub fn vertical(player: &str, table: &Vec<Vec<&str>>) -> bool {
}
```

### Usage

Here is a program to test your function

```rust
use tic_tac_toe::*;

fn main() {
    println!(
        "{:?}",
        tic_tac_toe(vec![
            vec!["O", "X", "O"],
            vec!["O", "O", "X"],
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
            vec!["X", "#", "X"]
        ];

    println!("{:?}",tic_tac_toe(dig));
    // "player X won"
}
```

And its output

```console
student@ubuntu:~/tic_tac_toe/test$ cargo run
"Tie"
"player O won"
"player X won"
student@ubuntu:~/tic_tac_toe/test$
```
