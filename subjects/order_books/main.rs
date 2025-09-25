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
