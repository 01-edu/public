## traits

### Instructions

Imagine you are designing a new video game and you have to create food that they players can take to gain strength. 

There are two types of food for now fruits and meet: fruits increases the strengths by 1 unit and meat increases it by 3 unit.

- Define both structures fruits and meat:

Define the std::fmt::Display trait of the Player structure so using the template {} inside a println! macro will print:

- In the first line the name of the player
- In the second line the strength, score and the money
- In the third line the weapons

### Expected Functions and Structures

```rust
#[derive(Debug)]
pub struct Player {
	pub name: String,
	pub strength: u32,
	pub score: i32,
	pub money: i32,
	pub weapons: Vec<String>,
}

pub struct Fruit {
	pub weight_in_kg: f64,
}

pub struct Meat {
	pub weight_in_kg: f64,
	pub fat_content: f64,
}

impl Player {
	fn eat(&mut self, food: T) {
		self.strength += food.gives();
	}
}

pub trait Food {
	fn gives(&self) -> u32;
}

impl Food for Fruit {
}

impl Food for Meat {
}
```

### Usage

Here is a program to test your function.

```rust
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
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
Before eating Player { name: "player1", strength: 1, score: 0, money: 0, weapons: ["knife"] }
After eating an apple
Player { name: "player1", strength: 5, score: 0, money: 0, weapons: ["knife"] }
After eating a steak
Player { name: "player1", strength: 14, score: 0, money: 0, weapons: ["knife"] }
student@ubuntu:~/[[ROOT]]/test$
```
