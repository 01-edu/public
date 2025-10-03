## vehicles

### Instructions

At a border crossing, you want to keep a list of all the vehicles that are waiting to enter the country. You intend to keep a waiting list of the vehicles, but the vehicles can be one of two types: `Car` or `Truck`.

Create a function that receives an array of references to structures that implements the `Vehicle` trait, and returns all the models in an array of same size.

> Once again, ensure that all the function signatures and their respective generics and lifetimes are correct.

### Expected Functions and Structures

```rust
pub struct Car<'a> {
    pub plate_nbr: &'a str,
    pub model: &'a str,
    pub horse_power: u32,
    pub year: u32,
}

pub struct Truck<'a> {
    pub plate_nbr: &'a str,
    pub model: &'a str,
    pub horse_power: u32,
    pub year: u32,
    pub load_tons: u32,
}

pub trait Vehicle {
    fn model(&self) -> &str;
    fn year(&self) -> u32;
}

impl Vehicle for Truck<'_> {

}

impl Vehicle for Car<'_> {

}

pub fn all_models<N>(list: [_; N]) -> [&str; N] {
    todo!()
}
```

### Usage

Here is a program to test your function.

```rust
use border_cross::*;

fn main() {
    println!(
        "{:?}",
        all_models([
            &Car {
                plate_nbr: "A3D5C7",
                model: "Model 3",
                horse_power: 325,
                year: 2010,
            },
            &Truck {
                plate_nbr: "V3D5CT",
                model: "Ranger",
                horse_power: 325,
                year: 2010,
                load_tons: 40,
            },
        ])
    );
}
```

And its output

```console
$ cargo run
["Model 3", "Ranger"]
$
```
