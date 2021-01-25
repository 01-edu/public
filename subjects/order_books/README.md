## order_books

### Instructions

Build a module called `library` with two sub-modules inside it:

- `writers` which contains a structure called `Writer` that has a first_name (String), last_name (String) and a set of books (Vec\<Book\>).
- `books` which contains a structure called `Book` that has a title (String) and a year of publish (u64).

You will also have to create (outside the previous modules) a function `order_books` that receives a writer (Writer) and orders the set of books alphabetically.

### Expected Functions and Structs

```rs
pub fn order_books(writer: &mut Writer) {

}
```

```rs
struct Writer {

}
```

```rs
struct Book {

}
```

### Example

Here is a program to test your function and structs:

```rs
fn main() {
    let mut writer_a = Writer {
        first_name: "William".to_string(),
        last_name: "Shakespeare".to_string(),
        books: vec![
            Book {
                title: "Hamlet".to_string(),
                year: 1600,
            },
            Book {
                title: "Othelo".to_string(),
                year: 1603,
            },
            Book {
                title: "Romeo and Juliet".to_string(),
                year: 1593,
            },
            Book {
                title: "MacBeth".to_string(),
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
