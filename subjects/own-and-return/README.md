## own_and_return

### Instructions

Create a function `only_return` that takes a `struct` Film that has one Field of type `String` and returns the string without owning it.

Create a function `own_and_return` which takes ownership of the struct, and returns the string.

### Expected functions

```rust
pub struct Film {
    pub name: String,
}

pub fn only_return(your_parameter) {
}

pub fn own_and_return(your_parameter) {
}


```

### Usage

Here is a possible program to test your function :

```rust
use own_and_return::*;

pub struct Film {
    pub name: String,
}

fn main() {
     println!("{}",  own_and_return(your_argument);
     println!("{}",  only_return(your_argument);
}
```

And its output should return the name of the film being passed.
