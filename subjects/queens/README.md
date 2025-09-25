## queens

### Instructions

In a chess game, a queen can attack pieces which are on the same rank (row), file (column), or diagonal.

The purpose of this exercise is to find out if two queens can attack each other.

The position of a chess piece on a chessboard will be represented by the struct `ChessPosition`. You must implement the associated function `new` which will return the position if it is valid, otherwise it will return `None`.

> Remember, chessboards have 8 files and 8 ranks (0 to 7).

You will create the `Queen` struct with the associate function `can_attack`, which will return `true` if the queens can attack each other or not. You also need to implement the function `new` which creates a new `Queen` with a `ChessPosition`.

### Example

If the white queen is at (2, 3) and the black queen is at (5, 6), then the set-up would be like the below. In this case, the two queens can attack each other because both pieces share a diagonal:

```
_ _ _ _ _ _ _ _
_ _ _ _ _ _ _ _
_ _ _ W _ _ _ _
_ _ _ _ _ _ _ _
_ _ _ _ _ _ _ _
_ _ _ _ _ _ B _
_ _ _ _ _ _ _ _
_ _ _ _ _ _ _ _
```

### Expected Function and Structures

```rust
#[derive(Debug, Clone, Copy)]
pub struct ChessPosition {
    pub rank: usize,
    pub file: usize,
}

impl ChessPosition {
    pub fn new(rank: usize, file: usize) -> Option<Self> {
        todo!()
    }
}

#[derive(Debug, Clone, Copy)]
pub struct Queen {
    pub position: ChessPosition,
}

impl Queen {
    pub fn new(position: ChessPosition) -> Self {
        todo!()
    }

    pub fn can_attack(self, other: Self) -> bool {
        todo!()
    }
}
```

### Usage

Here is a possible program to test your function:

```rust
use queens::*;

fn main() {
    let white_queen = Queen::new(ChessPosition::new(2, 2).unwrap());
    let black_queen = Queen::new(ChessPosition::new(0, 4).unwrap());

    println!(
        "Is it possible for the queens to attack each other? {}",
        white_queen.can_attack(black_queen)
    );

    let white_queen = Queen::new(ChessPosition::new(1, 2).unwrap());
    let black_queen = Queen::new(ChessPosition::new(0, 4).unwrap());

    println!(
        "Is it possible for the queens to attack each other? {}",
        white_queen.can_attack(black_queen)
    );
}
```

And its output:

```console
$ cargo run
Is it possible for the queens to attack each other? true
Is it possible for the queens to attack each other? false
$
```
