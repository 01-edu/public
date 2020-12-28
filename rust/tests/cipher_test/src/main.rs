use cipher::*;

fn main() {
    println!("{:?}", cipher("1Hello 2world!", "1Svool 2dliow!"));
    println!("{:?}", cipher("1Hello 2world!", "svool"));
    println!("{:?}", cipher("", "svool"));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_cipher() {
        assert_eq!(cipher("1Hello 2world!", "1Svool 2dliow!"), Some(Ok(true)));
        assert_eq!(cipher("", "1Svool 2dliow!"), None);
        assert_eq!(cipher("1Hello 2world!", ""), None);
        assert_eq!(cipher("1Hello 2world!", "1svool 2dliow!"), Some(Err(CipherError { validation: false, expected: String::from("1Svool 2dliow!") })));
        assert_eq!(cipher("asdasd", "zhwzhw"), Some(Ok(true)));
        assert_eq!(cipher("asdasd", "lkdas"), Some(Err(CipherError { validation: false, expected: String::from("zhwzhw") })));
        assert_eq!(cipher("3(/&%fsd 32das", "3(/&%uhw 32wzh"), Some(Ok(true)));
        assert_eq!(cipher("3(/&%sd 32das", "3(/&%uhw 32wzh"), Some(Err(CipherError { validation: false, expected: String::from("3(/&%hw 32wzh") })));
    }
}