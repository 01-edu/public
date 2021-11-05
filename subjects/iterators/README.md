## iterators

### Instructions

Create a method `new` that takes one number `usize` and initializes the `Number` struct.
This method will have to determinate if the given number is even or odd.

After that you will implement an iterator for the struct `Number` that returns a tuple (usize,usize,usize) containing each field of the struct Number.

The first position of the tuple will have the even numbers, the second will have the odd numbers, and the third will have the factorial numbers.

That being said, You will have to follow this rules:

- If the number is even you will have to calculate the factorial of the next odd number.
- If the number is odd you will have to calculate the factorial of the next even number.
- The tupple has to return all the numbers in the right positions.

Example:

If `a = 5`, 5 is odd, the next even is 6, so the factorial of 6 is 720. Your program will return the following:

```console
Some((6, 5, 720))
```

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
    let mut a = Number::new(4);
    println!("{:?}", a.next());
    println!("{:?}", a.next());
    println!("{:?}", a.next());
    println!("{:?}", b.next());
    println!("{:?}", b.next());
}
```

And its output:

```console
$ cargo run
Some((6, 5, 720))
Some((6, 7, 5040))
Some((8, 7, 40320))
Some((8, 9, 362880))
Some((10, 9, 3628800))
$
```
