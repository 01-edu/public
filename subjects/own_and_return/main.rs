use own_and_return::*;

pub struct Film {
    pub name: String,
}

fn main() {
    let my_film = Film {
        name: "Terminator".toString(),
    };
    println!("{}", take_film_name(/* to be implemented */));
    // the order of the print statements is intentional, if your implementation is correct,
    // you should have a compile error because my_film was consumed
    println!("{}", read_film_name(/* to be implemented */));
    println!("{}", take_film_name(/*to be implemented*/))
    // you can test this function by commenting out the first print statement,
    // you should see the expected output without errors in this case
}
