/*
## even_iterator

### Instructions

Create a method `new` that takes one number `usize` and initializes the `Number` struct.
This method will have to determinate if the given number is even or odd,
if it is even you will have to increment one to the odd number and
if it is odd you have to increment one to the even number.

After that you will implement an iterator for the struct `Number` that returns a tuple (usize,usize,usize).
containing each field of the struct Number.
The first position of the tuple will be the even numbers,
the second will be the odd numbers,
and the third will be the factorial numbers.

### Notions

- https://doc.rust-lang.org/std/iter/trait.Iterator.html

*/

fn main() {
    let mut a = Number::new(5);
    println!("{:?}", a.next());     // Some((6, 5, 120))
    println!("{:?}", a.next());     // Some((8, 7, 720))
    println!("{:?}", a.next());     // Some((10, 9, 5040))
    println!()
    let mut a = Number::new(18); 
    println!("{:?}", a.next());     // Some((18, 19, 6402373705728000))
    println!("{:?}", a.next());     // Some((20, 21, 121645100408832000))
    println!("{:?}", a.next());     // Some((22, 23, 2432902008176640000))
}

struct Number {
    even: usize,
    odd: usize,
    fact: usize,
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_first_seven() {
        let test_even = vec![0, 2, 4, 6, 8, 10, 12];
        let test_odd = vec![1, 3, 5, 7, 9, 11, 13];
        let test_fact = vec![1, 1, 2, 6, 24, 120, 720];

        for (i, x) in Number::new(0).take(7).enumerate() {
            assert_eq!(x.0, test_even[i]);
            assert_eq!(x.1, test_odd[i]);
            assert_eq!(x.2, test_fact[i]);
        }
    }

    #[test]
    fn test_next() {
        let mut a = Number::new(6);
        assert_eq!(a.next().unwrap(), (6, 7, 720));
        assert_eq!(a.next().unwrap(), (8, 9, 5040));
        assert_eq!(a.next().unwrap(), (10, 11, 40320));
        assert_eq!(a.next().unwrap(), (12, 13, 362880));
        assert_eq!(a.next().unwrap(), (14, 15, 3628800));
    }
}
