## macro_calculator

### Instructions

Create a **function** `calculate_macros` which receives a vector of `serde_json::Value` and returns a `serde_json::Value`.

The vector the function will receive will contain jsons in the following format:

```json
{
    "name": <name>,
    "calories": [<value_in_kJ>, <value_in_kcal>],
    "fats": <fats_in_g>,
    "carbs": <carbs_in_g>,
    "proteins": <proteins_in_g>,
    "nbr_of_portions": <portions>
}
```

The name and the content of the calories array will be `strings`, all other values will be `floats`.

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

Hint: You will need the `serde`, `serde_json` and `serde_derive` crates.

### Notions

- [serde](https://crates.io/crates/serde)
- [serde JSON](https://crates.io/crates/serde_json)
- [serde_derive](https://crates.io/crates/serde_derive)

### Usage

Here is a program to test your function:

```rust
use macro_calculator::*;

fn main() {
    let a = serde_json::json!(
        {
        "name": "light milk",
        "calories": ["148kJ", "35kcal"],
        "protein": 3.5,
        "fats": 0.1,
        "carbs": 5.0,
        "nbr_of_portions": 0.7
    });

    let b = serde_json::json!({
        "name": "oat cookies",
        "calories": ["1996kJ", "477kcal"],
        "protein": 8.2,
        "fats": 21.0,
        "carbs": 60.4,
        "nbr_of_portions": 1.2,
    });

    let macros: Macros = serde_json::from_value(calculate_macros(vec![a, b])).unwrap();
    println!("{:?}", macros);
}
```

And its output:

```sh
$ cargo run
Macros { cals: 596.9, carbs: 75.98, proteins: 12.29, fats: 25.27 }
$
```
