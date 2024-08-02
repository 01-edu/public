## borrow box

### Instructions

Game time.

You will implement some **CRUD** functionality for a game session. You will need to implement the `GameSession` structure with the following associated functions:

- `new`: which initializes a game session state with player names and some other information. This function returns the structure wrapped in a `Box`.

- `read_winner`: which returns a tuple with the name and score of the player who is currently winning. In the case that no player is winning, it should return the same tuple with the string `"Same score! tied"` and the tied score.

- `update_score`: which receives the name of a player, and increments their score. This function should **do nothing** if the the game session is already finished or if the name received doesn't match any player.

- `delete`: which takes ownership of the boxed game session and returns a string: `"game deleted: id -> 0"`, where `0` is the id of the `GameSession`.

> Examples for `nb_games`:
>
> `nb_games` = **5** --> the maximum of games will be **5** and if some player has a score of 3, the game is finished (not enough games for the other player to win).\
> `nb_games` = **11** --> the maximum of games will be **11** and if some player has a score of 6, the game is finished (not enough games for the other player to win).

### Expected Functions

```rust
#[derive(Debug, Clone, Eq, PartialEq)]
pub struct GameSession {
    pub id: u32,
    pub p1: (String, u16),
    pub p2: (String, u16),
    pub nb_games: u16
}

impl GameSession {
    pub fn new(id: u32, p1_name: String, p2_name: String, nb_games: u16) -> Box<GameSession> {

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
    let mut game = GameSession::new(0, String::from("Joao"), String::from("Susana"), 5);
    println!("{:?}", game.read_winner());
    // output : ("Same score! tied", 0)

    game.update_score(String::from("Joao"));
    game.update_score(String::from("Joao"));
    game.update_score(String::from("Susana"));
    game.update_score(String::from("Susana"));
    println!("{:?}", game.read_winner());
    // output : ("Same score! tied", 2)

    game.update_score(String::from("Joao"));
    // this one will not count because it already 5 games played, the nb_games
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
$ cargo run
("Same score! tied", 0)
("Same score! tied", 2)
("Joao", 3)
"game deleted: id -> 0"
$
```

### Notions

- [Box Pointers](https://doc.rust-lang.org/book/ch15-01-box.html)
