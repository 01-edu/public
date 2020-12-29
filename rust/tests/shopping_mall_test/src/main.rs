/*
## shopping_mall

### Instructions

You will have to create several functions to help run a shopping mall, with the help of the `mall` module provided:

- `biggest_store`: receives a `mall::Mall` and returns the `Store` with the biggest `square_meters`;
- `highest_paid_employees`: receives a `mall::Mall` and returns a vector containing the `Employee`(s) with the highest salaries;
- `nbr_of_employees`: receives a `mall::Mall` and returns the number of employees and securities, as an `usize`, in that mall.
- `fire_old_securities`: receives a `mall::Mall` and removes from the `mall::Mall.securities` all securities who are 50 years old or older.
- `check_securities`: receives a `mall::Mall` and a vector of `Security` and if there are not at least 1 security for every 200 square meters of floor size, there should be added a security to the `mall::Mall.securities`
- `cut_or_raise`: receives a `mall::Mall and raises or cuts the salary  of every employee in the mall by 10% depending if the employee works for more than 10 hours

### Expected Functions (and Structures)

```rust
pub fn biggest_store(mall: mall::Mall) -> store::Store {}

pub fn highest_paid_employee(mall: mall::Mall) -> Vec<employee::Employee> {}

pub fn nbr_of_employees(mall: mall::Mall) -> usize {}

pub fn fire_old_securities(mall: &mut mall::Mall) {}

pub fn check_for_securities(mall: &mut mall::Mall, available_sec: Vec<security::Security>) {}

pub fn cut_or_raise(mall: &mut mall::Mall) {}
```

### Usage

Here is a program to test your function:

```rust
fn main() {
    let secs = vec![
        mall::security::Security::new("John Oliver", 34, 7),
        mall::security::Security::new("Logan West", 23, 2),
        mall::security::Security::new("Bob Schumacher", 53, 15),
    ];

    let footzo_emp = vec![
        mall::floor::store::employee::Employee::new("Finbar Haines", 36, 9, 14, 650.88),
        mall::floor::store::employee::Employee::new("Roksanna Rocha", 45, 13, 22, 772.00),
        mall::floor::store::employee::Employee::new("Sienna-Rose Penn", 26, 9, 22, 1000.43),
    ];
    let swashion_emp = vec![
        mall::floor::store::employee::Employee::new("Abdallah Stafford", 54, 8, 22, 1234.21),
        mall::floor::store::employee::Employee::new("Marian Snyder", 21, 8, 14, 831.90),
        mall::floor::store::employee::Employee::new("Amanda Mclean", 29, 13, 22, 1222.12),
        mall::floor::store::employee::Employee::new("Faizaan Castro", 32, 11, 18, 1106.43),
    ];
    let pizbite_emp = vec![
        mall::floor::store::employee::Employee::new("Juniper Cannon", 21, 16, 23, 804.35),
        mall::floor::store::employee::Employee::new("Alena Simon", 28, 9, 15, 973.54),
        mall::floor::store::employee::Employee::new("Yasemin Collins", 29, 9, 19, 986.33),
        mall::floor::store::employee::Employee::new("Areeb Roberson", 54, 9, 22, 957.82),
        mall::floor::store::employee::Employee::new("Rocco Amin", 44, 13, 23, 689.21),
    ];
    let grill_emp = vec![
        mall::floor::store::employee::Employee::new("Rhian Crowther", 45, 9, 15, 841.18),
        mall::floor::store::employee::Employee::new("Nikkita Steadman", 52, 14, 22, 858.61),
        mall::floor::store::employee::Employee::new("Reginald Poole", 32, 9, 22, 1197.64),
        mall::floor::store::employee::Employee::new("Minnie Bull", 54, 14, 22, 1229.73),
    ];
    let sumo_emp = vec![
        mall::floor::store::employee::Employee::new("Chantelle Barajas", 20, 8, 22, 969.22),
        mall::floor::store::employee::Employee::new("Hywel Rudd", 49, 12, 22, 695.74),
        mall::floor::store::employee::Employee::new("Marianne Beasley", 55, 8, 14, 767.83),
    ];
    let supermaket_emp = vec![
        mall::floor::store::employee::Employee::new("Amara Schaefer", 23, 9, 14, 796.21),
        mall::floor::store::employee::Employee::new("Yara Wickens", 39, 9, 14, 853.42),
        mall::floor::store::employee::Employee::new("Tomi Boyer", 64, 9, 14, 881.83),
        mall::floor::store::employee::Employee::new("Greta Dickson", 42, 9, 14, 775.10),
        mall::floor::store::employee::Employee::new("Caroline Finnegan", 41, 9, 14, 702.92),
        mall::floor::store::employee::Employee::new("Indiana Baxter", 33, 13, 20, 991.71),
        mall::floor::store::employee::Employee::new("Jadine Page", 48, 13, 20, 743.21),
        mall::floor::store::employee::Employee::new("Husna Ryan", 43, 13, 20, 655.75),
        mall::floor::store::employee::Employee::new("Tyler Hunt", 63, 13, 20, 668.25),
        mall::floor::store::employee::Employee::new("Dahlia Caldwell", 56, 13, 20, 781.38),
        mall::floor::store::employee::Employee::new("Chandler Mansell", 20, 19, 24, 656.75),
        mall::floor::store::employee::Employee::new("Mohsin Mcgee", 30, 19, 24, 703.83),
        mall::floor::store::employee::Employee::new("Antoine Goulding", 45, 19, 24, 697.12),
        mall::floor::store::employee::Employee::new("Mark Barnard", 53, 19, 24, 788.81),
    ];

    let ground_stores = vec![
        store::Store::new("Footzo", 50, footzo_emp),
        store::Store::new("Swashion", 43, swashion_emp),
    ];
    let food_stores = vec![
        store::Store::new("PizBite", 60, pizbite_emp),
        store::Store::new("Chillout Grill", 50, grill_emp),
        store::Store::new("Sumo Food", 30, sumo_emp),
    ];
    let supermarket = vec![store::Store::new("Pretail", 950, supermaket_emp)];

    let floors = vec![
        floor::Floor::new("Ground Floor", ground_stores, 300),
        floor::Floor::new("Food Floor", food_stores, 500),
        floor::Floor::new("Supermarket", supermarket, 1000),
    ];

    let mall_la_vie = mall::Mall::new("La Vie Funchal", secs, floors);

    println!("{:?}", &mall_la_vie);
}
```


And its output:

```sh
$ cargo run
Mall { name: "La Vie Funchal", securities: [Security { name: "John Oliver", age: 34, years_experience: 7 }, Security { name: "Logan West", age: 23, years_experience: 2 }, Security { name: "Bob Schumacher", age: 53, years_experience: 15 }], floors: [Floor { name: "Ground Floor", stores: [Store { name: "Footzo", square_meters: 50, employees: [Employee { name: "Finbar Haines", age: 36, working_hours: (9, 14), salary: 650.88 }, Employee { name: "Roksanna Rocha", age: 45, working_hours: (13, 22), salary: 772.0 }, Employee { name: "Sienna-Rose Penn", age: 26, working_hours: (9, 22), salary: 1000.43 }] }, Store { name: "Swashion", square_meters: 43, employees: [Employee { name: "Abdallah Stafford", age: 54, working_hours: (8, 22), salary: 1234.21 }, Employee { name: "Marian Snyder", age: 21, working_hours: (8, 14), salary: 831.9 }, Employee { name: "Amanda Mclean", age: 29, working_hours: (13, 22), salary: 1222.12 }, Employee { name: "Faizaan Castro", age: 32, working_hours: (11, 18), salary: 1106.43 }] }], size_limit: 300 }, Floor { name: "Food Floor", stores: [Store { name: "PizBite", square_meters: 60, employees: [Employee { name: "Juniper Cannon", age: 21, working_hours: (16, 23), salary: 804.35 }, Employee { name: "Alena Simon", age: 28, working_hours: (9, 15), salary: 973.54 }, Employee { name: "Yasemin Collins", age: 29, working_hours: (9, 19), salary: 986.33 }, Employee { name: "Areeb Roberson", age: 54, working_hours: (9, 22), salary: 957.82 }, Employee { name: "Rocco Amin", age: 44, working_hours: (13, 23), salary: 689.21 }] }, Store { name: "Chillout Grill", square_meters: 50, employees: [Employee { name: "Rhian Crowther", age: 45, working_hours: (9, 15), salary: 841.18 }, Employee { name: "Nikkita Steadman", age: 52, working_hours: (14, 22), salary: 858.61 }, Employee { name: "Reginald Poole", age: 32, working_hours: (9, 22), salary: 1197.64 }, Employee { name: "Minnie Bull", age: 54, working_hours: (14, 22), salary: 1229.73 }] }, Store { name: "Sumo Food", square_meters: 30, employees: [Employee { name: "Chantelle Barajas", age: 20, working_hours: (8, 22), salary: 969.22 }, Employee { name: "Hywel Rudd", age: 49, working_hours: (12, 22), salary: 695.74 }, Employee { name: "Marianne Beasley", age: 55, working_hours: (8, 14), salary: 767.83 }] }], size_limit: 500 }, Floor { name: "Supermarket", stores: [Store { name: "Pretail", square_meters: 950, employees: [Employee { name: "Amara Schaefer", age: 23, working_hours: (9, 14), salary: 796.21 }, Employee { name: "Yara Wickens", age: 39, working_hours: (9, 14), salary: 853.42 }, Employee { name: "Tomi Boyer", age: 64, working_hours: (9, 14), salary: 881.83 }, Employee { name: "Greta Dickson", age: 42, working_hours: (9, 14), salary: 775.1 }, Employee { name: "Caroline Finnegan", age: 41, working_hours: (9, 14), salary: 702.92 }, Employee { name: "Indiana Baxter", age: 33, working_hours: (13, 20), salary: 991.71 }, Employee { name: "Jadine Page", age: 48, working_hours: (13, 20), salary: 743.21 }, Employee { name: "Husna Ryan", age: 43, working_hours: (13, 20), salary: 655.75 }, Employee { name: "Tyler Hunt", age: 63, working_hours: (13, 20), salary: 668.25 }, Employee { name: "Dahlia Caldwell", age: 56, working_hours: (13, 20), salary: 781.38 }, Employee { name: "Chandler Mansell", age: 20, working_hours: (19, 24), salary: 656.75 }, Employee { name: "Mohsin Mcgee", age: 30, working_hours: (19, 24), salary: 703.83 }, Employee { name: "Antoine Goulding", age: 45, working_hours: (19, 24), salary: 697.12 }, Employee { name: "Mark Barnard", age: 53, working_hours: (19, 24), salary: 788.81 }] }], size_limit: 1000 }] }
$
}
```
*/

