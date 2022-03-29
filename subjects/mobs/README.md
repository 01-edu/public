## mobs

### Instructions

Create a module `mobs`, in which has to be present:

- a `Mob` structure which consists of:
  - a String `name`
  - a Boss struct `boss`
  - a vector of Members `members`
  - a vector of tuples containing a name String and a value u8 `cities`
  - a u32 `wealth`

The Mob struct should implement the following functions:

- `recruit`, which adds a member to the members vector of the Mob when receiving a `member_name` (&str) and a `member_age` (u8) (the role should be Associate)
- `attack`, which receives another Mob and will remove the last member from the vector of Members from the Mob with the less power combat (in case of draw, the loser is the attacker). The power combat is calculated by its members:
  - an Underboss power combat is 4
  - a Caporegime power combat is 3
  - a Soldier power combat is 2
  - an Associate power combat is 1
  - In case of one of the mobs stays without members, the winner mob adds to its cities every city of the loser mob and the same happens to the wealth, and the loser mob loses all cities and all wealth
- `steal`, which receives the targeted mob (Mob) and a value (u32) and adds to the own mob a value and subtracts the value
- `conquer_city`, which receives a vector of mobs, a city name (String) and a value (u8) and adds it to the vector of cities of the own mob, only if no mob in the vector owns a city with this name

Create also two submodules of mob:

- `boss` submodule which should contain:
  - a `Boss` struct which consists of:
  - a String `name`
  - an u8 `age`
  - a function `new` which returns a Boss on receiving a name (&str) and an age (u8)
- `member` submodule which consists of:
  - a enum `Role` with the variants:
    - Underboss
    - Caporegime
    - Soldier
    - Associate
  - a `Member` struct which consists of:
    - a String name
    - a enum Role `role`
    - a u8 `age`
  - the Member struct should implement a function `get_promotion` which will change the member role. If a member is an Associate, it will get promoted to Soldier; a Soldier is promoted to a Caporegime and a Caporegime is promoted to Underboss
  - a function `new` which return a Member on receiving a name(&str), a role (Role) and an age (u8)

The submodules should be created inside a folder called `mobs`. We advise you to create two files one to each submodule, but it is up to you.

You must include `#[derive(Debug, CLone, PartialEq)]` on top of every struct and the enum.

### Expected Function

#### For this exercise the signature of the function has to be found out.

### Usage

Here is a program to test your function:

```rust
use mobs::*;

fn main() {
  let (mafia1, mafia2) = (
    Mob {
      name: "Hairy Giants".to_string(),
      boss: boss::new("Louie HaHa", 36),
      cities: vec![("San Francisco".to_string(), 7)],
      members: vec![
        member::new("Benny Eggs", member::Role::Soldier, 28),
        member::new("Jhonny", member::Role::Associate, 17),
        member::new("Greasy Thumb", member::Role::Soldier, 30),
        member::new("No Finger", member::Role::Caporegime, 32),
      ],
      wealth: 100000,
    },
    Mob {
      name: "Red Thorns".to_string(),
      boss: boss::new("Big Tuna", 30),
      cities: vec![("San Jose".to_string(), 5)],
      members: vec![
        member::new("Knuckles", member::Role::Soldier, 25),
        member::new("Baldy Dom", member::Role::Caporegime, 36),
        member::new("Crazy Joe", member::Role::Underboss, 23),
      ],
      wealth: 70000,
    },
  );

  println!("{:?}\n{:?}", mafia1, mafia2);
}
```

And its output:

```sh
$ cargo run
Mob { name: "Hairy Giants", boss: Boss { name: "Louie HaHa", age: 36 }, members: [Member { name: "Benny Eggs", role: Soldier, age: 28 }, Member { name: "Jhonny", role: Associate, age: 17 }, Member { name: "Greasy Thumb", role: Soldier, age: 30 }, Member { name: "No Finger", role: Caporegime, age: 32 }], cities: [("San Francisco", 7)], wealth: 100000 }
Mob { name: "Red Thorns", boss: Boss { name: "Big Tuna", age: 30 }, members: [Member { name: "Knuckles", role: Soldier, age: 25 }, Member { name: "Baldy Dom", role: Caporegime, age: 36 }, Member { name: "Crazy Joe", role: Underboss, age: 23 }], cities: [("San Jose", 5)], wealth: 70000 }
$
```
