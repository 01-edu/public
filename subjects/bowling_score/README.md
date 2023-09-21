## bowling_score

### Instructions

Create a library that implements a `struct` to calculate the score of a player in the game of [Bowling](https://en.wikipedia.org/wiki/Bowling).

The methods of `BowlingGame` will be:

- `roll`: It will take the number of knocked pins and return `Ok(())` if the roll is valid or an `Err(Error::...)` if the roll is invalid.
- If the pins knocked out are more than the available ones the error will be `NotEnoughPinsLeft`.
- If the roll happens after the game is complete the error will be `GameComplete`.
- `score`: It will calculate the score of the player and return `Some(score)` in case of success or `None` if there is still some rolls missing to complete the game.

#### The basic rules of bowling

- The game is divided in 10 frames
- At the start of each frame 10 pins are placed at the end of the lane
- Each frame has at maximum 2 rolls
- `Strike`: You knock 10 pins on the first roll, the game moves directly to the next frame since there is no more pins to knock in this one. You will also get extra points based on the next two rolls.
- `Spare`: On the second frame you knock all the pins left. You move to the next frame as you would normally but you will get extra points based on the next roll.

Last frame scenario:

- If on the 10th frame you make a `Strike` you will get two extra rolls (2 fill balls).
- If on the 10th frame you make a `Spare` you will get one extra roll (1 fill ball).

#### The scoring system

- Each pin knocked is one point.
- If you do a strike (10 pins knocked on first roll of a frame) it will count as 10 + pins knocked in the next two rolls.
- If you do a spare (10 pins knocked in between the two rolls of the frame) it will count as 10 + pins knocked in the next roll.

Last frame scenario:

The purpose of the filling balls is only to allow the calculation of the 10th frame, so if you do a strike on the 10th and then two more strikes the score will be `30`, because the last two are fill balls.

> 

### Expected Function

```rust
#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    NotEnoughPinsLeft,
    GameComplete,
}

pub struct BowlingGame {
    ...
}

impl BowlingGame {
    pub fn new() -> Self {
        unimplemented!();
    }

    pub fn roll(&mut self, pins: u16) -> Result<(), Error> {
        unimplemented!();
    }

    pub fn score(&mut self) -> Some(u16) {
        unimplemented!();
    }
}
```

### Usage

Here is a possible program to test your function,

```rust
fn main() -> Result<(), Error> {
    let mut game = BowlingGame::new();
    game.roll(0)?; // frame 1
    game.roll(10)?; // frame 1: spare
    game.roll(10)?; // frame 2: strike
    game.roll(5)?; // frame 3
    game.roll(5)?; // frame 3: spare
    game.roll(10)?; // frame 4: strike
    game.roll(10)?; // frame 5: strike
    game.roll(10)?; // frame 6: strike
    game.roll(10)?; // frame 7: strike
    game.roll(10)?; // frame 8: strike
    game.roll(10)?; // frame 9: strike
    game.roll(10)?; // frame 10: strike
    game.roll(2)?; // fill ball 1
    game.roll(8)?; // fill ball 2
    println!("{:?}", game.score());

    let mut perfect_game = BowlingGame::new();
    for _ in 0..12 {
        perfect_game.roll(10)?;
    }
    println!("{:?}", perfect_game.score());
    Ok(())
}
```

And its output:

```console
$ cargo run
Some(252)
Some(300)
$
```
