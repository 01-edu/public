## tic_tac_toe

### Instructions

You must create some functions for a tic-tac-toe checker.

Create a function named `tic_tac_toe`, which receives a tic-tac-toe table. It should return the appropriate string: `"player O won"`, `"player X won"` or `"tie"`.


Also create the following functions, which each accept a player and a table. These functions should return `true` if the player has completed one of the diagonals, rows or columns:

### Expected functions
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

Here is a program to test your `tic_tac_toe`. You'll need to test the other functions yourself. But they'll probably be useful for implementing your `tic_tac_toe` checker.

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

    println!(
        "{:?}",
        tic_tac_toe(vec![
            vec!["X", "O", "O"],
            vec!["X", "O", "O"],
            vec!["#", "O", "X"]
        ])
    );

    let dig = vec![
            vec!["O", "O", "X"],
            vec!["O", "X", "O"],
            vec!["X", "#", "X"]
        ];

    println!("{:?}",tic_tac_toe(dig));
}
```

And its output

```console
$ cargo run
"tie"
"player O won"
"player X won"
$
```

### Notions

- [references and borrowing](https://doc.rust-lang.org/book/ch04-02-references-and-borrowing.html)
