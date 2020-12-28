use std::collections::HashMap;
use banner::*;

fn main() {
    let mut handler = FlagsHandler { flags: HashMap::new() };

    let d = Flag::opt_flag("division", "divides the values, formula (a / b)");
    let r = Flag::opt_flag(
        "remainder",
        "remainder of the division between two values, formula (a % b)",
    );

    handler.add_flag((d.short_hand, d.long_hand), div);
    handler.add_flag((r.short_hand, r.long_hand), rem);

    println!("{:?}", handler.exec_func(("-d".to_string(), "--division".to_string()), &["1.0", "2.0"]));
    // output: "0.5"

    println!("{:?}",handler.exec_func(("-r".to_string(), "--remainder".to_string()), &["2.0", "2.0"]));
    // output: "0.0"
    
    println!("{:?}",handler.exec_func(("-d".to_string(), "--division".to_string()), &["a", "2.0"]));
    // output: "invalid float literal"
    
    println!("{:?}",handler.exec_func(("-r".to_string(), "--remainder".to_string()), &["2.0", "fd"]));
    // output: "invalid float literal"
}

#[cfg(test)]
mod tests {
    use super::*;

    fn init() -> FlagsHandler {
        let d = Flag::opt_flag("division", "divides two numbers");
        let r = Flag::opt_flag(
            "remainder",
            "gives the remainder of the division between two numbers",
        );
        let mut handler = FlagsHandler { flags: HashMap::new() };

        handler.add_flag((d.short_hand, d.long_hand), div);
        handler.add_flag((r.short_hand, r.long_hand), rem);
        return handler;
    }

    #[test]
    fn ok_test() {
        let mut handler = init();
        assert_eq!(
            handler.exec_func(("-d".to_string(), "--division".to_string()), &["1.0", "2.0"]),
            "0.5"
        );
        assert_eq!(
            handler.exec_func(("-r".to_string(), "--remainder".to_string()), &["2.0", "2.0"]),
            "0"
        );
        assert_eq!(
            handler.exec_func(("-d".to_string(), "--division".to_string()), &["12.323", "212.32"]),
            "0.05803975"
        );
        assert_eq!(
            handler.exec_func(("-r".to_string(), "--remainder".to_string()), &["12.323", "212.32"]),
            "12.323"
        );
    }

    #[test]
    fn error_test() {
        let mut handler = init();
        assert_eq!(
            handler.exec_func(("-d".to_string(), "--division".to_string()), &["a", "2.0"]),
            "invalid float literal"
        );
        assert_eq!(
            handler.exec_func(("-r".to_string(), "--remainder".to_string()), &["2.0", "f"]),
            "invalid float literal"
        );
        assert_eq!(
            handler.exec_func(("-d".to_string(), "--division".to_string()), &["1.0", "0.0"]),
            "inf"
        );
        assert_eq!(
            handler.exec_func(("-r".to_string(), "--remainder".to_string()), &["2.0", "0.0"]),
            "NaN"
        );
    }
}

