## get_document_id

### Instructions

Create the following structures which will help you get the `document_id` of the document you need.

- `OfficeOne`: `next_office` as type `Result<OfficeTwo, ErrorOffice>`.
- `OfficeTwo`: `next_office` as type `Result<OfficeThree, ErrorOffice>`.
- `OfficeThree`: `next_office` as type `Result<OfficeFour, ErrorOffice>`.
- `OfficeFour`: `document_id` as type `Result<u32, ErrorOffice>`.

`ErrorOffice` is an `enum` with `OfficeClosed(u32)`, `OfficeNotFound(u32)` or `OfficeFull(u32)`.

The `u32` is the `id` of the office generating the error.

Beside the structures, you must create a **function** named `get_document_id`, and associate it to the `OfficeOne` structure.

This **function** should return the `Result` value in the `OfficeFour` structure or the first `Err` it finds in the chain.

### Expected Function and structures

```rust
#[derive(Debug, PartialEq, Eq, Clone, Copy)]
pub enum ErrorOffice {
    OfficeClose(u32),
    OfficeNotFound(u32),
    OfficeFull(u32),
}

#[derive(PartialEq, Eq, Clone, Copy)]
pub struct OfficeOne {
    pub next_office: Result<OfficeTwo, ErrorOffice>,
    pub id: u32,
}

#[derive(PartialEq, Eq, Clone, Copy)]
pub struct OfficeTwo {
    pub next_office: Result<OfficeThree, ErrorOffice>,
    pub id: u32,
}

#[derive(PartialEq, Eq, Clone, Copy)]
pub struct OfficeThree {
    pub next_office: Result<OfficeFour, ErrorOffice>,
    pub id: u32,
}

#[derive(PartialEq, Eq, Clone, Copy)]
pub struct OfficeFour {
    pub document_id: Result<u32, ErrorOffice>,
}

impl OfficeOne {
    pub fn get_document_id(&self) -> Result<u32, ErrorOffice> {}
}
```

### Usage

Here is a program to test your function.

```rust
use get_document_id::*;

fn main() {
    let office_ok = OfficeOne {
        next_office: Ok(OfficeTwo {
            next_office: Ok(OfficeThree {
                next_office: Ok(OfficeFour {
                    document_id: Ok(13),
                    id: 4,
                }),
                id: 3,
            }),
            id: 2,
        }),
        id: 1,
    };
    let office_closed = {
        OfficeOne {
            next_office: Ok(OfficeTwo {
                next_office: Err(ErrorOffice::OfficeClose(2)),
                id: 2,
            }),
            id: 1,
        }
    };

    match office_ok.get_document_id() {
        Ok(id) => println!("Found a document with id {}", id),
        Err(err) => match err {
            ErrorOffice::OfficeClose(id) => println!("Error: office {id} closed!"),
            ErrorOffice::OfficeNotFound(id) => println!("Error: office {id} not found!"),
            ErrorOffice::OfficeFull(id) => println!("Error: office {id} full!"),
        },
    };
    match office_closed.get_document_id() {
        Ok(id) => println!("Found a document with id {}", id),
        Err(err) => match err {
            ErrorOffice::OfficeClose(id) => println!("Error: office {id} closed!"),
            ErrorOffice::OfficeotFound(id) => println!("Error: office {id} not found!"),
            ErrorOffice::OfficeFull(id) => println!("Error: office {id} full!"),
        },
    };
}
```

And its output:

```console
$ cargo run
Found a document with id 13
Error: office 2 closed!
$
```
