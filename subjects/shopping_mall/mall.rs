#[derive(Debug, Clone, PartialEq)]
pub struct Mall {
    pub name: String,
    pub securities: Vec<security::Security>,
    pub floors: Vec<floor::Floor>,
}

impl Mall {
    pub fn new(name: &str, securities: Vec<security::Security>, floors: Vec<floor::Floor>) -> Mall {
        Mall {
            name: name.to_string(),
            securities: securities,
            floors: floors,
        }
    }

    #[warn(dead_code)]
    pub fn change_name(&mut self, new_name: &str) {
        self.name = new_name.to_string();
    }

    #[warn(dead_code)]
    pub fn hire_security(&mut self, security: security::Security) {
        self.securities.push(security);
    }

    #[warn(dead_code)]
    pub fn fire_security(&mut self, name: String) {
        self.securities.retain(|x| x.name != name);
    }
}

pub mod security {

    #[derive(Debug, Clone, PartialEq)]
    pub struct Security {
        pub name: String,
        pub age: u8,
        pub years_experience: u8,
    }

    impl Security {
        pub fn new(name: &str, age: u8, years_experience: u8) -> Security {
            Security {
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
        pub fn new(name: &str, stores: Vec<store::Store>, store_limit: u64) -> Floor {
            Floor {
                name: name.to_string(),
                stores: stores,
                size_limit: store_limit,
            }
        }

        #[warn(dead_code)]
        pub fn change_store(&mut self, store: &str, new_store: store::Store) {
            let pos = self.stores.iter().position(|x| x.name == store).unwrap();
            self.stores[pos] = new_store;
        }

        #[warn(dead_code)]
        pub fn add_store(&mut self, new_store: store::Store) {
            let mut current_floor_size = 0;

            for store in self.stores.iter() {
                current_floor_size += store.square_meters;
            }

            if self.size_limit < current_floor_size + new_store.square_meters {
                self.stores.push(new_store);
            }
        }

        #[warn(dead_code)]
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
            pub fn new(name: &str, space: u64, employees: Vec<employee::Employee>) -> Store {
                Store {
                    name: name.to_string(),
                    square_meters: space,
                    employees: employees,
                }
            }

            #[warn(dead_code)]
            pub fn hire_employee(&mut self, employee: employee::Employee) {
                self.employees.push(employee);
            }
            #[warn(dead_code)]
            pub fn fire_employee(&mut self, employee_name: &str) {
                self.employees.retain(|x| x.name != employee_name);
            }
            #[warn(dead_code)]
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

                #[warn(dead_code)]
                pub fn birthday(&mut self) {
                    self.age += 1;
                }

                #[warn(dead_code)]
                pub fn change_workload(&mut self, entry_hour: u8, exit_hour: u8) {
                    self.working_hours = (entry_hour, exit_hour);
                }

                #[warn(dead_code)]
                pub fn raise(&mut self, amount: f64) {
                    self.salary += amount;
                }

                #[warn(dead_code)]
                pub fn cut(&mut self, amount: f64) {
                    self.salary = self.salary - amount;
                }
            }
        }
    }
}
