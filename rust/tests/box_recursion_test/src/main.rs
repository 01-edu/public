use box_recursion::*;

fn main() {
    let mut list = WorkEnvironment::new();
    list.add_worker(String::from("CEO"), String::from("Marie"));
    list.add_worker(String::from("Manager"), String::from("Monica"));
    list.add_worker(String::from("Normal Worker"), String::from("Ana"));
    list.add_worker(String::from("Normal Worker"), String::from("Alice"));
    println!("{:?}", list);

    println!("{:?}", list.search_worker());

    list.remove_worker();
    list.remove_worker();
    list.remove_worker();
    list.remove_worker();
    println!("{:?}", list);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_new() {
        let list = WorkEnvironment::new();
        assert!(list.grade.is_none());
    }

    #[test]
    fn test_one_worker() {
        let mut list = WorkEnvironment::new();
        list.add_worker(String::from("CEO"), String::from("Marie"));
        list.remove_worker();
        assert!(list.grade.is_none());
    }

    #[test]
    fn test_two_workers() {
        let mut list = WorkEnvironment::new();
        list.add_worker(String::from("CEO"), String::from("Marie"));
        list.add_worker(String::from("Manager"), String::from("Monica"));
        list.remove_worker();

        assert_eq!(list.grade.as_ref().unwrap().worker_type, "CEO");
        assert_eq!(list.grade.as_ref().unwrap().worker_name, "Marie");
    }

    #[test]
    fn test_more_workers() {
        let mut list = WorkEnvironment::new();
        list.add_worker(String::from("CEO"), String::from("Marie"));
        list.add_worker(String::from("Manager"), String::from("Monica"));
        list.add_worker(String::from("Normal Worker"), String::from("Ana"));
        list.add_worker(String::from("Normal Worker"), String::from("Alice"));
        list.remove_worker();

        assert_eq!(list.grade.as_ref().unwrap().worker_type, "Normal Worker");
        assert_eq!(list.grade.as_ref().unwrap().worker_name, "Ana");

        list.remove_worker();
        list.remove_worker();
        assert_eq!(list.grade.as_ref().unwrap().worker_type, "CEO");
        assert_eq!(list.grade.as_ref().unwrap().worker_name, "Marie");
    }

    #[test]
    fn test_search() {
        let mut list = WorkEnvironment::new();
        list.add_worker(String::from("CEO"), String::from("Marie"));
        list.add_worker(String::from("Manager"), String::from("Monica"));
        list.add_worker(String::from("Normal Worker"), String::from("Ana"));
        list.add_worker(String::from("Normal Worker"), String::from("Alice"));

        assert_eq!(
            list.search_worker().unwrap(),
            (String::from("Alice"), String::from("Normal Worker"))
        );

        list.remove_worker();
        assert_eq!(
            list.search_worker().unwrap(),
            (String::from("Ana"), String::from("Normal Worker"))
        );

        list.remove_worker();
        assert_eq!(
            list.search_worker().unwrap(),
            (String::from("Monica"), String::from("Manager"))
        );

        list.remove_worker();
        assert_eq!(
            list.search_worker().unwrap(),
            (String::from("Marie"), String::from("CEO"))
        );
    }
}
