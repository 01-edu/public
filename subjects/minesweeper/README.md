## minesweeper

### Instructions

Create a function that takes a minesweeper board as an array of strings and return the board solved.

Minesweeper is a very old game where some mines are placed in a board and you should calculate how many mines are touching every free field and write the count in the respective place.

> We will only test your function with empty and valid boards.

### Expected Function

```rust
pub fn solve_board(minefield: &[&str]) -> Vec<String> {
}
```

### Usage

Here is a possible program to test your function,

```rust
fn main() {
    println!("{:?}", solve_board(&[]));
    println!("{:?}", solve_board(&[""]));
    println!("{:?}", solve_board(&["***"]));
    println!("{:#?}", solve_board(&["   ", " * ", "   ",]));
    println!("{:#?}", solve_board(&["*  ", "   ", "  *",]));
}
```

And its output:

```console
$ cargo run
[]
[""]
["***"]
[
    "111",
    "1*1",
    "111",
]
[
    "*1 ",
    "121",
    " 1*",
]
$
```
