use own_and_return::*;

fn main() {
    let my_film = Film {
        name: "Terminator".to_owned(),
    };

    // println!("{}", take_film_name(/* to be implemented */));

    println!("{}", read_film_name(/* to be implemented */));
    println!("{}", take_film_name(/* to be implemented */));

    // the order of the print statements is intentional.
    // you can test this exercise properly by uncommenting out the first print statement,
    // you should get a compilation error then if your implementation is correct.
}
