## boxing_todo

### Instructions

The objective is to create an API to parse a list of _todos_ that are organized in a JSON file. You must handle all possible errors in a multiple error system.

Organization of the JSON file:

```json
{
  "title": "TODO LIST FOR PISCINE RUST",
  "tasks": [
    { "id": 0, "description": "do this", "level": 0 },
    { "id": 1, "description": "do that", "level": 5 }
  ]
}
```

#### err.rs

Create a module in a file named **err.rs** which handles the boxing of errors.

This module must implement an `enum` called `ParseErr` which will take care of the parsing errors. It must have the following elements:

- `Empty`
- `Malformed`: which has a dynamic boxed error as element

A structure called `ReadErr` which will take care of the reading errors, with an element called `child_err` of type `Box<dyn Error>`.

For each data structure, you will have to implement a function called `fmt` for the `Display` trait. It should write the message **"Fail to parse todo"** in the case of any parsing error. Otherwise, it should write the message **"Fail to read todo file"**.

For the `Error` trait, the following functions (methods) have to be implemented:

- `source` which returns an `Option` with the error:

  - For the `ReadErr`, it must return the option with the error.
  - For the `ParseErr`, it will return an option which is `None` if the tasks are **empty**, and the error if the parsing is **malformed**.

#### lib.rs

In the **lib** file you will have to implement a **function** called `get_todo` which receives a string and returns a `Result` which can be the structure `TodoList` or a boxing error. This **function** must be able to deserialize the json file.

Basically it must parse and read the JSON file and return the `TodoList` if everything is fine, otherwise it returns the error.

### Dependencies

[json = "0.12.4"](https://docs.rs/json/0.12.4/json/)

### Expected Functions

For **err.rs**

```rust
use std::fmt;
use std::fmt::Display;
use std::error::Error;

pub enum ParseErr {
    // expected public fields
}

// required by error trait
impl Display for ParseErr {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {

    }
}

pub struct ReadErr {
    // expected public fields
}

// required by error trait
impl Display for ReadErr {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {

    }
}

impl Error for ParseErr {
    fn source(&self) -> Option<&(dyn Error + 'static)> {

    }
}

impl Error for ReadErr {
    fn source(&self) -> Option<&(dyn Error + 'static)> {

    }
}
```

for **lib.rs**

```rust
mod err;
use err::{ ParseErr, ReadErr };

pub use json::{parse, stringify};
pub use std::error::Error;

#[derive(Debug, Eq, PartialEq)]
pub struct Task {
    pub id: u32,
    pub description: String,
    pub level: u32,
}

#[derive(Debug, Eq, PartialEq)]
pub struct TodoList {
    pub title: String,
    pub tasks: Vec<Task>,
}

impl TodoList {
    pub fn get_todo(path: &str) -> Result<TodoList, Box<dyn Error>> {

    }
}
```

### Usage

Here is a program to test your function.

You can create some _todos_ yourself, inside the boxing_todo file in order to test it. The JSON structure can be found above.

```rust
use boxing_todo::TodoList;
fn main() {
    let todos = TodoList::get_todo("todo.json");
    match todos {
        Ok(list) => println!("{:?}", list),
        Err(e) => {
            println!("{}{:?}", e.to_string(), e.source());
        }
    }
    let todos = TodoList::get_todo("todo_empty.json");
    match todos {
        Ok(list) => println!("{:?}", list),
        Err(e) => {
            println!("{}{:?}", e.to_string(), e.source());
        }
    }
    let todos = TodoList::get_todo("malformed_object.json");
    match todos {
        Ok(list) => println!("{:?}", list),
        Err(e) => {
            println!("{}{:?}", e.to_string(), e.source());
        }
    }
}
```

And its output:

```console
$ cargo run
TodoList { title: "TODO LIST FOR PISCINE RUST", tasks: [Task { id: 0, description: "do this", level: 0 }, Task { id: 1, description: "do that", level: 5 }] }
Fail to parse todoNone
Fail to parse todo Some(Malformed(UnexpectedCharacter { ch: ',', line: 2, column: 18 }))
$
```

### Notions

- [Module std::fmt](https://doc.rust-lang.org/std/fmt/)
- [JSON](https://docs.rs/json/0.12.4/json/)
- [Boxing errors](https://doc.rust-lang.org/stable/rust-by-example/error/multiple_error_types/boxing_errors.html)
- [Returning Traits with dyn](https://doc.rust-lang.org/stable/rust-by-example/trait/dyn.html)