fn main() {
    let secs = vec![
        mall::security::Security::new("John Oliver", 34, 7),
        mall::security::Security::new("Logan West", 23, 2),
        mall::security::Security::new("Bob Schumacher", 53, 15),
    ];

    let footzo_emp = vec![
        mall::floor::store::employee::Employee::new("Finbar Haines", 36, 9, 14, 650.88),
        mall::floor::store::employee::Employee::new("Roksanna Rocha", 45, 13, 22, 772.00),
        mall::floor::store::employee::Employee::new("Sienna-Rose Penn", 26, 9, 22, 1000.43),
    ];
    let swashion_emp = vec![
        mall::floor::store::employee::Employee::new("Abdallah Stafford", 54, 8, 22, 1234.21),
        mall::floor::store::employee::Employee::new("Marian Snyder", 21, 8, 14, 831.90),
        mall::floor::store::employee::Employee::new("Amanda Mclean", 29, 13, 22, 1222.12),
        mall::floor::store::employee::Employee::new("Faizaan Castro", 32, 11, 18, 1106.43),
    ];
    let pizbite_emp = vec![
        mall::floor::store::employee::Employee::new("Juniper Cannon", 21, 16, 23, 804.35),
        mall::floor::store::employee::Employee::new("Alena Simon", 28, 9, 15, 973.54),
        mall::floor::store::employee::Employee::new("Yasemin Collins", 29, 9, 19, 986.33),
        mall::floor::store::employee::Employee::new("Areeb Roberson", 54, 9, 22, 957.82),
        mall::floor::store::employee::Employee::new("Rocco Amin", 44, 13, 23, 689.21),
    ];
    let grill_emp = vec![
        mall::floor::store::employee::Employee::new("Rhian Crowther", 45, 9, 15, 841.18),
        mall::floor::store::employee::Employee::new("Nikkita Steadman", 52, 14, 22, 858.61),
        mall::floor::store::employee::Employee::new("Reginald Poole", 32, 9, 22, 1197.64),
        mall::floor::store::employee::Employee::new("Minnie Bull", 54, 14, 22, 1229.73),
    ];
    let sumo_emp = vec![
        mall::floor::store::employee::Employee::new("Chantelle Barajas", 20, 8, 22, 969.22),
        mall::floor::store::employee::Employee::new("Hywel Rudd", 49, 12, 22, 695.74),
        mall::floor::store::employee::Employee::new("Marianne Beasley", 55, 8, 14, 767.83),
    ];
    let supermaket_emp = vec![
        mall::floor::store::employee::Employee::new("Amara Schaefer", 23, 9, 14, 796.21),
        mall::floor::store::employee::Employee::new("Yara Wickens", 39, 9, 14, 853.42),
        mall::floor::store::employee::Employee::new("Tomi Boyer", 64, 9, 14, 881.83),
        mall::floor::store::employee::Employee::new("Greta Dickson", 42, 9, 14, 775.10),
        mall::floor::store::employee::Employee::new("Caroline Finnegan", 41, 9, 14, 702.92),
        mall::floor::store::employee::Employee::new("Indiana Baxter", 33, 13, 20, 991.71),
        mall::floor::store::employee::Employee::new("Jadine Page", 48, 13, 20, 743.21),
        mall::floor::store::employee::Employee::new("Husna Ryan", 43, 13, 20, 655.75),
        mall::floor::store::employee::Employee::new("Tyler Hunt", 63, 13, 20, 668.25),
        mall::floor::store::employee::Employee::new("Dahlia Caldwell", 56, 13, 20, 781.38),
        mall::floor::store::employee::Employee::new("Chandler Mansell", 20, 19, 24, 656.75),
        mall::floor::store::employee::Employee::new("Mohsin Mcgee", 30, 19, 24, 703.83),
        mall::floor::store::employee::Employee::new("Antoine Goulding", 45, 19, 24, 697.12),
        mall::floor::store::employee::Employee::new("Mark Barnard", 53, 19, 24, 788.81),
    ];

    let ground_stores = vec![
        store::Store::new("Footzo", 50, footzo_emp),
        store::Store::new("Swashion", 43, swashion_emp),
    ];
    let food_stores = vec![
        store::Store::new("PizBite", 60, pizbite_emp),
        store::Store::new("Chillout Grill", 50, grill_emp),
        store::Store::new("Sumo Food", 30, sumo_emp),
    ];
    let supermarket = vec![store::Store::new("Pretail", 950, supermaket_emp)];

    let floors = vec![
        floor::Floor::new("Ground Floor", ground_stores, 300),
        floor::Floor::new("Food Floor", food_stores, 500),
        floor::Floor::new("Supermarket", supermarket, 1000),
    ];

    let mall_la_vie = mall::Mall::new("La Vie Funchal", secs, floors);

    println!("{:?}", &mall_la_vie);
}

