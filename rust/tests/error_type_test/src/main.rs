use error_type::*;

#[allow(dead_code)]
fn create_date(date: &str) -> NaiveDate {
    NaiveDate::parse_from_str(date, "%Y-%m-%d").unwrap()
}

fn main() {
    let mut form_output = Form::new(
        String::from("Lee"),
        String::from("Silva"),
        create_date("2015-09-05"),
        SexType::Male,
        String::from("Africa"),
        String::from("qwqwsa1dty_"));

    println!("{:?}", form_output);
    println!("{:?}", form_output.validate().unwrap());

    form_output.first_name = String::from("");
    println!("{:?}", form_output.validate().unwrap_err());

    form_output.first_name = String::from("as");
    form_output.password = String::from("dty_1");
    println!("{:?}", form_output.validate().unwrap_err());

    form_output.password = String::from("asdasASd(_");
    println!("{:?}", form_output.validate().unwrap_err());

    form_output.password = String::from("asdasASd123SA");
    println!("{:?}", form_output.validate().unwrap_err());
}

#[cfg(test)]
mod tests {
    use super::*;

    #[derive(Debug)]
    struct TestForm<'a> {
        form: Form,
        validation: Result<Vec<&'a str>, FErr>
    }

    impl <'a> TestForm<'_> {
        // all test cases
        fn new() -> Vec<TestForm<'a>> {
            vec![
                TestForm {
                    form : Form::new(
                    String::from("Katy"),
                    String::from("Silva"),
                    create_date("2015-09-05"),
                    SexType::Female,
                    String::from("Africa"),
                    String::from("qwTw12&%$3sa1dty_")),
                    validation: Ok(vec!["Valid first name", "Valid password"]),
                },
                TestForm {
                    form : Form::new(
                    String::from(""),
                    String::from("Bear"),
                    create_date("2015-09-05"),
                    SexType::Male,
                    String::from("Africa"),
                    String::from("qwTw12&%$3sa1dty_")),
                    validation: Err(FErr {
                        form_values: (String::from("first_name"),
                        String::from("")),
                        date: Utc::now().format("%Y-%m-%d %H:%M:%S").to_string(),
                        err: String::from("No user name")}),
                },
                TestForm {
                    form : Form::new(
                    String::from("Someone"),
                    String::from("Bear"),
                    create_date("2015-09-05"),
                    SexType::Male,
                    String::from("Africa"),
                    String::from("12345")),
                    validation: Err(FErr {
                        form_values: (String::from("password"), String::from("12345")),
                        date: Utc::now().format("%Y-%m-%d %H:%M:%S").to_string(),
                        err: String::from("At least 8 characters") }),
                },
                TestForm {
                    form : Form::new(
                    String::from("Someone"),
                    String::from("Bear"),
                    create_date("2015-09-05"),
                    SexType::Male,
                    String::from("Africa"),
                    String::from("sdASDsrW")),
                    validation: Err(FErr {
                        form_values: (String::from("password"), String::from("sdASDsrW")),
                        date: Utc::now().format("%Y-%m-%d %H:%M:%S").to_string(),
                        err: String::from("Combination of different ASCII character types (numbers, letters and none alphanumeric characters)") }),
                },
                TestForm {
                    form : Form::new(
                    String::from("Someone"),
                    String::from("Bear"),
                    create_date("2015-09-05"),
                    SexType::Female,
                    String::from("Africa"),
                    String::from("dsGE1SAD213")),
                    validation: Err(FErr {
                        form_values: (String::from("password"), String::from("dsGE1SAD213")),
                        date: Utc::now().format("%Y-%m-%d %H:%M:%S").to_string(),
                        err: String::from("Combination of different ASCII character types (numbers, letters and none alphanumeric characters)") }),
                },
                TestForm {
                    form : Form::new(
                    String::from("Someone"),
                    String::from("Bear"),
                    create_date("2015-09-05"),
                    SexType::Female,
                    String::from("Africa"),
                    String::from("dsaSD&%DF!?=")),
                    validation: Err(FErr {
                        form_values: (String::from("password"), String::from("dsaSD&%DF!?=")),
                        date: Utc::now().format("%Y-%m-%d %H:%M:%S").to_string(),
                        err: String::from("Combination of different ASCII character types (numbers, letters and none alphanumeric characters)") }),
                }
            ]
        }
    }

    #[test]
    fn test_error_type() {
        let form_cases = TestForm::new();

        for v in form_cases {
            assert_eq!(v.form.validate(), v.validation);
        }
    }
}