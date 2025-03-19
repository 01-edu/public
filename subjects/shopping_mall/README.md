## shopping_mall

### Instructions

Using the `mall` module provided, create the following **functions** to help run a shopping mall:

- `biggest_store`: receives a `Mall` and returns the `Store` with the most `square_meters`.
- `highest_paid_employee`: receives a `Mall` and returns a vector containing the `Employee`(s) with the highest salary.
- `nbr_of_employees`: receives a `Mall` and returns the number of employees and guards as a `usize`.
- `check_for_securities`: receives a `Mall` and a vector of `Guard`. If there is not at least 1 guard for every 200 square meters of total floor size, a guard should be added to the `Mall.guards`.
- `cut_or_raise`: receives a `Mall`. For each employee, the salary will be raised by 10% if they work for 10 hours or more, else their salary will be decreased by 10%. You can consider that guards are not employees of the mall.

### Expected Function

> You'll need to work out the function signatures for yourself.

### mall.rs

```rust
use std::collections::HashMap;

#[inline]
fn coerce_map<V>(m: HashMap<impl Into<String>, V>) -> HashMap<String, V> {
    m.into_iter().map(|(k, v)| (k.into(), v)).collect()
}

#[derive(Debug, Clone, PartialEq)]
pub struct Mall {
    pub name: String,
    pub guards: HashMap<String, Guard>,
    pub floors: HashMap<String, Floor>,
}

impl Mall {
    pub fn new(
        name: impl Into<String>,
        guards: HashMap<impl Into<String>, Guard>,
        floors: HashMap<impl Into<String>, Floor>,
    ) -> Self {
        Self {
            name: name.into(),
            guards: coerce_map(guards),
            floors: coerce_map(floors),
        }
    }

    pub fn change_name(&mut self, new_name: impl Into<String>) {
        self.name = new_name.into();
    }

    pub fn hire_guard(&mut self, name: impl Into<String>, guard: Guard) {
        self.guards.insert(name.into(), guard);
    }

    pub fn fire_guard(&mut self, name: impl Into<String>) {
        self.guards.remove(&name.into());
    }
}

#[derive(Debug, Clone, Copy, PartialEq)]
pub struct Guard {
    pub age: u32,
    pub years_experience: u32,
}

#[derive(Debug, Clone, PartialEq)]
pub struct Floor {
    pub stores: HashMap<String, Store>,
    pub size_limit: u64,
}

impl Floor {
    pub fn new(stores: HashMap<impl Into<String>, Store>, size_limit: u64) -> Self {
        Self {
            stores: coerce_map(stores),
            size_limit,
        }
    }

    pub fn replace_store(&mut self, store: impl Into<String>, with: Store) {
        self.stores.entry(store.into()).and_modify(|v| *v = with);
    }

    pub fn add_store(&mut self, name: impl Into<String>, store: Store) -> Result<(), ()> {
        let has_space = self.size_limit
            >= self.stores.values().map(|s| s.square_meters).sum::<u64>() + store.square_meters;

        if has_space {
            self.stores.insert(name.into(), store);
            Ok(())
        } else {
            Err(())
        }
    }

    pub fn remove_store(&mut self, name: impl Into<String>) {
        self.stores.remove(&name.into());
    }
}

#[derive(Debug, Clone, PartialEq)]
pub struct Store {
    pub employees: HashMap<String, Employee>,
    pub square_meters: u64,
}

impl Store {
    pub fn new(employees: HashMap<impl Into<String>, Employee>, square_meters: u64) -> Self {
        Self {
            employees: coerce_map(employees),
            square_meters,
        }
    }

    pub fn hire_employee(&mut self, name: impl Into<String>, employee: Employee) {
        self.employees.insert(name.into(), employee);
    }

    pub fn fire_employee(&mut self, name: impl Into<String>) {
        self.employees.remove(&name.into());
    }

    pub fn expand(&mut self, by: u64) {
        self.square_meters += by;
    }
}

#[derive(Debug, Clone, Copy, PartialEq)]
pub struct Employee {
    pub age: u32,
    // The employee works from `working_hours.0` to `working_hours.1`
    pub working_hours: (u32, u32),
    pub salary: f64,
}

impl Employee {
    pub fn birthday(&mut self) {
        self.age += 1;
    }

    pub fn change_workload(&mut self, from: u32, to: u32) {
        self.working_hours = (from, to);
    }

    pub fn raise(&mut self, amount: f64) {
        self.salary += amount;
    }

    pub fn cut(&mut self, amount: f64) {
        self.salary -= amount;
    }
}
```