use mall::*;
use shopping_mall as mall;

mod tests {
    use super::*;

    fn create_mall() -> mall::Mall {
        let secs = vec![
            mall::security::Security::new("John Oliver", 34, 7),
            mall::security::Security::new("Logan West", 23, 2),
            mall::security::Security::new("Bob Schumacher", 53, 15),
        ];

        let footzo_emp = vec![
            mall::floor::store::employee::Employee::new("Finbar Haines", 36, 9, 14, 650.88),
            mall::floor::store::employee::Employee::new("Roksanna Rocha", 45, 13, 22, 772.00),
            mall::floor::store::employee::Employee::new("Sienna-Rose Penn", 26, 9, 22, 1000.43),
        ];
        let swashion_emp = vec![
            mall::floor::store::employee::Employee::new("Abdallah Stafford", 54, 8, 22, 1234.21),
            mall::floor::store::employee::Employee::new("Marian Snyder", 21, 8, 14, 831.90),
            mall::floor::store::employee::Employee::new("Amanda Mclean", 29, 13, 22, 1222.12),
            mall::floor::store::employee::Employee::new("Faizaan Castro", 32, 11, 18, 1106.43),
        ];
        let pizbite_emp = vec![
            mall::floor::store::employee::Employee::new("Juniper Cannon", 21, 16, 23, 804.35),
            mall::floor::store::employee::Employee::new("Alena Simon", 28, 9, 15, 973.54),
            mall::floor::store::employee::Employee::new("Yasemin Collins", 29, 9, 19, 986.33),
            mall::floor::store::employee::Employee::new("Areeb Roberson", 54, 9, 22, 957.82),
            mall::floor::store::employee::Employee::new("Rocco Amin", 44, 13, 23, 689.21),
        ];
        let grill_emp = vec![
            mall::floor::store::employee::Employee::new("Rhian Crowther", 45, 9, 15, 841.18),
            mall::floor::store::employee::Employee::new("Nikkita Steadman", 52, 14, 22, 858.61),
            mall::floor::store::employee::Employee::new("Reginald Poole", 32, 9, 22, 1197.64),
            mall::floor::store::employee::Employee::new("Minnie Bull", 54, 14, 22, 1229.73),
        ];
        let sumo_emp = vec![
            mall::floor::store::employee::Employee::new("Chantelle Barajas", 20, 8, 22, 969.22),
            mall::floor::store::employee::Employee::new("Hywel Rudd", 49, 12, 22, 695.74),
            mall::floor::store::employee::Employee::new("Marianne Beasley", 55, 8, 14, 767.83),
        ];
        let supermaket_emp = vec![
            mall::floor::store::employee::Employee::new("Amara Schaefer", 23, 9, 14, 796.21),
            mall::floor::store::employee::Employee::new("Yara Wickens", 39, 9, 14, 853.42),
            mall::floor::store::employee::Employee::new("Tomi Boyer", 64, 9, 14, 881.83),
            mall::floor::store::employee::Employee::new("Greta Dickson", 42, 9, 14, 775.10),
            mall::floor::store::employee::Employee::new("Caroline Finnegan", 41, 9, 14, 702.92),
            mall::floor::store::employee::Employee::new("Indiana Baxter", 33, 13, 20, 991.71),
            mall::floor::store::employee::Employee::new("Jadine Page", 48, 13, 20, 743.21),
            mall::floor::store::employee::Employee::new("Husna Ryan", 43, 13, 20, 655.75),
            mall::floor::store::employee::Employee::new("Tyler Hunt", 63, 13, 20, 668.25),
            mall::floor::store::employee::Employee::new("Dahlia Caldwell", 56, 13, 20, 781.38),
            mall::floor::store::employee::Employee::new("Chandler Mansell", 20, 19, 24, 656.75),
            mall::floor::store::employee::Employee::new("Mohsin Mcgee", 30, 19, 24, 703.83),
            mall::floor::store::employee::Employee::new("Antoine Goulding", 45, 19, 24, 697.12),
            mall::floor::store::employee::Employee::new("Mark Barnard", 53, 19, 24, 788.81),
        ];

        let ground_stores = vec![
            mall::floor::store::Store::new("Footzo", 50, footzo_emp),
            mall::floor::store::Store::new("Swashion", 43, swashion_emp),
        ];
        let food_stores = vec![
            mall::floor::store::Store::new("PizBite", 60, pizbite_emp),
            mall::floor::store::Store::new("Chillout Grill", 50, grill_emp),
            mall::floor::store::Store::new("Sumo Food", 30, sumo_emp),
        ];
        let supermarket = vec![mall::floor::store::Store::new(
            "Pretail",
            950,
            supermaket_emp,
        )];

        let floors = vec![
            mall::floor::Floor::new("Ground Floor", ground_stores, 300),
            mall::floor::Floor::new("Food Floor", food_stores, 500),
            mall::floor::Floor::new("Supermarket", supermarket, 1000),
        ];

        mall::Mall::new("La Vie Funchal", secs, floors)
    }

