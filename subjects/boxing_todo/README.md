## boxing_todo

### Instructions

The objective is create an api to parse a list of _todos_ that is organized in a JSON file,
handling all possible errors in a multiple error system.

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

#### Error.rs

Create a module in another file called **error.rs** which handles the boxing of errors.
This module must implement an `enum` called `ParseErr` which will take care of the
parsing errors. It must have the following elements:

- Empty
- Malformed, which has a dynamic boxed error as element

A structure called `ReadErr` which will take care of the reading errors, having just an element called `child_err` of type `Box<dyn Error>`.

For each data structure you will have to implement a function called `fmt` for the trait `Display` which writes
out the message **"Failed to parse todo"** in case it is a parsing error. Otherwise, it should write the message
**"Failed to read todo file"**.
For the `Error` trait the following functions (methods) have to be implemented:

- `source` which returns an `Option` with the error:

  - For the `ReadErr` it must just return the option with the error
  - For the `ParseErr` it will return an option which can be `None` if the tasks are **empty** otherwise the error, if
    the parsing is **malformed**.

#### lib.rs

In the **lib** file you will have to implement a **function** called `get_todo` which receives a string and returns a Result
which can be the structure `TodoList` or a boxing error. This **function** must be able to deserialize the json file.
Basically it must parse and read the JSON file and return the `TodoList` if everything is fine, otherwise it returns the error.

### Notions

- [Module std::fmt](https://doc.rust-lang.org/std/fmt/)
- [JSON](https://docs.rs/json/0.12.4/json/)
- [Boxing errors](https://doc.rust-lang.org/stable/rust-by-example/error/multiple_error_types/boxing_errors.html)
- [Returning Traits wirh dyn](https://doc.rust-lang.org/stable/rust-by-example/trait/dyn.html)

### Expected Functions

For **error.rs**

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
mod error;
use error::{ ParseErr, ReadErr };

pub use json::{parse, stringify};
pub use std::error::Error;

#[derive(Debug, Eq, PartialEq)]
pub struct Task {
    id: u32,
    description: String,
    level: u32,
}

#[derive(Debug, Eq, PartialEq)]
pub struct TodoList {
    title: String,
    tasks: Vec<Task>,
}

impl TodoList {
    pub fn get_todo(path: &str) -> Result<TodoList, Box<dyn Error>> {

    }
}
```

### Usage

Here is a program to test your function.
Note that you can create some todo list your self to test it, but you can find the JSON files that
are being tested [here](https://github.com/01-edu/public/blob/master/subjects/boxing_todo)

```rust
mod lib;
use lib::{ TodoList };

fn main() {
    let todos = TodoList::get_todo("todo.json");
    match todos {
        Ok(list) => println!("{:?}", list),
        Err(e) => {
            println!("{}{:?}", e.to_string(), e.source());
        }
    }

    let todos = TodoList::get_todo("no_todo_list.json");
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
student@ubuntu:~/boxing_todo/test$ cargo run
TodoList { title: "TODO LIST FOR PISCINE RUST", tasks: [Task { id: 0, description: "do this", level: 0 }, Task { id: 1, description: "do that", level: 5 }] }
Todo List parse failed: None
Fail to parses todo Some(Malformed(UnexpectedCharacter { ch: ',', line: 2, column: 18 }))
student@ubuntu:~/boxing_todo/test$
```
