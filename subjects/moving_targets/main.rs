use moving_targets::*;

fn main() {
    let mut field = Field::new();

    println!("{:?}", field.pop());
    field.push(Target { size: 12, xp: 2 });
    println!("{:?}", *field.peek().unwrap());
    field.push(Target { size: 24, xp: 4 });
    println!("{:?}", field.pop());
    let last_target = field.peek_mut().unwrap();
    *last_target = Target { size: 2, xp: 0 };
    println!("{:?}", field.pop());
}
