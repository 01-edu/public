## iterators

### Instructions

Create a method `new` that takes one number `usize` and initializes the struct `Number`.

This method will have to determinate if the given number is even or odd. If it is even you will have to increment it by one to the next odd number and if it is odd you have to increment by one to the next even number.

After that you will implement an iterator for the struct `Number` that returns a tuple (usize,usize,usize) containing each field of the struct Number.

The first position of the tuple will be the even number, the second will be the odd number, and the third will be the factorial number.

So the purpose is to return the given number in the right position. If it is even it will be at the first position, and if it is odd it will be in the second position. Apart from that you have to return the factorial of the given number in the third position.

### Notions

- [Trait Iterator](https://doc.rust-lang.org/std/iter/trait.Iterator.html)

### Expected functions

```rust
impl Number {
    pub fn new(nbr: usize) -> Number {

    }
}

impl Iterator for Number {
    pub fn next(&mut self) -> Option<Self::Item> {}
}
```

### Usage

Here is a program to test your function.

```rust
use iterators::*;

struct Number {
    even: usize,
    odd: usize,
    fact: usize,
}

fn main() {
    let mut a = Number::new(5);
    println!("{:?}", a.next());
    println!("{:?}", a.next());
    println!("{:?}", a.next());
    println!();
    let mut b = Number::new(18);
    println!("{:?}", b.next());
    println!("{:?}", b.next());
    println!("{:?}", b.next());
}
```

And its output:

```console
$ cargo run
Some((6, 5, 120))
Some((8, 7, 720))
Some((10, 9, 5040))

Some((18, 19, 6402373705728000))
Some((20, 21, 121645100408832000))
Some((22, 23, 2432902008176640000))
$
```
