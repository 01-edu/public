## borrow box

### Instructions

In this exercise, a **CRUD** functionality will have to be created. Therefore the following functions will have to be defined :

- `new`, which receives two players and initializes them with a name and a score. This functions should
  return the structure wrapped in a Box.

- `read_winner`, which returns a tuple with the name and the score of the player who is winning.
  In case none of the players are winning, it should return the same tuple with the string "Same score! tied" and the tied score.

  - `update_score`, which receives the name of a player.
    This function should increment the score of the player. The score should only be increased if it does not pass the `nbr_of_games`.
    Example: if the nbr_of_games is 3 then the winner of the game should be the one who is the best out of three games. So if one play as 2 wins then
    he is the winner and the function should not increase the score anymore for either players.

- `delete`, which takes the ownership of the boxed game and returning a string : "Game deleted: id -> 0".

### Notions

- [Box Pointers](https://doc.rust-lang.org/book/ch15-01-box.html)

### Expected Functions

```rust
#[derive(Debug, Clone, Eq, PartialEq)]
pub struct Game {
    // expected public fields
}
impl Game {
    pub fn new(i: u32, pl1: String, pl2: String, n: u16) -> Box<Game> {

    }
    pub fn read_winner(&self) -> (String, u16) {

    }
    pub fn update_score(&mut self, user_name: String) {

    }
    pub fn delete(self) -> String {

    }
}
```

### Usage

Here is a program to test your functions,

```rust
use borrow_box::*;

fn main() {
    let mut game = Game::new(0, String::from("Joao"), String::from("Susana"), 5);
    println!("{:?}", game.read_winner());
    // output : ("Same score! tied", 0)

    game.update_score(String::from("Joao"));
    game.update_score(String::from("Joao"));
    game.update_score(String::from("Susana"));
    game.update_score(String::from("Susana"));
    println!("{:?}", game.read_winner());
    // output : ("Same score! tied", 2)

    game.update_score(String::from("Joao"));
    // this one will not count because it already 5 games played, the nbr_of_games
    game.update_score(String::from("Susana"));

    println!("{:?}", game.read_winner());
    // output : ("Joao", 3)

    println!("{:?}", game.delete());
    // output : "game deleted: id -> 0"

    // game.read_winner();
    // this will give error
    // because the game was dropped, no longer exists on the heap
}
```

And its output:

```console
student@ubuntu:~/borrow_box/test$ cargo run
("Same score! tied", 0)
("Same score! tied", 2)
("Joao", 3)
"game deleted: id -> 0"
student@ubuntu:~/borrow_box/test$
```
