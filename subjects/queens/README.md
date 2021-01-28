## queens

### Instructions

In a chess game, a queen can attack pieces which are on the same row,
column, or diagonal.

Your purpose in this exercise is to find out if two queens can attack
each other using the same rules.

The chess board will be represented as an 8 by 8 array.

So, given the position of the two queens on a chess board, you will have to
implement the function `can_attack` in the given struct `Queen` with
the purpose of finding out if the two queens can attack each other or not.

For this to be possible, you will also have to implement the struct `ChessPosition`
with the function `new` that will allow you to verify if the position is valid or not. If the position is valid it will return that position and if it is invalid it will return `None`.

So if you are told that the white queen is at (2, 3) and the black queen is at (5, 6),
then you would know you have got a set-up like so:

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

### Expected Function

```rust
#[derive(Debug)]
pub struct ChessPosition {
    rank: i32,
    file: i32,
}

#[derive(Debug)]
pub struct Queen {
    position: ChessPosition,
}

impl ChessPosition {
    pub fn new(rank: i32, file: i32) -> Option<Self> {}
}

impl Queen {
    pub fn can_attack(&self, other: &Queen) -> bool {}
}
```

### Usage

Here is a possible program to test your function :

```rust
use queens::queens;

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
student@ubuntu:~/[[ROOT]]/test$ cargo run
Is it possible for the queens to attack each other? => true
Is it possible for the queens to attack each other? => false
student@ubuntu:~/[[ROOT]]/test$
```
