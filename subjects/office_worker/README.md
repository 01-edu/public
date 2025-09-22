## office_worker

### Instructions

Create a structure `OfficeWorker` with the following public fields:

- `name` as `String`.
- `age` as `u32`.
- `role` as `WorkerRole`.

Create an enum `WorkerRole` which can be `Admin`, `User` or `Guest`.

Implement the trait `From<&str>` for both `OfficeWorker` and `WorkerRole`.
For `OfficeWorker` the string will have the format `{name},{age},{role}`.
For `WorkerRole` the format of the string will be `{role}` in lowercase.

> Invalid inputs won't be tested.

### Expected Functions and Data Structures

```rust
#[derive(Debug, PartialEq, Eq)]
pub struct OfficeWorker {
}

#[derive(Debug, PartialEq, Eq)]
pub enum WorkerRole {
}

impl From<&str> for OfficeWorker {
    fn from(s: &str) -> Self {
        todo!()
    }
}

impl From<&str> for WorkerRole {
    fn from(s: &str) -> Self {
        todo!()
    }
}
```

### Usage

Here is a program to test your function.

```rust
use office_worker::*;

fn main() {
    println!("New worker: {:?}", OfficeWorker::from("Manuel,23,admin"));
    println!(
        "New worker: {:?}",
        OfficeWorker::from("Jean Jacques,44,guest")
    );
}
```

And its output:

```console
$ cargo run
New worker: OfficeWorker { name: "Manuel", age: 23, role: Admin }
New worker: OfficeWorker { name: "Jean Jacques", age: 44, role: Guest }
$
```
