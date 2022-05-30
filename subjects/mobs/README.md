## mobs

### Instructions

Create a module named `mobs`, containing a structure `Mob` which has:

- `name`: `String`
- `boss`: `Boss`
- `members`: a vector of `Member`
- `cities`: a vector of tuples containing a city name and a `u8`
- `wealth`: `u32`
- `recruit`: an associated function which adds a `Member` to the `members` vector. It should accept a `name`, and an `age`. The member's role should be set to `Associate`.
- `attack`: an associated function which receives another `Mob`. It will remove the last member from the vector of `Member` from whichever mob has the least power combat score. In the case of a draw, the attacker loses. In the case that one of the mobs is left with zero members, the victorious mob will **also** take the cities and wealth from the losing mob. The power combat score is calculated from the sum of the `role` of each mob `member`:
  - `Underboss`: 4
  - `Caporegime`: 3
  - `Soldier`: 2
  - `Associate`: 1
- `steal`: an associated function which receives a `Mob` to target, and a `u32` value to steal. The 'self' mob steals the value from the wealth of the target mob, and adds the value to its own wealth. Only as much money as the target mob has can be stolen.
- `conquer_city`: an associated function which receives a vector of `Mob`, a city name and a `u8` value. The city name and `u8` value are added to its list of cities if non of the other mobs in the vector have a city with the same name.

You will also need to create two submodules of mob:

- `boss`: which should contain:
  - `Boss`: a struct which consists of:
  - `name`: `String`
  - `age`: `u8`
  - `new`: an associated function which accepts a `name` and `age`, and returns a `Boss`.
- `member` submodule which consists of:
  - `Role`: an enum with the variants:
    - `Underboss`
    - `Caporegime`
    - `Soldier`
    - `Associate`
  - `Member`: a struct which consists of:
    - `name`: `String`
    - `role`: `Role`
    - `age`: `u8`
    `get_promotion`: an associated function which when invoked should promote the member from:
      - `Associate` -> `Soldier`
      - `Soldier` -> `Caporegime`
      - `Caporegime` -> `Underboss`
    - `new`: accepts a `name`, `role` and `age`, returning a `Member`.

The submodules should be created inside a folder named `mobs`. We advise you to create two files for each submodule, but that is up to you.

You must include `#[derive(Debug, CLone, PartialEq)]` above every struct and enum.

### Expected Function
> You'll need to work out the function signatures for yourself.


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
