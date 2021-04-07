## sales

### Instructions

In this exercise a shopping system will have to be created. There will be :

- A store that will save all the products in it
- A cart that will have `items`, that the client will buy, and a `receipt`

This store is having a promotion, "Buy three and get one for free" (the free item must be the cheapest). The receipt must not present
any value as 0, so the promotion must be a reduction to be applied to all items instead.(see the example)

### Notions

- [closures](https://doc.rust-lang.org/rust-by-example/fn/closures.html)

### Expected Function

```rust

#[derive(Debug, Clone)]
pub struct Store {
    pub products: Vec<(String, f32)>,
}
impl Store {
    pub fn new(products: Vec<(String, f32)>) -> Store {
        Store { products }
    }
}

#[derive(Debug, Clone)]
pub struct Cart {
    // expected public fields
}
impl Cart {
    pub fn new() -> Cart {}
    pub fn insert_item(&mut self, s: &Store, ele: String) {}
    pub fn get_prices(&self) -> Vec<f32> {}
    pub fn generate_receipt(&mut self) -> Vec<f32> {}
}

```

### Example

`[1.23, 3.12, 23.1]` -> the receipt will be `[1.17, 2.98, 22.07]`

So `1.17 + 2.98 + 22.07 == 3.12 + 23.1 + 0`

This is a percentage calculation, and it can be applied to a set of three items.
If the client purchase 9 items, the promotion will be applied, three for free, to all items

|--------------| |---------------| |---------------|
`[1.23, 23.1, 3.12, 9.75, 1.75, 23.75, 2.75, 1.64, 15.23]` -> the receipt will be `[1.16, 1.55, 1.65, 2.6, 2.94, 9.2, 14.38, 21.8, 22.42]`

|--------| |--------| |--------|
`[3.12, 9.75, 1.75, 23.75, 2.75, 1.64, 15.23]` -> the receipt will be `[1.54, 1.65, 2.59, 2.94, 9.18, 14.34, 22.36]`

and so on... (hint: Closures is the way)

You will have to implement for the Cart structure the following **functions**:

- `new`, that will initialize the cart
- `insert_item`, that will receive a reference to `Store` and a `String`. Just like the name says you will
  have to insert the item to the cart
- `generate_receipt`, that returns a vector of sorted floats. This function must generate the receipt just
  like the example above, using the promotion. Also saving the result in the filed `receipt`.

### Usage

Here is a program to test your function,

```rust
use sales::*;

fn main() {
    let store = Store::new(vec![
        (String::from("product A"), 1.23),
        (String::from("product B"), 23.1),
        (String::from("product C"), 3.12)]);

    println!("{:?}", store);

    let mut cart = Cart::new();
    cart.insert_item(&store, String::from("product A"));
    cart.insert_item(&store, String::from("product B"));
    cart.insert_item(&store, String::from("product C"));

    println!("{:?}", cart.generate_receipt());

    println!("{:?}", cart);
}
```

And its output:

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
Store { products: [("product A", 1.23), ("product B", 23.1), ("product C", 3.12)] }
[1.17, 2.98, 22.07]
Cart { items: [("product A", 1.23), ("product B", 23.1), ("product C", 3.12)], receipt: [1.17, 2.98, 22.07] }
student@ubuntu:~/[[ROOT]]/test$
```
