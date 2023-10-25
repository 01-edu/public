fn main() {
    let mut list = Queue::new();
    list.add(String::from("Marie"), 20);
    list.add(String::from("Monica"), 15);
    list.add(String::from("Ana"), 5);
    list.add(String::from("Alice"), 35);
    println!("{:?}", list);

    println!("{:?}", list.search("Marie"));
    println!("{:?}", list.search("Alice"));
    println!("{:?}", list.search("someone"));

    println!("removed {:?}", list.rm());
    println!("list {:?}", list);
    list.invert_queue();
    println!("invert {:?}", list);
}