### Usage

Here is a program to test your function:

```rust
use shopping_mall::*;

fn main() {
    let mut mall = Mall::new(
        "La Vie Funchal",
        [
            (
                "John Oliver",
                Guard {
                    age: 34,
                    years_experience: 7,
                },
            ),
            (
                "Bob Schumacher",
                Guard {
                    age: 53,
                    years_experience: 15,
                },
            ),
        ]
        .into(),
        [
            (
                "Ground Floor",
                Floor::new(
                    [
                        (
                            "Footzo",
                            Store::new(
                                [
                                    (
                                        "Finbar Haines",
                                        Employee {
                                            age: 36,
                                            working_hours: (9, 14),
                                            salary: 650.88,
                                        },
                                    ),
                                    (
                                        "Sienna-Rose Penn",
                                        Employee {
                                            age: 26,
                                            working_hours: (9, 22),
                                            salary: 1000.43,
                                        },
                                    ),
                                ]
                                .into(),
                                50,
                            ),
                        ),
                        (
                            "Swashion",
                            Store::new(
                                [
                                    (
                                        "Abdallah Stafford",
                                        Employee {
                                            age: 54,
                                            working_hours: (8, 22),
                                            salary: 1234.21,
                                        },
                                    ),
                                    (
                                        "Marian Snyder",
                                        Employee {
                                            age: 21,
                                            working_hours: (8, 14),
                                            salary: 831.9,
                                        },
                                    ),
                                ]
                                .into(),
                                43,
                            ),
                        ),
                    ]
                    .into(),
                    300,
                ),
            ),
            (
                "Supermarket",
                Floor::new(
                    [(
                        "Pretail",
                        Store::new(
                            [
                                (
                                    "Yara Wickens",
                                    Employee {
                                        age: 39,
                                        working_hours: (9, 14),
                                        salary: 853.42,
                                    },
                                ),
                                (
                                    "Indiana Baxter",
                                    Employee {
                                        age: 33,
                                        working_hours: (13, 20),
                                        salary: 991.71,
                                    },
                                ),
                                (
                                    "Jadine Page",
                                    Employee {
                                        age: 48,
                                        working_hours: (13, 20),
                                        salary: 743.21,
                                    },
                                ),
                                (
                                    "Tyler Hunt",
                                    Employee {
                                        age: 63,
                                        working_hours: (13, 20),
                                        salary: 668.25,
                                    },
                                ),
                                (
                                    "Mohsin Mcgee",
                                    Employee {
                                        age: 30,
                                        working_hours: (19, 24),
                                        salary: 703.83,
                                    },
                                ),
                                (
                                    "Antoine Goulding",
                                    Employee {
                                        age: 45,
                                        working_hours: (19, 24),
                                        salary: 697.12,
                                    },
                                ),
                                (
                                    "Mark Barnard",
                                    Employee {
                                        age: 53,
                                        working_hours: (19, 24),
                                        salary: 788.81,
                                    },
                                ),
                            ]
                            .into(),
                            950,
                        ),
                    )]
                    .into(),
                    1000,
                ),
            ),
        ]
        .into(),
    );

    // returns the biggest store
    println!("Biggest store: {:#?}", biggest_store(&mall));

    // returns the list with the highest paid employees
    println!("Highest paid employee: {:#?}", highest_paid_employee(&mall));

    // returns the number of employees
    println!("Number of employees: {}", nbr_of_employees(&mall));

    // checks if it is needed to add securities
    check_for_securities(
        &mut mall,
        [
            (
                "Peter Solomons",
                Guard {
                    age: 45,
                    years_experience: 20,
                },
            ),
            (
                "William Charles",
                Guard {
                    age: 32,
                    years_experience: 10,
                },
            ),
            (
                "Leonardo Changretta",
                Guard {
                    age: 23,
                    years_experience: 0,
                },
            ),
            (
                "Vlad Levi",
                Guard {
                    age: 38,
                    years_experience: 8,
                },
            ),
            (
                "Faruk Berkai",
                Guard {
                    age: 40,
                    years_experience: 15,
                },
            ),
            (
                "Chritopher Smith",
                Guard {
                    age: 35,
                    years_experience: 9,
                },
            ),
            (
                "Jason Mackie",
                Guard {
                    age: 26,
                    years_experience: 2,
                },
            ),
            (
                "Kenzie Mair",
                Guard {
                    age: 34,
                    years_experience: 8,
                },
            ),
            (
                "Bentley Larson",
                Guard {
                    age: 33,
                    years_experience: 10,
                },
            ),
            (
                "Ray Storey",
                Guard {
                    age: 37,
                    years_experience: 12,
                },
            ),
        ]
        .map(|(n, d)| (n.to_owned(), d))
        .into(),
    );

    // raises or cuts the salary of every employee
    cut_or_raise(&mut mall);

    println!("{:#?}", mall);
}
```

