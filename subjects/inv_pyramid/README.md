## inv_pyramid

### Instructions

Create a function named `inv_pyramid` that takes a string and an integer as input and returns a vector of strings.
This function should create a pyramid structure. Each element of the vector must be the given string after indentation represented as spaces.

### Example

For i = 5,

```console
[
    " >",
    "  >>",
    "   >>>",
    "    >>>>",
    "     >>>>>",
    "    >>>>",
    "   >>>",
    "  >>",
    " >"
]
```

### Expected Functions

```rust
pub fn inv_pyramid(v: String, i: usize) -> Vec<String> {
    todo!()
}
```

### Usage

Here is a program to test your function

```rust
use inv_pyramid::*;

fn main() {
    println!("{:#?}", inv_pyramid(String::from("#"), 1));
    println!("{:#?}", inv_pyramid(String::from("a"), 2));
    println!("{:#?}", inv_pyramid(String::from(">"), 5));
    println!("{:#?}", inv_pyramid(String::from("&"), 8));
}
```

And its output

```console
$ cargo run
[
    " #",
]
[
    " a",
    "  aa",
    " a",
]
[
    " >",
    "  >>",
    "   >>>",
    "    >>>>",
    "     >>>>>",
    "    >>>>",
    "   >>>",
    "  >>",
    " >",
]
[
    " &",
    "  &&",
    "   &&&",
    "    &&&&",
    "     &&&&&",
    "      &&&&&&",
    "       &&&&&&&",
    "        &&&&&&&&",
    "       &&&&&&&",
    "      &&&&&&",
    "     &&&&&",
    "    &&&&",
    "   &&&",
    "  &&",
    " &",
]
$
```