    #[test]
    fn biggest_store_tests() {
        let mut shopping_mall = create_mall();

        let mut tested = biggest_store(shopping_mall.clone());

        assert_eq!("Pretail", tested.name);
        assert_eq!(950, tested.square_meters);

        (&mut shopping_mall).floors[2].remove_store("Pretail".to_string());

        tested = biggest_store(shopping_mall.clone());

        assert_eq!("PizBite", tested.name);
        assert_eq!(60, tested.square_meters);
    }

    #[test]
    fn highest_paid_test() {
        let mut shopping_mall = create_mall();

        let mut tested = highest_paid_employee(shopping_mall.clone());
        assert_eq!(1, tested.len());
        assert_eq!("Abdallah Stafford", tested[0].name);
        assert_eq!(54, tested[0].age);

        let salary = shopping_mall.clone().floors[0].stores[0].employees[0].salary;
        shopping_mall.floors[0].stores[0].employees[0].raise(tested[0].salary - salary);
        tested = highest_paid_employee(shopping_mall.clone());

        assert_eq!(2, tested.len());
        assert_eq!("Finbar Haines", tested[0].name);
        assert_eq!(36, tested[0].age);

        assert_eq!(tested[1].salary, tested[0].salary);
    }

    #[test]
    fn fire_old_sec_test() {
        let mut shopping_mall = create_mall();

        fire_old_securities(&mut shopping_mall);
        assert_eq!(2, shopping_mall.securities.len());

        shopping_mall.securities.append(&mut vec![
            mall::security::Security::new("Chris Esparza", 54, 12),
            mall::security::Security::new("Kane Holloway", 53, 20),
            mall::security::Security::new("Connor Wardle", 22, 1),
            mall::security::Security::new("Louis Pickett", 26, 3),
            mall::security::Security::new("Olly Middleton", 36, 9),
        ]);

        assert_eq!(7, shopping_mall.securities.len());

        fire_old_securities(&mut shopping_mall);

        assert_eq!(5, shopping_mall.securities.len());
    }

