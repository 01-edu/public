## order_books

### Instructions

Build a module named `library` which contains two sub-modules:

- `writers`: which contains:
  - `Writer`: a structure with:
    - `first_name`: `String`
    - `last_name`: `String`
    - `books`: `Vec<Book>`
- `books`: which contains:
  - `Book`: a structure with:
    - `title`: `String`
    - `year`: `u32` as its year of publication

A function `order_books` should be created outside of the previous modules which receives a `Writer`, and orders the set of books alphabetically (case insensitive!).

### Expected Functions and Structs

> You'll need to complete the function and structs, and add them to the appropriate place, so that the `main` in the usage can be successfully compiled and run.

```rs
pub struct Writer {

}
```

```rs
pub struct Book {

}
```

```rs
pub fn order_books(writer: &mut Writer) {
    todo!()
}
```

### Example

Here is a program to test your function and structs:

```rs
pub use order_books::{
    library::{books::Book, writers::Writer},
    order_books,
};

fn main() {
    let mut writer_a = Writer {
        first_name: "William".to_owned(),
        last_name: "Shakespeare".to_owned(),
        books: vec![
            Book {
                title: "Hamlet".to_owned(),
                year: 1600,
            },
            Book {
                title: "Othelo".to_owned(),
                year: 1603,
            },
            Book {
                title: "Romeo and Juliet".to_owned(),
                year: 1593,
            },
            Book {
                title: "MacBeth".to_owned(),
                year: 1605,
            },
        ],
    };

    println!("Before ordering");
    for b in &writer_a.books {
        println!("{:?}", b.title);
    }

    order_books(&mut writer_a);

    println!("\nAfter ordering");
    for b in writer_a.books {
        println!("{:?}", b.title);
    }
}
```

And its output:

```sh
$ cargo run
Before ordering
"Hamlet"
"Othelo"
"Romeo and Juliet"
"MacBeth"

After ordering
"Hamlet"
"MacBeth"
"Othelo"
"Romeo and Juliet"
$
```
