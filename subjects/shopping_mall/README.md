## shopping_mall

### Instructions

Using the `mall` module provided, create the following **functions** to help run a shopping mall:

- `biggest_store`: receives a `Mall` and returns the `Store` with the biggest `square_meters`.
- `highest_paid_employee`: receives a `Mall` and returns a vector containing the `Employee`(s) with the highest salary.
- `nbr_of_employees`: receives a `Mall` and returns the number of employees and guards as a `usize`.
- `check_for_securities`: receives a `Mall` and a vector of `Guard`. If there is not at least 1 guard for every 200 square meters of floor size, a guard should be added to the `Mall.guards`.
- `cut_or_raise`: receives a `Mall`. For each employee, the salary will be raised by 10% if they work more than 10 hours, else their salary will be decreased by 10%. You can consider that guards are not employees of the mall.

### Expected Function

> You'll need to work out the function signatures for yourself.

### mall&#46;rs

```rust
#[derive(Debug, Clone, PartialEq)]
pub struct Mall {
    pub name: String,
    pub guards: Vec<guard::Guard>,
    pub floors: Vec<floor::Floor>,
}

impl Mall {
    #[allow(dead_code)]
    pub fn new(name: &str, guards: Vec<guard::Guard>, floors: Vec<floor::Floor>) -> Mall {
        Mall {
            name: name.to_string(),
            guards: guards,
            floors: floors,
        }
    }

    #[allow(dead_code)]
    pub fn change_name(&mut self, new_name: &str) {
        self.name = new_name.to_string();
    }

    #[allow(dead_code)]
    pub fn hire_guard(&mut self, guard: guard::Guard) {
        self.guards.push(guard);
    }

    #[allow(dead_code)]
    pub fn fire_guard(&mut self, name: String) {
        self.guards.retain(|x| x.name != name);
    }
}

pub mod guard {

    #[derive(Debug, Clone, PartialEq)]
    pub struct Guard {
        pub name: String,
        pub age: u8,
        pub years_experience: u8,
    }

    impl Guard {
        #[allow(dead_code)]
        pub fn new(name: &str, age: u8, years_experience: u8) -> Guard {
            Guard {
                name: name.to_string(),
                age: age,
                years_experience: years_experience,
            }
        }
    }
}

pub mod floor {

    #[derive(Debug, Clone, PartialEq)]
    pub struct Floor {
        pub name: String,
        pub stores: Vec<store::Store>,
        pub size_limit: u64,
    }

    impl Floor {
        #[allow(dead_code)]
        pub fn new(name: &str, stores: Vec<store::Store>, store_limit: u64) -> Floor {
            Floor {
                name: name.to_string(),
                stores: stores,
                size_limit: store_limit,
            }
        }

        #[allow(dead_code)]
        pub fn change_store(&mut self, store: &str, new_store: store::Store) {
            let pos = self.stores.iter().position(|x| x.name == store).unwrap();
            self.stores[pos] = new_store;
        }

        #[allow(dead_code)]
        pub fn add_store(&mut self, new_store: store::Store) {
            let mut current_floor_size = 0;

            for store in self.stores.iter() {
                current_floor_size += store.square_meters;
            }

            if self.size_limit < current_floor_size + new_store.square_meters {
                self.stores.push(new_store);
            }
        }

        #[allow(dead_code)]
        pub fn remove_store(&mut self, store_name: String) {
            self.stores.retain(|x| x.name != store_name);
        }
    }

    pub mod store {

        #[derive(Debug, Clone, PartialEq)]
        pub struct Store {
            pub name: String,
            pub square_meters: u64,
            pub employees: Vec<employee::Employee>,
        }

        impl Store {
            #[allow(dead_code)]
            pub fn new(name: &str, space: u64, employees: Vec<employee::Employee>) -> Store {
                Store {
                    name: name.to_string(),
                    square_meters: space,
                    employees: employees,
                }
            }

            #[allow(dead_code)]
            pub fn hire_employee(&mut self, employee: employee::Employee) {
                self.employees.push(employee);
            }
            #[allow(dead_code)]
            pub fn fire_employee(&mut self, employee_name: &str) {
                self.employees.retain(|x| x.name != employee_name);
            }
            #[allow(dead_code)]
            pub fn expand(&mut self, square_meters: u64) {
                self.square_meters += square_meters;
            }
        }

        pub mod employee {

            #[derive(Debug, Clone, PartialEq)]
            pub struct Employee {
                pub name: String,
                pub age: u8,
                pub working_hours: (u8, u8),
                pub salary: f64,
            }

            impl Employee {
                #[allow(dead_code)]
                pub fn new(
                    name: &str,
                    age: u8,
                    entry_hour: u8,
                    exit_hour: u8,
                    salary: f64,
                ) -> Employee {
                    Employee {
                        name: name.to_string(),
                        age: age,
                        working_hours: (entry_hour, exit_hour),
                        salary: salary,
                    }
                }

                #[allow(dead_code)]
                pub fn birthday(&mut self) {
                    self.age += 1;
                }

                #[allow(dead_code)]
                pub fn change_workload(&mut self, entry_hour: u8, exit_hour: u8) {
                    self.working_hours = (entry_hour, exit_hour);
                }

                #[allow(dead_code)]
                pub fn raise(&mut self, amount: f64) {
                    self.salary += amount;
                }

                #[allow(dead_code)]
                pub fn cut(&mut self, amount: f64) {
                    self.salary = self.salary - amount;
                }
            }
        }
    }
}
```

