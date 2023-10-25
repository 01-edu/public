use office_worker::*;

fn main() {
    println!("New worker: {:?}", OfficeWorker::from("Manuel,23,admin"));
    println!(
        "New worker: {:?}",
        OfficeWorker::from("Jean Jacques,44,guest")
    );
}
