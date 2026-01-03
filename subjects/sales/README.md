## sales

### Instructions

You are going to make a shopping system. It will have a store where the products will be saved, and a cart which will contain items from which a receipt will be generated.

**"Buy three, get one free".**

The store is having a promotion. For every three items, the cheapest will be free. But there is a problem with the printer interface: It cannot receive any zero values. We can create a workaround. We will modify the price of all of the items in the cart by a small percentage that will add up to the correct total price. You can find examples below.

You will have to implement for the `Cart` structure the following **methods**:

- `new`: will initialize the cart.
- `insert_item_by_name`: will receive a reference to `Store` and a `&str`. As the name suggests, it will insert into the cart the item of the store with the same name. If such item does not exist, it will return `Err(())`. On success, we return `Ok(())`.
- `generate_receipt`: returns a vector of the final prices, sorted. This function must generate the receipt just like the example below, with the promotion applied.

### Expected Function

```rust
#[derive(Clone, Copy, Debug, PartialEq)]
pub struct Item(pub &str, pub f64);

#[derive(Clone, Copy, Debug, PartialEq)]
pub struct Store {
    pub products: [Item; _],
}

impl Store {
    pub fn new(products: [Item; _]) -> Self {
        todo!()
    }
}

#[derive(Debug, Clone, PartialEq)]
pub struct Cart {
    pub items: Vec<Item>,
}

impl Cart {
    pub fn new() -> Self {
        todo!()
    }

    pub fn insert_item_by_name(&mut self, s: &Store, name: &str) {
        todo!()
    }

    pub fn generate_prices(&self) -> Vec<f64> {
        todo!()
    }
}
```

### Example

```
[1.23, 3.12, 23.1]` => `[1.17, 2.98, 22.06]
```

Because `1.17 + 2.98 + 22.06` == `0 + 3.12 + 23.1`

**Floats are rounded with a precision of two decimals** which can create small discrepancies as per the example above.

This is a percentage calculation, and it can be applied to a set of three items. If the client purchases 9 items, they will receive three for free, with the discount applied to all items.

```
[1.23, 23.1, 3.12, 9.75, 1.75, 23.75, 2.75, 1.64, 15.23] => [1.16, 1.55, 1.65, 2.6, 2.94, 9.2, 14.38, 21.8, 22.42]
```

```
[3.12, 9.75, 1.75, 23.75, 2.75, 1.64, 15.23] => [1.54, 1.65, 2.59, 2.94, 9.18, 14.34, 22.36]
```

> Hint: Closures will be particularly helpful for many of these operations inside lists and vectors.

### Usage

Here is a program to test your function,

```rust
use sales::*;

fn main() {
    let store = Store::new([
        Item("product A", 1.23),
        Item("product B", 23.1),
        Item("product C", 3.12)
    ]);

    println!("{:?}", store);

    let mut cart = Cart::new();
    cart.insert_item_by_name(&store, "product A").unwrap();
    cart.insert_item_by_name(&store, "product B").unwrap();
    cart.insert_item_by_name(&store, "product C").unwrap();

    println!("{:?}", cart);

    println!("{:?}", cart.generate_receipt());
}
```

And its output:

```console
$ cargo run
Store { products: [Item("product A", 1.23), Item("product B", 23.1), Item("product C", 3.12)] }
Cart { items: [Item("product A", 1.23), Item("product B", 23.1), Item("product C", 3.12)] }
[1.17, 2.98, 22.06]
$
```

### Notions

- [closures](https://doc.rust-lang.org/rust-by-example/fn/closures.html)
