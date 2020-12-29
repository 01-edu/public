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

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn testing_macros_values() {
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

        assert_eq!(
            Macros {
                cals: 596.9,
                carbs: 75.98,
                proteins: 12.29,
                fats: 25.27
            },
            macros
        );
    }

    #[test]
    fn testing_no_food() {
        let macros: Macros = serde_json::from_value(calculate_macros(vec![])).unwrap();

        assert_eq!(
            Macros {
                cals: 0.0,
                carbs: 0.0,
                proteins: 0.0,
                fats: 0.0
            },
            macros
        );
    }

    #[test]
    fn big_values() {
        let macros: Macros = serde_json::from_value(calculate_macros(vec![
            serde_json::json!({
                "name": "big mac",
                "calories": ["2133.84kJ", "510kcal"],
                "protein": 27,
                "fats": 26,
                "carbs": 41,
                "nbr_of_portions": 2,
            }),
            serde_json::json!({
                "name": "pizza margherita",
                "calories": ["1500.59kJ", "358.65kcal"],
                "protein": 13.89,
                "fats": 11.21,
                "carbs": 49.07,
                "nbr_of_portions": 4.9,
            }),
        ]))
        .unwrap();

        assert_eq!(
            Macros {
                cals: 2777.39,
                carbs: 322.44,
                proteins: 122.06,
                fats: 106.93
            },
            macros
        );
    }
}
