## iterators

### Instructions

The Collatz Conjecture or 3x+1 problem can be summarized as follows:

Take any positive integer `n`.

- If `n` is even, you will divide `n` by 2 to get `n / 2`.
- If `n` is odd, you will multiply `n` by 3 and add 1 to get `3n + 1`.

Repeat the process indefinitely. The conjecture states that no matter which number you start with, you will always reach 1 eventually.

But sometimes the number grow significantly before it reaches 1. This can lead to an integer overflow and makes the algorithm unsolvable within the range of a number in u64. You will not have to worry about that in this exercise.

Given a number `n`, return the number of steps required to reach 1.

Examples:

Starting with n = 16, the steps would be as follows:

0. 16
1. 8
2. 4
3. 2
4. 1

Resulting in 4 steps. So for input n = 16, the return value would be 4.

### Notions

- [Trait Iterator](https://doc.rust-lang.org/std/iter/trait.Iterator.html)
- [Collatz Conjecture](https://en.wikipedia.org/wiki/Collatz_conjecture)

### Expected functions

```rust
#[derive(Copy, Clone)]
pub struct Collatz {
    pub v: u64,
}

impl Iterator for Collatz {}

impl Collatz {
	pub fn new(n: u64) -> Self {}
}

pub fn collatz(n: u64) -> usize {}
```

### Usage

Here is a program to test your function.

```rust
use iterators::*;

fn main() {
    println!("{:?}", collatz(0));
    println!("{:?}", collatz(1));
    println!("{:?}", collatz(4));
    println!("{:?}", collatz(5));
    println!("{:?}", collatz(6));
    println!("{:?}", collatz(7));
    println!("{:?}", collatz(12));
}
```

And its output:

```console
$ cargo run
0
0
2
5
8
16
9
$
```
