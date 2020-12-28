use box_it::*;

fn main() {
    let new_str = String::from("5.5k 8.9k 32");

    // creating a variable and we save it in the Heap
    let a_h = transform_and_save_on_heap(new_str);
    println!("Box value : {:?}", &a_h);
    println!("size occupied in the stack : {:?} bytes", (std::mem::size_of_val(&a_h)));
    
    let a_b_v = take_value_ownership(a_h);
    println!("value : {:?}", &a_b_v);
    println!("size occupied in the stack : {:?} bytes", (std::mem::size_of_val(&a_b_v)));
    // whenever the box, in this case "a_h", goes out of scope it will be deallocated, freed
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::mem;

    #[test]
    fn test_transform() {
        let new_str = String::from("5.5k 8.9k 32");
        let new_str_1 = String::from("6.8k 13.5k");
        let new_str_2 = String::from("20.3k 3.8k 7.7k 992");

        let a = transform_and_save_on_heap(new_str);
        let b = transform_and_save_on_heap(new_str_1);
        let c = transform_and_save_on_heap(new_str_2);

        assert_eq!(a, Box::new(vec![5500, 8900, 32]));
        assert_eq!(b, Box::new(vec![6800, 13500]));
        assert_eq!(c, Box::new(vec![20300, 3800, 7700, 992]));
        assert_eq!(mem::size_of_val(&a), 8);
        assert_eq!(mem::size_of_val(&b), 8);
        assert_eq!(mem::size_of_val(&c), 8);
    }

    #[test]
    fn test_take_value_from_box() {
        let new_str = String::from("5.5k 8.9k 32");
        let new_str_1 = String::from("6.8k 13.5k");
        let new_str_2 = String::from("20.3k 3.8k 7.7k 992");
        let a = take_value_ownership(transform_and_save_on_heap(new_str));
        let b = take_value_ownership(transform_and_save_on_heap(new_str_1));
        let c = take_value_ownership(transform_and_save_on_heap(new_str_2));

        assert_eq!(a, vec![5500, 8900, 32]);
        assert_eq!(b, vec![6800, 13500]);
        assert_eq!(c, vec![20300, 3800, 7700, 992]);
        assert_eq!(mem::size_of_val(&a), 24);
        assert_eq!(mem::size_of_val(&b), 24);
        assert_eq!(mem::size_of_val(&c), 24);
    }
}
