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

Create a file named `err.rs` which handles the boxing of errors.

This file must:

- Implement an `enum` called `ParseErr` which will represent parsing errors. It must have the following variants:
  - `Empty`
  - `Malformed`: which has a dynamic boxed error as element
- Implement a `struct` called `ReadErr` which will represent reading errors with a single field called `child_err` of type `Box<dyn Error>`.

These structures must implement the `Display` trait. They should write `"Failed to parse todo file"` and `"Failed to read todo file"`, depending on the respective error case.

These structures should also naturally implement the `Error` trait. We will override its method `source` to manipulate the error source output.

- For the `ReadErr` structure, `child_err` should be returned, wrapped in `Some()`.
- For the `ParseErr` structure, `None` should be returned if we have no tasks, otherwise, with a parsing malformation, we should just return `self` wrapped in `Some()`.

In the `lib.rs` file you will have to declare and implement a `TodoList` and a `Task` structure.

- The `Task` structure serves merely to represent and encapsulate the fields of the tasks in the JSON file.
- The `TodoList` structure will have two fields called `title` and `tasks` as shown below, and an associated **function** called `get_todo` which receives a `&str` and returns a `Result` which will represent either the deserialized and parsed content in a `Self` instance on success, or any error type on failure.

### Dependencies

[json = "0.12.4"](https://docs.rs/json/0.12.4/json/)

### Expected Functions

For `err.rs`

```rust
use std::{error::Error, fmt::Display};

pub enum ParseErr {
    // expected public fields
}

impl Display for ParseErr {
}

impl Error for ParseErr {
}

pub struct ReadErr {
    // expected public fields
}

impl Display for ReadErr {
}

impl Error for ReadErr {
}
```

For `lib.rs`

```rust
mod err;

use std::error::Error;

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
        todo!()
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
            println!("{}: {:?}", e.to_string(), e.source());
        }
    }

    let todos = TodoList::get_todo("todo_empty.json");
    match todos {
        Ok(list) => println!("{:?}", list),
        Err(e) => {
            println!("{}: {:?}", e.to_string(), e.source());
        }
    }

    let todos = TodoList::get_todo("malformed_object.json");
    match todos {
        Ok(list) => println!("{:?}", list),
        Err(e) => {
            println!("{}: {:?}", e.to_string(), e.source());
        }
    }
}
```

And its output:

```console
$ cargo run
TodoList { title: "TODO LIST FOR PISCINE RUST", tasks: [Task { id: 0, description: "do this", level: 0 }, Task { id: 1, description: "do that", level: 5 }] }
Failed to parse todo file: None
Failed to parse todo file: Some(Malformed(UnexpectedCharacter { ch: ',', line: 2, column: 18 }))
$
```

### Notions

- [Module std::fmt](https://doc.rust-lang.org/std/fmt/)
- [JSON](https://docs.rs/json/0.12.4/json/)
- [Boxing errors](https://doc.rust-lang.org/stable/rust-by-example/error/multiple_error_types/boxing_errors.html)
- [Returning Traits with dyn](https://doc.rust-lang.org/stable/rust-by-example/trait/dyn.html)
