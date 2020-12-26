## card_deck

### Instructions

Represent cards from a desk

- A standard deck of cards has 52 cards: 4 suits and 13 cards per suit

- Start by creating the `Suit` enum and implement the associated
function `random` which returns a random `Suit` (`Heart`,
`Diamond`, `Spade` or `Club`)

- Then create the `Rank` enum that can have the value
`Ace`, `King`, `Queen`, `Jack`, and `Number` associated to an `u8`
value to represent the ranks 2 through 10
After create an associated function to `Rank` called `Random` that
returns a random `Rank`

- Finally create a structure name `Card` which has the fields `suit`
and `rank`

Define:

The associated function `translate` for Rank and Suite
- for `Suit`, `tranlate` makes the translation between an integer value (u8) and the suit of a card (1 -> Heart, 2 -> Diamonds, 3 -> Spade, 4 -> Club)
- for `Rank`, `translate` makes the tranlation between an integer value (u8) and the rank ( 1 -> Ace, 2 -> 2, .., 10 -> 10, 11 -> Jack, 12 -> Queen, 13 -> King)

The associated function `random` for `Rank` and `Suit` which returns a random rand and suit respectively

Finally define the function `winner_card` that returns true if the card passed as an argument is an Ace of spades

### Expected Functions and Structures

```rust
pub fn random() -> Suit {
}

pub fn translate(value: u8) -> Suit {
}

pub fn random() -> Rank {
}

pub fn traslate(value: u8) -> Rank {
}

pub enum Suit {
}

pub enum Rank {
}

pub struct Card {
	pub suit: Suit,
	pub rank: Rank,
}
```

### Usage

Here is a program to test your function

```rust
fn main() {
	let your_card = Card {
		rank: Rank::random(),
		suit: Suit::random(),
	};

	println!("You're card is {:?}", your_card);

	// Now if the card is an Ace of Spades print "You are the winner"
	if card_deck::winner_card(your_card) {
		println!("You are the winner!");
	}
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
You're card is Card { suit: Club, rank: Ace }
student@ubuntu:~/[[ROOT]]/test$
```
