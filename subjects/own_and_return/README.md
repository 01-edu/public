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
    todo!()
}

pub fn take_film_name(/* to be implemented */) -> String {
    todo!()
}

```

### Usage

Here is a possible program to test your function:

```rust
use own_and_return::*;

fn main() {
    let my_film = Film {
        name: "Terminator".to_owned(),
    };

    // println!("{}", take_film_name(/* to be implemented */));

    println!("{}", read_film_name(/* to be implemented */));
    println!("{}", take_film_name(/* to be implemented */));

    // the order of the print statements is intentional.
    // you can test this exercise properly by uncommenting out the first print statement,
    // you should get a compilation error then if your implementation is correct.
}
```

And its output:

```console
$ cargo run
Terminator
Terminator
$
```