### Usage

Here is a program to test your function:

```rust
use shopping_mall::*;

fn main() {
    let secs = vec![
        mall::guard::Guard::new("John Oliver", 34, 7),
        mall::guard::Guard::new("Logan West", 23, 2),
        mall::guard::Guard::new("Bob Schumacher", 53, 15),
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

    let mut mall_la_vie = mall::Mall::new("La Vie Funchal", secs, floors);

    //returns the biggest store
    println!("Biggest store: {:#?}",biggest_store(mall_la_vie.clone()));

    //returns the list with the highest paid employees
    println!("Highest paid employee: {:#?}", highest_paid_employee(mall_la_vie.clone()));

    //returns the number of employees
    println!("Number of employees: {:?}", nbr_of_employees(mall_la_vie.clone()));

    //checks if it is needed to add securities
    check_for_securities(
        &mut mall_la_vie,
        vec![
            mall::guard::Guard::new("Peter Solomons", 45, 20),
            mall::guard::Guard::new("William Charles", 32, 10),
            mall::guard::Guard::new("Leonardo Changretta", 23, 0),
            mall::guard::Guard::new("Vlad Levi", 38, 8),
            mall::guard::Guard::new("Faruk Berkai", 40, 15),
            mall::guard::Guard::new("Chritopher Smith", 35, 9),
            mall::guard::Guard::new("Jason Mackie", 26, 2),
            mall::guard::Guard::new("Kenzie Mair", 34, 8),
            mall::guard::Guard::new("Bentley Larson", 33, 10),
            mall::guard::Guard::new("Ray Storey", 37, 12),
        ],
    );

    //raises or cuts the salary  of every employee
    cut_or_raise(&mut mall_la_vie);

    println!("{:#?}", &mall_la_vie);
}
```

And its ouput:

