## borrow box

### Instructions

Game time.

You will implement some **CRUD** functionality for a game session. You will need to implement the `GameSession` structure with the following associated functions:

- `new`: which initializes a game session state with player names and some other information.

- `read_winner`: which returns a tuple referencing the player who is currently winning. In the case that no player is winning, it should return `None`.

- `update_score`: which receives the name of a player, and increments their score. This function should **do nothing** if the the game session is already finished or if the name received doesn't match any player.

- `delete`: which takes ownership of the game session and returns a string: `"game deleted: id -> {id}"`, where `{id}` is the id of the `GameSession`.

> Examples for `nb_games`:
>
> When `nb_games` is 5, the game is best out of 5, and if some player has a score of 3, the game is finished (there aren't enough games for the other player to draw).
> When `nb_games` is 11, the game is best out of 11, and if some player has a score of 6, the game is finished (there aren't enough games for the other player to draw).

### Expected Functions

```rust
#[derive(Debug, Clone, Eq, PartialEq)]
pub struct GameSession {
    pub id: u32,
    pub p1: (String, u32),
    pub p2: (String, u32),
    pub nb_games: u32,
}

impl GameSession {
    pub fn new(id: u32, p1_name: String, p2_name: String, nb_games: u32) -> GameSession {
        todo!()
    }

    pub fn read_winner(&self) -> Option<&(String, u32)> {
        todo!()
    }

    pub fn update_score(&mut self, user_name: &str) {
        todo!()
    }

    pub fn delete(self) -> String {
        todo!()
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

    game.update_score("Joao");
    game.update_score("Joao");
    game.update_score("Susana");
    game.update_score("Susana");
    println!("{:?}", game.read_winner());

    game.update_score("Joao");
    // This one will not count because it already 5 games played, the `nb_games`
    game.update_score("Susana");

    println!("{:?}", game.read_winner());

    println!("{:?}", game.delete());

    // game.read_winner();
    // This will give an error as the game was dropped with `delete` and no longer exists
}
```

And its output:

```console
$ cargo run
None
None
Some(("Joao", 3))
"game deleted: id -> 0"
$
```

### Notions

- [Box Pointers](https://doc.rust-lang.org/book/ch15-01-box.html)