And its ouput:

```rs
$ cargo run
Biggest store: (
    "Pretail",
    Store {
        employees: {
            "Mohsin Mcgee": Employee {
                age: 30,
                working_hours: (
                    19,
                    24,
                ),
                salary: 703.83,
            },
            "Jadine Page": Employee {
                age: 48,
                working_hours: (
                    13,
                    20,
                ),
                salary: 743.21,
            },
            "Yara Wickens": Employee {
                age: 39,
                working_hours: (
                    9,
                    14,
                ),
                salary: 853.42,
            },
            "Antoine Goulding": Employee {
                age: 45,
                working_hours: (
                    19,
                    24,
                ),
                salary: 697.12,
            },
            "Indiana Baxter": Employee {
                age: 33,
                working_hours: (
                    13,
                    20,
                ),
                salary: 991.71,
            },
            "Mark Barnard": Employee {
                age: 53,
                working_hours: (
                    19,
                    24,
                ),
                salary: 788.81,
            },
            "Tyler Hunt": Employee {
                age: 63,
                working_hours: (
                    13,
                    20,
                ),
                salary: 668.25,
            },
        },
        square_meters: 950,
    },
)
Highest paid employee: [
    (
        "Abdallah Stafford",
        Employee {
            age: 54,
            working_hours: (
                8,
                22,
            ),
            salary: 1234.21,
        },
    ),
]
Number of employees: 13
Mall {
    name: "La Vie Funchal",
    guards: {
        "Ray Storey": Guard {
            age: 37,
            years_experience: 12,
        },
        "Jason Mackie": Guard {
            age: 26,
            years_experience: 2,
        },
        "John Oliver": Guard {
            age: 34,
            years_experience: 7,
        },
        "Chritopher Smith": Guard {
            age: 35,
            years_experience: 9,
        },
        "Peter Solomons": Guard {
            age: 45,
            years_experience: 20,
        },
        "Bob Schumacher": Guard {
            age: 53,
            years_experience: 15,
        },
    },
    floors: {
        "Supermarket": Floor {
            stores: {
                "Pretail": Store {
                    employees: {
                        "Mohsin Mcgee": Employee {
                            age: 30,
                            working_hours: (
                                19,
                                24,
                            ),
                            salary: 633.447,
                        },
                        "Jadine Page": Employee {
                            age: 48,
                            working_hours: (
                                13,
                                20,
                            ),
                            salary: 668.889,
                        },
                        "Yara Wickens": Employee {
                            age: 39,
                            working_hours: (
                                9,
                                14,
                            ),
                            salary: 768.078,
                        },
                        "Antoine Goulding": Employee {
                            age: 45,
                            working_hours: (
                                19,
                                24,
                            ),
                            salary: 627.408,
                        },
                        "Indiana Baxter": Employee {
                            age: 33,
                            working_hours: (
                                13,
                                20,
                            ),
                            salary: 892.539,
                        },
                        "Mark Barnard": Employee {
                            age: 53,
                            working_hours: (
                                19,
                                24,
                            ),
                            salary: 709.929,
                        },
                        "Tyler Hunt": Employee {
                            age: 63,
                            working_hours: (
                                13,
                                20,
                            ),
                            salary: 601.425,
                        },
                    },
                    square_meters: 950,
                },
            },
            size_limit: 1000,
        },
        "Ground Floor": Floor {
            stores: {
                "Swashion": Store {
                    employees: {
                        "Abdallah Stafford": Employee {
                            age: 54,
                            working_hours: (
                                8,
                                22,
                            ),
                            salary: 1357.631,
                        },
                        "Marian Snyder": Employee {
                            age: 21,
                            working_hours: (
                                8,
                                14,
                            ),
                            salary: 748.71,
                        },
                    },
                    square_meters: 43,
                },
                "Footzo": Store {
                    employees: {
                        "Finbar Haines": Employee {
                            age: 36,
                            working_hours: (
                                9,
                                14,
                            ),
                            salary: 585.792,
                        },
                        "Sienna-Rose Penn": Employee {
                            age: 26,
                            working_hours: (
                                9,
                                22,
                            ),
                            salary: 1100.473,
                        },
                    },
                    square_meters: 50,
                },
            },
            size_limit: 300,
        },
    },
}
$
```
