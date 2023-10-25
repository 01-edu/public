use display_table::*;

fn main() {
    let mut table = Table::new();
    println!("{}", table);
    table.headers = vec![
        String::from("Model"),
        String::from("Piece N°"),
        String::from("In Stock"),
        String::from("Description"),
    ];
    table.add_row(&[
        String::from("model 1"),
        String::from("43-EWQE304"),
        String::from("30"),
        String::from("Piece for x"),
    ]);
    table.add_row(&[
        String::from("model 2"),
        String::from("98-QCVX5433"),
        String::from("100000000"),
        String::from("-"),
    ]);
    table.add_row(&[
        String::from("model y"),
        String::from("78-NMNH"),
        String::from("60"),
        String::from("nothing"),
    ]);
    println!("{}", table);
}
