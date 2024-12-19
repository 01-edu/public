## tic_tac_toe

### Instructions

You must create some functions for a tic-tac-toe checker.

Create a function named `tic_tac_toe`, which receives a tic-tac-toe table. It should return the appropriate string: `"player O won"`, `"player X won"` or `"tie"`.

Also create the following functions, which each accept a player and a table. These functions should return `true` if the player has completed one of the diagonals, rows or columns:

### Expected functions

```rust
pub fn tic_tac_toe(table: [[char; 3]; 3]) -> String {
}

pub fn diagonals(player: char, table: [[char; 3]; 3]) -> bool {
}

pub fn horizontal(player: char, table: [[char; 3]; 3]) -> bool {
}

pub fn vertical(player: char, table: [[char; 3]; 3]) -> bool {
}
```

### Usage

Here is a program to test your `tic_tac_toe`. You'll need to test the other functions yourself.

```rust
use tic_tac_toe::*;

fn main() {
    println!(
        "{}",
        tic_tac_toe([['O', 'X', 'O'], ['O', 'P', 'X'], ['X', '#', 'X']])
    );
    // tie
    println!(
        "{}",
        tic_tac_toe([['X', 'O', 'O'], ['X', 'O', 'O'], ['#', 'O', 'X']])
    );
    // player O won

    let diag = [['O', 'O', 'X'], ['O', 'X', 'O'], ['X', '#', 'X']];

    println!("{}", tic_tac_toe(diag));
    // player X won
}
```

And its output

```console
$ cargo run
tie
player O won
player X won
$
```

### Notions

- [references and borrowing](https://doc.rust-lang.org/book/ch04-02-references-and-borrowing.html)
