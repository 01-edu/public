## generics_list 

### Instructions

Create a linked list of generic values with the following methods.

- `new`: returns a new empty list.
- `push`: adds an element to the beginning of the list.
- `pop`: deletes an element from the list based on LIFO.
- `len`: returns the size of the list.

### Expected Functions

```rust
#[derive(Clone, Debug)]
pub struct List<T> {
    pub head: Option<Node<T>>,
}

#[derive(Clone, Debug)]
pub struct Node<T> {
    pub value: T,
    pub next: Option<Box<Node<T>>>,
}

impl<T> List<T> {
    pub fn new() -> List<T> {
    }

    pub fn push(&mut self, value: T) {
    }

    pub fn pop(&mut self) {
    }

    pub fn len(&self) -> usize {
    }
}
```

### Usage

Here is a program to test your function.

```rust
use generics_list::*;

fn main() {
    let mut new_list_str = List::new();
    new_list_str.push("String Test 1");
    println!("The size of the list is {}", new_list_str.len());

    new_list_str.push("String Test 2");
    println!("The size of the list is {}", new_list_str.len());

    new_list_str.push("String Test 3");
    println!("The size of the list is {}", new_list_str.len());

    new_list_str.pop();
    println!("The size of the list is {}", new_list_str.len());
}
```

And its output

```console
$ cargo run
The size of the list is 1
The size of the list is 2
The size of the list is 3
The size of the list is 2
$
```
