// Imagine you are designing a new video game and you have to create
// food that they players can take to gain strength there are two
// types of food for now fruits and meet: fruits increases the
// strengths by 1 unit and meat increases it by 3 unit.

// Define both structures fruits and meat
// Define the std::fmt::Display trait of the Player structure so using
// the template {} inside a println! macro will print in the first
// line the name of the player
// in the second line the strength, score and the money
// and in the third line the weapons
use traits::{Food, Fruit, Meat, Player};

fn main() {
	let apple = Fruit { weight_in_kg: 1.0 };
	assert_eq!(apple.gives(), 4);
	let steak = Meat {
		weight_in_kg: 1.0,
		fat_content: 1.0,
	};

	let mut player1 = Player {
		name: String::from("player1"),
		strength: 1,
		score: 0,
		money: 0,
		weapons: vec![String::from("knife")],
	};
	println!("Before eating {:?}", player1);
	player1.eat(apple);
	println!("After eating an apple\n{:?}", player1);
	player1.eat(steak);
	println!("After eating a steak\n{:?}", player1);
}

#[test]
fn test_gives() {
	let apple = Fruit { weight_in_kg: 1.0 };
	assert_eq!(apple.gives(), 4);
	let steak = Meat {
		weight_in_kg: 1.0,
		fat_content: 1.0,
	};
	assert_eq!(steak.gives(), 9);
}

#[test]
fn test_eat() {
	let apple = Fruit { weight_in_kg: 1.0 };
	assert_eq!(apple.gives(), 4);
	let steak = Meat {
		weight_in_kg: 1.0,
		fat_content: 1.0,
	};

	let mut player1 = Player {
		name: String::from("player1"),
		strength: 1,
		score: 0,
		money: 0,
		weapons: vec![String::from("knife")],
	};
	player1.eat(apple);
	assert_eq!(player1.strength, 5);
	player1.eat(steak);
	assert_eq!(player1.strength, 14);
}
