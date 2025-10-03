## car_rental

### Instructions

Create a struct named `Car` and another one named `RentalBusiness`.

The scope of the exercise will be to modify and operate on the `Car` in the `RentalBusiness` even if this element is not declared as mutable, introducing the concept of interior mutability.

To accomplish that, you will create some methods for `RentalBusiness` that will return the field `car` in many different ways.

### Expected Functions and Structures

```rust
use std::cell::{Ref, RefCell, RefMut};

#[derive(Debug, Default, PartialEq, Eq)]
pub struct Car {
    pub color: String,
    pub plate: String,
}

#[derive(Debug)]
pub struct RentalBusiness {
    pub car: RefCell<Car>,
}

impl RentalBusiness {
    pub fn rent_car(&self) -> Ref<'_, Car> {
        todo!()
    }

    pub fn sell_car(&self) -> Car {
        todo!()
    }

    pub fn repair_car(&self) -> RefMut<'_, Car> {
        todo!()
    }

    pub fn change_car(&self, new_car: Car) {
        todo!()
    }
}
```

### Usage

Here is a program to test your function:

```rust
use car_rental::*;
use std::cell::RefCell;

fn main() {
    let car_rental = RentalBusiness {
        car: RefCell::new(Car {
            color: "red".to_string(),
            plate: "AAA".to_string(),
        }),
    };

    println!("{:?}", car_rental.rent_car());
    println!("{:?}", car_rental.repair_car());

    {
        let mut car = car_rental.repair_car();
        car.color = "blue".to_string();
    }

    println!("{:?}", car_rental.rent_car());

    car_rental.change_car(Car {
        color: "pink".to_string(),
        plate: "WWW".to_string(),
    });

    println!("{:?}", car_rental.rent_car());

    println!("{:?}", car_rental.sell_car());
    println!("{:?}", car_rental.sell_car());
}
```

And its output:

```console
$ cargo run
Car { color: "red", plate: "AAA" }
Car { color: "red", plate: "AAA" }
Car { color: "blue", plate: "AAA" }
Car { color: "pink", plate: "WWW" }
Car { color: "pink", plate: "WWW" }
Car { color: "", plate: "" }
$
```

### Notions

- [std::cell::RefCell](https://doc.rust-lang.org/std/cell/struct.RefCell.html)
- [Struct std::rc::Rc](https://doc.rust-lang.org/std/rc/struct.Rc.html)
