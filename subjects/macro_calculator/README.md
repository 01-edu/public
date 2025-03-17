## macro_calculator

### Instructions

Create a **function** named `calculate_macros` which receives a vector of `Food` structures and returns a `json::JsonValue`.

The `Food` structure has the following fields:

```rust
Food {
    name: <name>,
    calories: (<calories_in_kJ>kJ, <calories_in_kcal>kcal),
    fats: <fats_in_grams>,
    carbs: <carbs_in_grams>,
    proteins: <proteins_in_grams>,
    nbr_of_portions: <portions>
}
```

The `name` and the values in the `calories` tuple should be of type `String`. All other values should be represented as `f64`.

The json returned by `calculate_macros` will have the following format:

```json
{
    "cals": <calories_in_kcal>,
    "carbs": <carbs_in_grams>,
    "proteins": <proteins_in_grams>,
    "fats": <fats_in_grams>,
}
```

Consider the number of portions, as the values of the macros refer to one portion. Each value should represent the sum of each micro-nutrient in the array. E.g. `cals` is the sum of all calories of all the foods combined.
Every `f64` should be rounded to two decimal points or one decimal point if it ends in a zero. E.g:
- `12.294` -> `12.29`
- `12.295` -> `12.30` -> `12.3`

### Expected Function

```rust
pub struct Food {
    // expected public fields
}

pub fn calculate_macros(foods: &[Food]) -> json::JsonValue {
    todo!()
}
```

### Usage

Here is a program to test your function:

```rust
use macro_calculator::*;

fn main() {
    let foods = [
        Food {
            name: "big mac".to_owned(),
            calories: ("2133.84kJ", "510kcal"),
            proteins: 27.,
            fats: 26.,
            carbs: 41.,
            nbr_of_portions: 2.,
        },
        Food {
            name: "pizza margherita".to_owned(),
            calories: ("1500.59kJ", "358.65kcal"),
            proteins: 13.89,
            fats: 11.21,
            carbs: 49.07,
            nbr_of_portions: 4.9,
        },
    ];

    println!("{:#}", calculate_macros(foods));
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

### Notions

- [json-0.12.4](https://docs.rs/json/0.12.4/json/)
