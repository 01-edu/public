use modify_letter::*;

fn main() {
    println!("{}", remove_letter_sensitive("Jojhn jis sljeepjjing", 'j'));
    println!("{}", remove_letter_insensitive("JaimA ais swiaAmmingA", 'A'));
    println!("{}", swap_letter_case("byE bye", 'e'));
}
