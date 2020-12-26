// Create a enum that represent the card suits
use card_deck::{self, Card, Rank, Suit};

// Write a program that takes that returns a random card in the deck
// A standard deck of cards has 52 cards: 4 suits and 13 cards per suit
#[allow(dead_code)]
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

#[test]
fn test_winner() {
	let winner = Card {
		rank: Rank::Ace,
		suit: Suit::Spade,
	};
	for rank in 1..14 {
		for suit in 1..5 {
			let card = Card {
				rank: Rank::traslate(rank),
				suit: Suit::translate(suit),
			};
			if card != winner {
				assert!(!card_deck::winner_card(card));
			} else {
				assert!(card_deck::winner_card(card));
			}
		}
	}
}