```rs
$ cargo run
Biggest store: Store {
    name: "Pretail",
    square_meters: 950,
    employees: [
        Employee {
            name: "Amara Schaefer",
            age: 23,
            working_hours: (
                9,
                14,
            ),
            salary: 796.21,
        },
        Employee {
            name: "Yara Wickens",
            age: 39,
            working_hours: (
                9,
                14,
            ),
            salary: 853.42,
        },
        Employee {
            name: "Tomi Boyer",
            age: 64,
            working_hours: (
                9,
                14,
            ),
            salary: 881.83,
        },
        Employee {
            name: "Greta Dickson",
            age: 42,
            working_hours: (
                9,
                14,
            ),
            salary: 775.1,
        },
        Employee {
            name: "Caroline Finnegan",
            age: 41,
            working_hours: (
                9,
                14,
            ),
            salary: 702.92,
        },
        Employee {
            name: "Indiana Baxter",
            age: 33,
            working_hours: (
                13,
                20,
            ),
            salary: 991.71,
        },
        Employee {
            name: "Jadine Page",
            age: 48,
            working_hours: (
                13,
                20,
            ),
            salary: 743.21,
        },
        Employee {
            name: "Husna Ryan",
            age: 43,
            working_hours: (
                13,
                20,
            ),
            salary: 655.75,
        },
        Employee {
            name: "Tyler Hunt",
            age: 63,
            working_hours: (
                13,
                20,
            ),
            salary: 668.25,
        },
        Employee {
            name: "Dahlia Caldwell",
            age: 56,
            working_hours: (
                13,
                20,
            ),
            salary: 781.38,
        },
        Employee {
            name: "Chandler Mansell",
            age: 20,
            working_hours: (
                19,
                24,
            ),
            salary: 656.75,
        },
        Employee {
            name: "Mohsin Mcgee",
            age: 30,
            working_hours: (
                19,
                24,
            ),
            salary: 703.83,
        },
        Employee {
            name: "Antoine Goulding",
            age: 45,
            working_hours: (
                19,
                24,
            ),
            salary: 697.12,
        },
        Employee {
            name: "Mark Barnard",
            age: 53,
            working_hours: (
                19,
                24,
            ),
            salary: 788.81,
        },
    ],
}
Highest paid employee: [
    Employee {
        name: "Abdallah Stafford",
        age: 54,
        working_hours: (
            8,
            22,
        ),
        salary: 1234.21,
    },
]
Number of employees: 36
Mall {
    name: "La Vie Funchal",
    guards: [
        Guard {
            name: "John Oliver",
            age: 34,
            years_experience: 7,
        },
        Guard {
            name: "Logan West",
            age: 23,
            years_experience: 2,
        },
        Guard {
            name: "Bob Schumacher",
            age: 53,
            years_experience: 15,
        },
        Guard {
            name: "Peter Solomons",
            age: 45,
            years_experience: 20,
        },
        Guard {
            name: "William Charles",
            age: 32,
            years_experience: 10,
        },
        Guard {
            name: "Leonardo Changretta",
            age: 23,
            years_experience: 0,
        },
        Guard {
            name: "Vlad Levi",
            age: 38,
            years_experience: 8,
        },
        Guard {
            name: "Faruk Berkai",
            age: 40,
            years_experience: 15,
        },
        Guard {
            name: "Chritopher Smith",
            age: 35,
            years_experience: 9,
        },
    ],
    floors: [
        Floor {
            name: "Ground Floor",
            stores: [
                Store {
                    name: "Footzo",
                    square_meters: 50,
                    employees: [
                        Employee {
                            name: "Finbar Haines",
                            age: 36,
                            working_hours: (
                                9,
                                14,
                            ),
                            salary: 585.792,
                        },
                        Employee {
                            name: "Roksanna Rocha",
                            age: 45,
                            working_hours: (
                                13,
                                22,
                            ),
                            salary: 694.8,
                        },
                        Employee {
                            name: "Sienna-Rose Penn",
                            age: 26,
                            working_hours: (
                                9,
                                22,
                            ),
                            salary: 1100.473,
                        },
                    ],
                },
                Store {
                    name: "Swashion",
                    square_meters: 43,
                    employees: [
                        Employee {
                            name: "Abdallah Stafford",
                            age: 54,
                            working_hours: (
                                8,
                                22,
                            ),
                            salary: 1357.631,
                        },
                        Employee {
                            name: "Marian Snyder",
                            age: 21,
                            working_hours: (
                                8,
                                14,
                            ),
                            salary: 748.71,
                        },
                        Employee {
                            name: "Amanda Mclean",
                            age: 29,
                            working_hours: (
                                13,
                                22,
                            ),
                            salary: 1099.908,
                        },
                        Employee {
                            name: "Faizaan Castro",
                            age: 32,
                            working_hours: (
                                11,
                                18,
                            ),
                            salary: 995.787,
                        },
                    ],
                },
            ],
            size_limit: 300,
        },
        Floor {
            name: "Food Floor",
            stores: [
                Store {
                    name: "PizBite",
                    square_meters: 60,
                    employees: [
                        Employee {
                            name: "Juniper Cannon",
                            age: 21,
                            working_hours: (
                                16,
                                23,
                            ),
                            salary: 723.915,
                        },
                        Employee {
                            name: "Alena Simon",
                            age: 28,
                            working_hours: (
                                9,
                                15,
                            ),
                            salary: 876.1859999999999,
                        },
                        Employee {
                            name: "Yasemin Collins",
                            age: 29,
                            working_hours: (
                                9,
                                19,
                            ),
                            salary: 1084.963,
                        },
                        Employee {
                            name: "Areeb Roberson",
                            age: 54,
                            working_hours: (
                                9,
                                22,
                            ),
                            salary: 1053.602,
                        },
                        Employee {
                            name: "Rocco Amin",
                            age: 44,
                            working_hours: (
                                13,
                                23,
                            ),
                            salary: 758.1310000000001,
                        },
                    ],
                },
                Store {
                    name: "Chillout Grill",
                    square_meters: 50,
                    employees: [
                        Employee {
                            name: "Rhian Crowther",
                            age: 45,
                            working_hours: (
                                9,
                                15,
                            ),
                            salary: 757.0619999999999,
                        },
                        Employee {
                            name: "Nikkita Steadman",
                            age: 52,
                            working_hours: (
                                14,
                                22,
                            ),
                            salary: 772.749,
                        },
                        Employee {
                            name: "Reginald Poole",
                            age: 32,
                            working_hours: (
                                9,
                                22,
                            ),
                            salary: 1317.404,
                        },
                        Employee {
                            name: "Minnie Bull",
                            age: 54,
                            working_hours: (
                                14,
                                22,
                            ),
                            salary: 1106.757,
                        },
                    ],
                },
                Store {
                    name: "Sumo Food",
                    square_meters: 30,
                    employees: [
                        Employee {
                            name: "Chantelle Barajas",
                            age: 20,
                            working_hours: (
                                8,
                                22,
                            ),
                            salary: 1066.142,
                        },
                        Employee {
                            name: "Hywel Rudd",
                            age: 49,
                            working_hours: (
                                12,
                                22,
                            ),
                            salary: 765.314,
                        },
                        Employee {
                            name: "Marianne Beasley",
                            age: 55,
                            working_hours: (
                                8,
                                14,
                            ),
                            salary: 691.047,
                        },
                    ],
                },
            ],
            size_limit: 500,
        },
        Floor {
            name: "Supermarket",
            stores: [
                Store {
                    name: "Pretail",
                    square_meters: 950,
                    employees: [
                        Employee {
                            name: "Amara Schaefer",
                            age: 23,
                            working_hours: (
                                9,
                                14,
                            ),
                            salary: 716.589,
                        },
                        Employee {
                            name: "Yara Wickens",
                            age: 39,
                            working_hours: (
                                9,
                                14,
                            ),
                            salary: 768.078,
                        },
                        Employee {
                            name: "Tomi Boyer",
                            age: 64,
                            working_hours: (
                                9,
                                14,
                            ),
                            salary: 793.647,
                        },
                        Employee {
                            name: "Greta Dickson",
                            age: 42,
                            working_hours: (
                                9,
                                14,
                            ),
                            salary: 697.59,
                        },
                        Employee {
                            name: "Caroline Finnegan",
                            age: 41,
                            working_hours: (
                                9,
                                14,
                            ),
                            salary: 632.6279999999999,
                        },
                        Employee {
                            name: "Indiana Baxter",
                            age: 33,
                            working_hours: (
                                13,
                                20,
                            ),
                            salary: 892.539,
                        },
                        Employee {
                            name: "Jadine Page",
                            age: 48,
                            working_hours: (
                                13,
                                20,
                            ),
                            salary: 668.889,
                        },
                        Employee {
                            name: "Husna Ryan",
                            age: 43,
                            working_hours: (
                                13,
                                20,
                            ),
                            salary: 590.175,
                        },
                        Employee {
                            name: "Tyler Hunt",
                            age: 63,
                            working_hours: (
                                13,
                                20,
                            ),
                            salary: 601.425,
                        },
                        Employee {
                            name: "Dahlia Caldwell",
                            age: 56,
                            working_hours: (
                                13,
                                20,
                            ),
                            salary: 703.242,
                        },
                        Employee {
                            name: "Chandler Mansell",
                            age: 20,
                            working_hours: (
                                19,
                                24,
                            ),
                            salary: 591.075,
                        },
                        Employee {
                            name: "Mohsin Mcgee",
                            age: 30,
                            working_hours: (
                                19,
                                24,
                            ),
                            salary: 633.447,
                        },
                        Employee {
                            name: "Antoine Goulding",
                            age: 45,
                            working_hours: (
                                19,
                                24,
                            ),
                            salary: 627.408,
                        },
                        Employee {
                            name: "Mark Barnard",
                            age: 53,
                            working_hours: (
                                19,
                                24,
                            ),
                            salary: 709.929,
                        },
                    ],
                },
            ],
            size_limit: 1000,
        },
    ],
}
$
```
