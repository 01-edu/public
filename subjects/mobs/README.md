## mobs

### Instructions

Create a structure `Mob` which has:

- `name`: `String`
- `boss`: `Boss`
- `members`: a HashMap of `Member`s keyed by `String`s
- `cities`: a `HashSet` of city names
- `wealth`: `u64`

- `recruit`: a method which adds a `Member` to the `members` vector. It should accept a tuple with the member's information, `name` and `age` (`(&str, u32)`). The member's role should be set to the lowest one, `Associate`.
- `attack`: a method which receives another `Mob` as reference. It will remove the youngest member(s) from the vector of `Member` from whichever mob has the least power combat score. In the case of a draw, the attacker loses. Furthermore, if the loser is left with zero members, the victorious mob will **also** take the cities and wealth from the losing mob. The power combat score is calculated from the sum of the `role` of each mob `member`:
  - `Underboss`: 4
  - `Caporegime`: 3
  - `Soldier`: 2
  - `Associate`: 1
- `steal`: a method which receives a `Mob` to target, and a `u64` value to steal. The `self` mob steals the value from the wealth of the target mob, and adds the value to its own wealth. Naturally, any mob can only be stolen as much money as they have in possession.
- `conquer_city`: a method which receives a slice of references of `Mob`, and a `String` city name. The city is added to its list of cities if none of the other mobs already have it.

You will also need to create two submodules:

- `boss`: which should contain:
  - `Boss`: a struct which consists of:
  - `name`: `String`
  - `age`: `u32`
  - `new`: an associated function which accepts a `name: &str` and `age: u32`, and returns a `Boss`.
- `member` submodule which consists of:
  - `Role`: an enum with the variants:
    - `Underboss`
    - `Caporegime`
    - `Soldier`
    - `Associate`
  - `Member`: a struct which consists of:
    - `role`: `Role`
    - `age`: `u32`

    - `get_promotion`: a method which when invoked should promote the member from:
      - `Associate` -> `Soldier`
      - `Soldier` -> `Caporegime`
      - `Caporegime` -> `Underboss`
      - `Underboss` -> Impossibility. Panic.

We advise you to have each submodule inside their own respective file, but that is up to you.

Every data structure should derive at least `Debug` and `PartialEq`.

### Expected Function

> You'll need to work out the function signatures for yourself.

### Usage

> You'll need to create your own tests and samples. Make sure you cover every possibility!

### Notions

- [Packages, Crates and Modules](https://doc.rust-lang.org/book/ch07-00-managing-growing-projects-with-packages-crates-and-modules.html)
