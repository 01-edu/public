## tic_tac_toe

### Instructions

You must create a tic tac toe checker.

Create the following functions:

- `tic_tac_toe` that receives a table of vectors (Vec<Vec<&str>>) and returns a string : `player O won` or `player X won` or `Tie`
- `diagonals` that will receive a player and a table. It should return a boolean, this must return true if all the diagonals are completed by the player
- `horizontal` that will receive a player and a table. It should return a boolean, this must return true if one of the horizontal lines are completed by the player
- `vertical` that will receive a player and a table. It should return a boolean, this must return true if one of the vertical lines are completed by the player

### Notions

- https://doc.rust-lang.org/book/ch04-02-references-and-borrowing.html

### Expected Functions

```rust
fn tic_tac_toe(table: Vec<Vec<&str>>) -> String {
}

fn diagonals(player: &str, table: &Vec<Vec<&str>>) -> bool {
}

fn horizontal(player: &str, table: &Vec<Vec<&str>>) -> bool {
}

fn vertical(player: &str, table: &Vec<Vec<&str>>) -> bool {
}
```

### Usage

Here is a program to test your function

```rust
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
            vec!["X", "#", "X"]
        ];

    println!("{:?}",tic_tac_toe(dig));
    // "player X won"
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
"Tie"
"player O won"
"player X won"
student@ubuntu:~/[[ROOT]]/test$
```
