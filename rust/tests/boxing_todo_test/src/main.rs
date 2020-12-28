use boxing_todo::*;

// Note that you can create some todo list your self to test it, but you can find the JSON files that
// are being tested [here](https://github.com/01-edu/public/blob/master/subjects/boxing_todo)
fn main() {
    let todos = TodoList::get_todo("todo.json");
    match todos {
        Ok(list) => println!("{:?}", list),
        Err(e) => {
            println!("{}{:?}", e.description(), e.cause());
        }
    }

    let todos = TodoList::get_todo("malforned_object.json");
    match todos {
        Ok(list) => println!("{:?}", list),
        Err(e) => {
            println!("{}{:?}", e.description(), e.cause().unwrap());
        }
    }

    let todos = TodoList::get_todo("permission_err.json");
    match todos {
        Ok(list) => println!("{:?}", list),
        Err(e) => {
            println!("{}{:?}", e.description(), e.cause().unwrap());
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::fs::File;
    use std::fs;

    fn new_todo(s: String, v: Vec<Task>) -> TodoList {
        TodoList { title: s, tasks: v }
    }
    fn run(s: &TodoList, f: &str) -> Result<TodoList, Box<dyn Error>> {
        serde_json::to_writer(&File::create(f)?, s)?;
        let result = TodoList::get_todo(f);
		fs::remove_file(f)?;
        return result;
    }

    #[test]
    fn test_good_todo() {
        let file_name = "todo.json";
        let good_struct = new_todo(
            String::from("todo list for something"),
            vec![
                Task { id: 0, description: "do this".to_string(), level: 0 },
                Task { id: 1, description: "do that".to_string(), level: 5 },
            ],
        );
        let result = run(&good_struct, file_name).unwrap();

        assert_eq!(result.title, good_struct.title);
        assert_eq!(&result.tasks, &good_struct.tasks);
    }

    #[test]
    fn test_empty_tasks() {
        let file_name = "empty_tasks.json";
        let result = run(&new_todo(String::from("empty tasks"), vec![]), file_name).unwrap_err();

        assert_eq!(result.to_string(), "Failed to parses todo");
        assert_eq!(result.description(), "Todo List parse failed: ");
        assert!(!result.cause().is_some());
    }
    #[test]
    fn test_read() {
        let result = TodoList::get_todo("no_file.json").unwrap_err();

        assert_eq!(result.to_string(), "Failed to read todo file");
        assert_eq!(result.description(), "Todo List read failed: ");
    }

    #[test]
    #[should_panic(expected = "Malformed(Error(\"missing field `title`\", line: 1, column: 2))")]
    fn test_malformed_json() {
        #[derive(Serialize, Deserialize)]
        struct Mal {};
        let file_name = "malformed.json";
        let malformed: Mal = serde_json::from_str(r#"{}"#).unwrap();
        serde_json::to_writer(&File::create(file_name).unwrap(), &malformed).unwrap();
        let result = TodoList::get_todo(file_name);
        fs::remove_file(file_name).unwrap();

        result.unwrap_or_else(|e| panic!("{:?}", e));
    }

    #[test]
    #[should_panic(expected = "ReadErr { child_err: Os { code: 2, kind: NotFound, message: \"No such file or directory\" } }")]
    fn test_read_error() {
        TodoList::get_todo("no_file.json").unwrap_or_else(|e| panic!("{:?}", e));
    }
}
