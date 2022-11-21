## office_worker

### Instructions

Create a structure `OfficeWorker` having the following public fields:

- `name` as `String`.
- `age` as `u32`.
- `role` as `WorkerRole`.

Create an enum `WorkerRole` which can be `Admin`, `User` or `Guest`.

Implement for both the trait `From<&str>`. For `OfficeWorker` the string will have the format `"name,age,role"`, for `WorkerRole` the format of the string will be the `"role"` name in lower case.

> Invalid inputs won't be tested.

### Expected Functions and Data Structures

```rust
use crate::OfficeWorker::*;

#[derive(Debug, PartialEq, Eq)]
pub struct OfficeWorker {
}

#[derive(Debug, PartialEq, Eq)]
pub enum WorkerRole {
}

impl From<&str> for OfficeWorker {
}

impl From<&str> for WorkerRole {
}
```

### Usage

Here is a program to test your function.

```rust
use office_worker::*;

fn main() {
    println!("New worker: {:?}",
        OfficeWorker::from("Manuel,23,admin"));
    println!("New worker: {:?}",
        OfficeWorker::from("Jean Jacques,44,guest"));
}
```

And its output:

```console
$ cargo run
New worker: OfficeWorker { name: "Manuel", age: 23, role: Admin }
New worker: OfficeWorker { name: "Jean Jacques", age: 44, role: Guest }
$
```
