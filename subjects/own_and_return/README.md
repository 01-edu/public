## own_and_return

### Instructions

Create a struct `Film` that has one field `name` of type `String`.

Create the following two functions:

- Create a function `take_film_name`, it will return the name and consume the film (you should not be able to reuse it after you passed it to the function).
- Create a function `read_film_name`, it will return the name without consuming the film (you can call the function multiple times with the same argument).

### Expected functions

```rust
pub struct Film {
    pub name: String,
}

pub fn read_film_name(/* to be implemented */) -> String {
}

pub fn take_film_name(/* to be implemented */) -> String {
}

```

### Usage

Here is a possible program to test your function:

```rust
use own_and_return::*;

pub struct Film {
    pub name: String,
}

fn main() {
    let my_film = Film { name: "Terminator".toString() };
    println!("{}", take_film_name(/* to be implemented */));
    // the order of the print statements is intentional, if your implementation is correct,
    // you should have a compile error because my_film was consumed
    println!("{}", read_film_name(/* to be implemented */));
    println!("{}", take_film_name(/*to be implemented*/))
    // you can test this function by commenting out the first print statement,
    // you should see the expected output without errors in this case
}

```

And its output:

```console
$ cargo run
Terminator
Terminator
$
```
