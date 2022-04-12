## macro_calculator

### Instructions

Create a **function** `calculate_macros` which receives a vector of `Food` structures and returns a `json::JsonValue`.

```rust
Food {
    name: <name>,
    calories: [<value_in_kJ>, <value_in_kcal>],
    fats: <fats_in_g>,
    carbs: <carbs_in_g>,
    proteins: <proteins_in_g>,
    nbr_of_portions: <portions>
}
```

Both values in the calories array will be `strings`, all other values will be `f64`s.

The function should return a json with the following format (Macros struct):

```json
    "cals": <calories>,
    "carbs": <carbs>,
    "proteins": <proteins>,
    "fats": <fats>,
```

The number of portions should be taken into account. The values of the macros refer to one portion.
All values should be floats (f64) and should be the addition of every macronutrient in the provided array (cals is the addition of every calories, fats is the addition of every fats, and so on...).
Every value should be rounded to two decimal places (ex: 12.29) or one decimal place if it ends in 0 (ex: 18.90 -> 18.9).

Hint: You will need the `json` crate.

### Expected Function

```rust
pub struct Food {
    //expected public fields
}

pub fn calculate_macros(foods: Vec<Food>) -> json::JsonValue {

}
```

### Notions

- [json](https://docs.rs/json/0.12.4/json/) crate

### Usage

Here is a program to test your function:

```rust
use macro_calculator::*;

fn main(){
    let a = vec![
        Food {
            name: String::from("big mac"),
            calories: ["2133.84kJ".to_string(), "510kcal".to_string()],
            proteins: 27.0,
            fats: 26.0,
            carbs: 41.0,
            nbr_of_portions: 2.0,
        },
        Food {
            name: "pizza margherita".to_string(),
            calories: ["1500.59kJ".to_string(), "358.65kcal".to_string()],
            proteins: 13.89,
            fats: 11.21,
            carbs: 49.07,
            nbr_of_portions: 4.9,
        },
    ];

    println!("{:#}", calculate_macros(a));
}
```

And its output:

```sh
$ cargo run
{
    "cals": 2777.39,
    "carbs": 322.44,
    "proteins": 122.06,
    "fats": 106.93
}
$
```
