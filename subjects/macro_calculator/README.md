## macro_calculator

### Instructions

Create a **function** named `calculate_macros` which receives a vector of `Food` structures and returns a `json::JsonValue`.

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

The values in the calories array will be of type `string`, all other values will be `f64`.

The json returned by `calculate_macros` will have the following format:

```json
"cals": <calories>,
"carbs": <carbs>,
"proteins": <proteins>,
"fats": <fats>,
```

Consider the number of portions, as the values of the macros refer to one portion. Each value should represent the sum of each micro-nutrient in the array. E.g. `cals` is the sum of all `calories`.
Every value should be `f64` and be rounded rounded to two decimal places, or one decimal place if it ends in a zero. E.g:
- `12.294` -> `12.29`
- `12.295` -> `12.30` -> `12.3`

### Expected Function

```rust
pub struct Food {
    //expected public fields
}

pub fn calculate_macros(foods: Vec<Food>) -> json::JsonValue {

}
```

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

### Notions

- [json](https://docs.rs/json/0.12.4/json/) crate
