use count_factorial_steps::count_factorial_steps;

fn main() {
    println!(
        "The factorial steps of 720 = {}",
        count_factorial_steps(720)
    );
    println!("The factorial steps of 13 = {}", count_factorial_steps(13));
    println!("The factorial steps of 6 = {}", count_factorial_steps(6));
}
