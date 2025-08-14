## traits

### Instructions

Imagine you are designing a new video game, and your next feature is to create food that the players can eat to gain strength.

There are two types of food for now:

- `Fruit`: increases the strength by 4 units for each kilogram of fruit consumed.
- `Meat`: has a weight in kilograms, and it's pure fat content as a decimal fraction. The remaining weight of meat which is not pure fat is considered to be protein. Each kilogram of protein increases the strength by 4 units. Each kilogram of fat increases the strength by 9 units.

Define the `Food` trait for `Fruit` and `Meat`. The required method `gives` returns the amount of strength that the food provides.

Implement the `std::fmt::Display` trait for the `Player` structure, so that when `Player` is printed, it will print information according to this example:

- player_name
- Strength: 4, Score: 10, Money: 400
- Weapons: ["toy knife", "tough glove"]

### Expected Functions and Structures

```rust
use std::fmt;

#[derive(Debug)]
pub struct Player<'a> {
	pub name: &'a str,
	pub strength: f64,
	pub score: u32,
	pub money: u32,
	pub weapons: Vec<&'a str>,
}

pub struct Fruit {
	pub weight_in_kg: f64,
}

pub struct Meat {
	pub weight_in_kg: f64,
	pub fat_content: f64,
}

impl Player<'_> {
	pub fn eat(&mut self, food: impl Food) {
		self.strength += food.gives();
	}
}

impl fmt::Display for Player<'_> {
	fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
		todo!()
    }
}

pub trait Food {
	fn gives(&self) -> f64;
}

impl Food for Fruit {
	fn gives(&self) -> f64 {
		todo!()
	}
}

impl Food for Meat {
	fn gives(&self) -> f64 {
		todo!()
	}
}
```

### Usage

Here is a program to test your functions and traits.

```rust
use traits::*;


fn main() {
	let apple = Fruit { weight_in_kg: 1.0 };

	println!("this apple gives {} units of strength", apple.gives());

	let steak = Meat {
		weight_in_kg: 1.0,
		fat_content: 1.0,
	};

	let mut player1 = Player {
		name: "player1",
		strength: 1.0,
		score: 0,
		money: 0,
		weapons: vec!["knife"],
	};

	println!("Before eating {:?}", player1);

	player1.eat(apple);
	println!("After eating an apple\n{}", player1);

	player1.eat(steak);
	println!("After eating a steak\n{}", player1);
}

```

And its output:

```console
$ cargo run
this apple gives 4 units of strength
Before eating Player { name: "player1", strength: 1.0, score: 0, money: 0, weapons: ["knife"] }
After eating an apple
player1
Strength: 5, Score: 0, Money: 0
Weapons: ["knife"]
After eating a steak
player1
Strength: 14, Score: 0, Money: 0
Weapons: ["knife"]
$
```

### Notions

- [Traits](https://doc.rust-lang.org/book/ch10-02-traits.html)
