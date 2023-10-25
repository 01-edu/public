use profanity_filter::*;

fn main() {
    let m0 = Message::new("hello there".to_string(), "toby".to_string());
    println!("{:?}", check_ms(&m0));

    let m1 = Message::new("".to_string(), "toby".to_string());
    println!("{:?}", check_ms(&m1));

    let m2 = Message::new("you are stupid".to_string(), "toby".to_string());
    println!("{:?}", check_ms(&m2));

    let m3 = Message::new("stupid".to_string(), "toby".to_string());
    println!("{:?}", check_ms(&m3));
}
