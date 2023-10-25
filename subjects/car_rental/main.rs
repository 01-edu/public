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
