use ref_cell::*;

fn main() {
    // initialize the worker
    let Logger = Worker::new(1);

    // initialize the tracker, with the max number of
    // called references as 10
    let track = Tracker::new(&Logger, 10);

    let _a = Logger.track_value.clone();    // |\
    let _a1 = Logger.track_value.clone();   // | -> increase the Rc to 4 references
    let _a2 = Logger.track_value.clone();   // |/

    // take a peek of how much we already used from our quota
    track.peek(&Logger.track_value);

    let _b = Logger.track_value.clone();  // |\
    let _b1 = Logger.track_value.clone(); // |  -> increase the Rc to 8 references
    let _b2 = Logger.track_value.clone(); // | /
    let _b3 = Logger.track_value.clone(); // |/

    // this will set the value and making a verification of
    // how much we already used of our quota
    track.set_value(&Logger.track_value);

    let _c = Logger.track_value.clone(); // | -> increase the Rc to 9 references

    // this will set the value and making a verification of
    // how much we already used of our quota
    track.set_value(&Logger.track_value);

    let _c1 = Logger.track_value.clone(); // | -> increase the Rc to 10 references, this will be the limit

    track.set_value(&Logger.track_value);

    for (k ,v) in Logger.mapped_messages.into_inner() {
        println!("{:?}", (k ,v));
    }
    println!("{:?}", Logger.all_messages.into_inner());
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_module() {
        #[derive(Clone, Debug)]
        struct TestMs {
            value: Rc<usize>,
            ms: RefCell<Vec<String>>,
            correct: Vec<String>,
        }

        impl Logger for TestMs {
            fn warning(&self, message: &str) {
                self.ms.borrow_mut().push(message.to_string());
            }
            fn info(&self, message: &str) {
                self.ms.borrow_mut().push(message.to_string());
            }
            fn error(&self, message: &str) {
                self.ms.borrow_mut().push(message.to_string());
            }
        }

        let l = TestMs {
            value: Rc::new(115),
            ms: RefCell::new(vec![]),
            correct: vec![
              String::from("Info: you are using up too 40% of your quote"),
              String::from("Warning: you have used up over 80% of your quota! Proceeds with precaution"),
              String::from("Error: you are over your quota!")
            ],
        };

        let track = Tracker::new(&l, 5);
        let _a = l.value.clone();
        track.peek(&l.value); // 40%
        let _a1 = l.value.clone();
        let _a2 = l.value.clone();
        track.set_value(&l.value); // 80%
        let _a3 = l.value.clone();
        track.set_value(&l.value); // 100%

        for (i, v) in l.ms.into_inner().iter().enumerate() {
            assert_eq!(v, &l.correct[i])
        }
    }

    #[test]
    fn test_module_usage_hasmap() {
        let log = Worker::new(1000);
        let track = Tracker::new(&log, 12);
        let _clone_test = log.track_value.clone();
        let _clone_test1 = log.track_value.clone();
        let _clone_test2 = log.track_value.clone();
        let _clone_test3 = log.track_value.clone();
        let _clone_test4 = log.track_value.clone();
        let _clone_test5 = log.track_value.clone();
        let _clone_test6 = log.track_value.clone();
        let _clone_test7 = log.track_value.clone();
        
        // warning: 75% of the quota
        track.set_value(&log.track_value);
        assert_eq!(log.mapped_messages.borrow().get("Warning").unwrap(), "you have used up over 75% of your quota! Proceeds with precaution");

        let _clone_test8 = log.track_value.clone();

        // warning: 83% of the quota <- most resent of the messages last onw to be added to the hashmap
        track.set_value(&log.track_value);
        assert_eq!(log.mapped_messages.borrow().get("Warning").unwrap(), "you have used up over 83% of your quota! Proceeds with precaution");

        // info: 83%
        track.peek(&log.track_value);
        assert_eq!(log.mapped_messages.borrow().get("Info").unwrap(), "you are using up too 83% of your quote");

        let _clone_test9 = log.track_value.clone();
        // info: 91%
        track.peek(&log.track_value);
        assert_eq!(log.mapped_messages.borrow().get("Info").unwrap(), "you are using up too 91% of your quote");

        let _clone_test10 = log.track_value.clone();
        // error: passed the quota
        track.set_value(&log.track_value);
        assert_eq!(log.mapped_messages.borrow().get("Error").unwrap(), "you are over your quota!");
    }

    #[test]
    fn test_module_usage_vector() {
        let correct = vec![
            "Info: you are using up too 40% of your quote",
            "Warning: you have used up over 80% of your quota! Proceeds with precaution",
            "Info: you are using up too 80% of your quote", "Error: you are over your quota!"
        ];
        let log = Worker::new(1);
        let track = Tracker::new(&log, 5);
        let _a = log.track_value.clone();
        // info: 40%
        track.peek(&log.track_value);
        let _a1 = log.track_value.clone();
        let _a2 = log.track_value.clone();

        // warning: 80%
        track.set_value(&log.track_value);
        // info: 80%
        track.peek(&log.track_value);
        let _a3 = log.track_value.clone();

        // error: passed the quota
        track.set_value(&log.track_value);
        
        for (i, v) in log.all_messages.into_inner().iter().enumerate() {
            assert_eq!(v, correct[i]);
        }
    }
}
