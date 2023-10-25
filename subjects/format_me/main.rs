use format_me::*;

fn main() {
    println!(
        "{}",
        Park {
            name: "Les Tuileries".to_string(),
            park_type: ParkType::Garden,
            address: "Pl. de la Concorde".to_string(),
            cap: "75001".to_string(),
            state: "France".to_string()
        }
    );
    println!(
        "{}",
        Park {
            name: "".to_string(),
            park_type: ParkType::Playground,
            address: "".to_string(),
            cap: "".to_string(),
            state: "".to_string()
        }
    );
}
