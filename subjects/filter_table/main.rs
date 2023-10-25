use filter_table::*;

fn main() {
    let mut table = Table::new();
    table.headers = vec![
        "Name".to_string(),
        "Last Name".to_string(),
        "ID Number".to_string(),
    ];
    table.add_row(&[
        "Adam".to_string(),
        "Philips".to_string(),
        "123456789".to_string(),
    ]);
    table.add_row(&[
        "Adamaris".to_string(),
        "Shelby".to_string(),
        "1111123456789".to_string(),
    ]);
    table.add_row(&[
        "Ackerley".to_string(),
        "Philips".to_string(),
        "123456789".to_string(),
    ]);
    let filter_names = |col: &str| col == "Name";
    println!("{:?}", table.filter_col(filter_names));

    let filter_philips = |lastname: &str| lastname == "Philips";
    println!("{:?}", table.filter_row("Last Name", filter_philips));
}
