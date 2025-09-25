use inv_pyramid::*;

fn main() {
    println!("{:#?}", inv_pyramid(String::from("#"), 1));
    println!("{:#?}", inv_pyramid(String::from("a"), 2));
    println!("{:#?}", inv_pyramid(String::from(">"), 5));
    println!("{:#?}", inv_pyramid(String::from("&"), 8));
}
