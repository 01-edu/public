use format_me::*;

fn main() {
    println!(
        "{}",
        Park {
            name: Some("Les Tuileries".to_owned()),
            park_type: ParkType::Garden,
            address: Some("Pl. de la Concorde".to_owned()),
            cap: Some("75001".to_owned()),
            state: Some("France".to_owned())
        }
    );
    println!(
        "{}",
        Park {
            name: None,
            park_type: ParkType::Playground,
            address: None,
            cap: None,
            state: None
        }
    );
}
