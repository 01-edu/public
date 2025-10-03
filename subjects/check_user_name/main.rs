use check_user_name::*;

fn main() {
    let user0 = User::new("Didier".to_owned(), AccessLevel::Admin);
    println!("{:?}", check_user_name(&user0));

    let user1 = User::new("Mary".to_owned(), AccessLevel::Normal);
    println!("{:?}", check_user_name(&user1));

    let user2 = User::new("John".to_owned(), AccessLevel::Guest);
    println!("{:?}", check_user_name(&user2));
}