    #[test]
    fn nbr_of_employees_test() {
        let mut shopping_mall = create_mall();

        let mut tested = nbr_of_employees(shopping_mall.clone());
        assert_eq!(36, tested);

        fire_old_securities(&mut shopping_mall);

        tested = nbr_of_employees(shopping_mall.clone());
        assert_eq!(35, tested);

        shopping_mall.floors[2].stores[0].employees = vec![];

        tested = nbr_of_employees(shopping_mall.clone());
        assert_eq!(21, tested);
    }

    #[test]
    fn check_for_securities_test() {
        let mut shopping_mall = create_mall();

        assert_eq!(3, shopping_mall.securities.len());

        check_for_securities(
            &mut shopping_mall,
            vec![
                mall::security::Security::new("Peter Solomons", 45, 20),
                mall::security::Security::new("William Charles", 32, 10),
                mall::security::Security::new("Leonardo Changretta", 23, 0),
                mall::security::Security::new("Vlad Levi", 38, 8),
                mall::security::Security::new("Faruk Berkai", 40, 15),
                mall::security::Security::new("Chritopher Smith", 35, 9),
                mall::security::Security::new("Jason Mackie", 26, 2),
                mall::security::Security::new("Kenzie Mair", 34, 8),
                mall::security::Security::new("Bentley Larson", 33, 10),
                mall::security::Security::new("Ray Storey", 37, 12),
            ],
        );

        assert_eq!(9, shopping_mall.securities.len());
    }

    #[test]
    fn cut_or_raise_test() {
        let mut shopping_mall = create_mall();

        cut_or_raise(&mut shopping_mall);
        assert_eq!(
            585.792,
            shopping_mall.floors[0].stores[0].employees[0].salary
        );
        assert_eq!(
            1100.473,
            shopping_mall.floors[0].stores[0].employees[2].salary
        );

        cut_or_raise(&mut shopping_mall);
        assert_eq!(
            527.2128,
            shopping_mall.floors[0].stores[0].employees[0].salary
        );
        assert_eq!(
            1210.5203,
            shopping_mall.floors[0].stores[0].employees[2].salary
        );
    }
}
