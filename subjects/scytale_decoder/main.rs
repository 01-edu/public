use scytale_decoder::*;

fn main() {
    println!(
        "\"sec yCtoadle\" size=2 -> {:?}",
        scytale_decoder("sec yCtoadle".to_string(), 2)
    );

    println!(
        "\"steoca dylCe\" size=4 -> {:?}",
        scytale_decoder("steoca dylCe".to_string(), 4)
    );
}
