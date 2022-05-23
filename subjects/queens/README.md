## queens

### Instructions

In a chess game, a queen can attack pieces which are on the same rank, file, or diagonal.

The purpose of this exercise is to find out if two queens can attack each other using the same rules.

The chess board will be represented by the struct `ChessPosition`. You must implement the function `new` 
that allows you to verify if the position is valid or not. If the position is valid the function will return that 
position, otherwise it will return `None`.

So, given the position of the two queens on a chess board, you will have to
implement the function `can_attack` in the given struct `Queen` with
the purpose of finding out whether the two queens can attack each other or not. You also need to implement the function `new` 
that allows you to create a new `Queen` given a ChessPosition.

For example, if the white queen is at (2, 3) and the black queen is at (5, 6),
then the set-up would be like so:

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

In this case, the two queens can attack each other because both pieces share a diagonal.

### Expected Function and Structures

```rust
#[derive(Debug)]
pub struct ChessPosition {
    pub rank: i32,
    pub file: i32,
}

#[derive(Debug)]
pub struct Queen {
    pub position: ChessPosition,
}

impl ChessPosition {
    pub fn new(rank: i32, file: i32) -> Option<Self> {

    }
}

impl Queen {
    pub fn new(position: ChessPosition) -> Self {

    }

    pub fn can_attack(&self, other: &Queen) -> bool {

    }
}
```

### Usage

Here is a possible program to test your function :

```rust
fn main() {
    let white_queen = Queen::new(ChessPosition::new(2, 2).unwrap());
    let black_queen = Queen::new(ChessPosition::new(0, 4).unwrap());

    println!(
        "Is it possible for the queens to attack each other? => {}",
        white_queen.can_attack(&black_queen)
    );

    let white_queen = Queen::new(ChessPosition::new(1, 2).unwrap());
    let black_queen = Queen::new(ChessPosition::new(0, 4).unwrap());

    println!(
        "Is it possible for the queens to attack each other? => {}",
        white_queen.can_attack(&black_queen)
    );
}
```

And its output:

```console
$ cargo run
Is it possible for the queens to attack each other? => true
Is it possible for the queens to attack each other? => false
$
```
