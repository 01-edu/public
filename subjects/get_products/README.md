## get_products

### Instructions

Create a function `get_products` that takes a vector of integers, and returns a vector of the products
of each index. For this exercise to be correct you will have to return the product of every index
except the current one.

Examples: [1,2,3,4]

- for the number `1` we get `2*3*4 = 24`
- for the number `3` we get `1*2*4 = 8`

### Notions

- [Trait iterator](https://doc.rust-lang.org/std/iter/trait.Iterator.html)

### Expected functions

```rust
pub fn get_products(arr: Vec<usize>) -> Vec<usize> {

}
```

### Usage

Here is a program to test your function.

```rust
use get_products::get_products;

fn main() {
    let arr: Vec<usize> = vec![1, 7, 3, 4];
    let output = get_products(arr);
    println!("{:?}", output);
}
```

And its output:

```console
student@ubuntu:~/get_products/test$ cargo run
[84, 12, 28, 21]
student@ubuntu:~/get_products/test$
```
