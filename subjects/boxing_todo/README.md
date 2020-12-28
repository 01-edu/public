## boxing_todo

### Instructions

The objective is to do an api to parse a list of *todos* that is organized in a JSON file,
handling all possible errors in a multiple error system.

Organization of the JSON file:

```json
{
    "title" : "TODO LIST FOR PISCINE RUST",
    "tasks": [
        { "id": 0, "description": "do this", "level": 0 },
        { "id": 1, "description": "do that", "level": 5 }
    ]
}
```

Create a module in another file called **error.rs** that handles the boxing of errors.
This module must implement an `enum` called `ParseErr` that will take care of the
parsing errors, it must have the following elements:

- Empty
- Malformed, that has a dynamic boxed error as element

A structure called `ReadErr` that will take care of the reading errors, having just an element called `child_err` of type `Box<dyn Error>`.

For each data structure you will have to implement a function called `fmt` for the trait `Display` that writes
out the message **"Fail to parse todo"** in case it's a parsing error, otherwise it writes the message
**"Failed to read todo file"**.
And for the `Error` trait the following functions:

- `description` that returns a string literal that says:
  - "Todo List parse failed: " for the `ParseErr`
  - "Todo List read failed: " for the `ReadErr`.

- `cause` that returns an `Option` with the error:
  - For the `ReadErr` it must just return the option with the error
  - For the `ParseErr` it will return an option that can be `None` if the tasks are **empty** otherwise the error, if
  the parsing is **malformed**.

In the **lib** file you will have to implement a function called `get_todo` that receives a string and returns a Result
that can be the structure `TodoList` or a boxing error. This function must be able to deserialize the json file,
basically it must parse and read the JSON file and return the `TodoList` if everything is fine otherwise the error.

### Expected Function

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

impl Error for ReadErr {
    fn description(&self) -> &str {}
    fn cause(&self) -> Option<&dyn Error> {}
}

impl Error for ParseErr {
    fn description(&self) -> &str {}
    fn cause(&self) -> Option<&dyn Error> {}
}

```

for **lib.rs**

```rust
mod error;
use error::{ ParseErr, ReadErr };
use std::error::Error;
use serde::{ Deserialize, Serialize };

#[derive(Serialize, Deserialize, Debug, Eq, PartialEq)]
pub struct Task {
    id: u32,
    description: String,
    level: u32,
}

#[derive(Serialize, Deserialize, Debug, Eq, PartialEq)]
pub struct TodoList {
    title: String,
    tasks: Vec<Task>,
}

impl TodoList {
    pub fn get_todo(path: &str) -> Result<TodoList, Box<dyn Error>> {}
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
            println!("{}{:?}", e.description(), e.cause());
        }
    }

    let todos = TodoList::get_todo("no_todo_list.json");
    match todos {
        Ok(list) => println!("{:?}", list),
        Err(e) => {
            println!("{}{:?}", e.description(), e.cause());
        }
    }

    let todos = TodoList::get_todo("malformed_object.json");
    match todos {
        Ok(list) => println!("{:?}", list),
        Err(e) => {
            println!("{}{:?}", e.description(), e.cause().unwrap());
        }
    }
}
```

And its output:

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
TodoList { title: "TODO LIST FOR PISCINE RUST", tasks: [Task { id: 0, description: "do this", level: 0 }, Task { id: 1, description: "do that", level: 5 }] }
Todo List parse failed: None
Todo List parse failed: Malformed(Error("missing field `title`", line: 1, column: 2))
student@ubuntu:~/[[ROOT]]/test$
```

### Notions

- https://serde.rs/
- https://doc.rust-lang.org/stable/rust-by-example/error/multiple_error_types/boxing_errors.html
- https://doc.rust-lang.org/stable/rust-by-example/trait/dyn.html
