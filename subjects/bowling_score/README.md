## bowling_score

### Instructions

Create a library that implements a `struct` to calculate the score of a player in the game of [Bowling](https://en.wikipedia.org/wiki/Bowling).

The methods of `BowlingGame` will be:

- `roll`: It will take the number of pins knocked down and return `Ok(())` if the roll is valid or an `Err(Error::...)` if the roll is invalid.
- If more pins than the ones available are knocked down the error will be `NotEnoughPinsLeft`.
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

- Each pin knocked down is one point.
- If you do a strike (knocking down all 10 pins with the first roll of a frame) it will count as 10 points + the number of pins knocked down in the next two rolls.
- If you do a spare (knocking down 10 pins in total between the two rolls of a frame) it will count as 10 points + the number of pins knocked down in the next roll.

Last frame scenario:

The purpose of the filling balls is solely to allow the calculation of the 10th frame. For example if you score a strike in the 10th frame and then you score two more strikes, your total score will be 30, as the last two rolls are fill balls.

### Expected Function

```rust
#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    NotEnoughPinsLeft,
    GameComplete,
}

pub struct BowlingGame {}

impl BowlingGame {
    pub fn new() -> Self {
        todo!();
    }

    pub fn roll(&mut self, pins: u32) -> Result<(), Error> {
        todo!();
    }

    pub fn score(&self) -> Option<u32> {
        todo!();
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
